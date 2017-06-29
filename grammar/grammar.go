package grammar

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Yara-Rules/yago/lexic"

	"github.com/Sirupsen/logrus"
)

// New returns a new parser
func New(name string) *Parser {
	return &Parser{
		Name:      name,
		peekCount: 0,
		log:       logrus.New(),
	}
}

// SetLogLevel sets the log level
func (p *Parser) SetLogLevel(level string) {
	if level != "" {
		var lvl logrus.Level
		lvl, err := logrus.ParseLevel(level)
		if err != nil {
			fmt.Printf("Not a valid log Level: %q.\nSetting info level as default.\n", level)
		} else {
			p.log.Level = lvl
		}
	}
}

// Parse starts parsing the input
func (p *Parser) Parse(text string) {
	p.log.Debugln(" ** Stating parser **")
	p.startParer(text)
	p.parse()
	p.log.Debugln(" ** Parser finished **")
}

func (p *Parser) startParer(text string) {
	p.Lex = lexic.Lex(p.Name, text)
}

func (p *Parser) errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)

	err["syntactical"] = "error"
	err["line"] = fmt.Sprintf("%d", p.LastItem.GetLine())
	err["file_name"] = p.Name
	err["msg"] = msg

	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}

func (p *Parser) warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)

	err["syntactical"] = "warning"
	err["line"] = fmt.Sprintf("%d", p.LastItem.GetLine())
	err["msg"] = msg

	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
}

func (p *Parser) getLastItem() lexic.Item {
	return p.token[p.peekCount]
}

// backup backs the input stream up one token.
func (p *Parser) backup() {
	p.peekCount++
}

// peek returns but does not consume the next token.
func (p *Parser) peek() lexic.Item {
	if p.peekCount > 0 {
		return p.token[p.peekCount-1]
	}
	p.peekCount = 1
	p.token[0] = p.nextNotComment()
	return p.token[0]
}

func (p *Parser) nextNotComment() lexic.Item {
	item := p.Lex.NextItem()
	for checkItemType(item, "__COMMENT__") { // Omit all comments
		item = p.Lex.NextItem()
	}
	return item
}

func (p *Parser) nextItem() lexic.Item {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		p.token[0] = p.nextNotComment()
	}
	p.LastItem = p.token[p.peekCount]
	return p.token[p.peekCount]
}

func (p *Parser) parse() {
	private := false
	global := false
	for item := range p.Lex.Items {
		// p.log.Debugln("--> ", item)
		switch {
		case checkItemType(item, "__KW_IMPORT__"):
			p.processImport()
		case checkItemType(item, "__KW_PRIVATE__"):
			private = true
			item = p.Lex.NextItem()
			if checkItemType(item, "__KW_GLOBAL__") {
				global = true
				item = p.Lex.NextItem()
			}
			if checkItemType(item, "__KW_RULE__") {
				p.processRule(global, private)
				private = false
				global = false
			}
		case checkItemType(item, "__KW_GLOBAL__"):
			global = true
			item = p.Lex.NextItem()
			if checkItemType(item, "__KW_PRIVATE__") {
				private = true
				item = p.Lex.NextItem()
			}
			if checkItemType(item, "__KW_RULE__") {
				p.processRule(global, private)
				private = false
				global = false
			}
		case checkItemType(item, "__KW_RULE__"):
			p.processRule(global, private)
		}
	}
}

func (p *Parser) processImport() {
	item := p.nextItem()
	if checkItemType(item, "__STRING__") {
		if !p.addImport(item) {
			p.warnf("Module %s already imported.\n", item.GetValue())
		} else {
			p.log.Debugln("Importing module: ", item)
		}
	} else {
		p.errorf("Expected %s and found %s.\n", lexic.ItemType["ItemString"], item.GetType())
	}
}

func (p *Parser) processRule(global, private bool) {
	item := p.nextItem()
	if checkItemType(item, "__IDENTIFIER__") {
		if !p.ruleAlreadyImported(item.GetValue()) {
			p.log.Debugln("Processing rule: ", item)
			newRule := RuleDef{
				Name:    item.GetValue(),
				Global:  global,
				Private: private,
			}
			item = p.nextItem()
			if checkItemType(item, "__COLON__") { // Tags comming
				newRule.Tags = p.processTags()
				p.log.Debugln("Tags list: ", newRule.Tags)
			}
			if checkItemType(p.LastItem, "__OPEN_CURLY__") {
				item = p.nextItem()
				if checkItemType(item, "__KW_META__") { // Meta comming
					item = p.nextItem()
					if checkItemType(item, "__COLON__") {
						newRule.Meta = p.processMeta()
						if len(newRule.Meta) == 0 {
							p.errorf("%s found but not meta items defined", lexic.ItemType["ItemKWMeta"])
						}
					} else {
						p.errorf("Expected %s found %s", lexic.ItemType["ItemColon"], item.GetType())
					}
				}
				if checkItemType(p.LastItem, "__KW_STRINGS__") { // Strings comming
					item = p.nextItem()
					if checkItemType(item, "__COLON__") {
						newRule.Strings = p.processStrings()
						if len(newRule.Strings) == 0 {
							p.errorf("%s found but not strings defined", lexic.ItemType["ItemKWStrings"])
						}
					} else {
						p.errorf("Expected %s found %s", lexic.ItemType["ItemColon"], item.GetType())
					}
				}
				if checkItemType(p.LastItem, "__KW_CONDITION__") { // Condition comming
					item = p.nextItem()
					if checkItemType(item, "__COLON__") {
						newRule.Condition = p.processCondition()
						if len(newRule.Condition) == 0 {
							p.errorf("%s found but not condition defined", lexic.ItemType["ItemKWCondition"])
						} else {
							p.log.Debugln("Condition: ", newRule.Condition)
						}
						p.addRule(newRule)
					} else {
						p.errorf("Expected %s found %s", lexic.ItemType["ItemColon"], item.GetType())
					}
				} else {
					p.errorf("Expected %s or %s or %s found %s", lexic.ItemType["ItemKWMeta"], lexic.ItemType["ItemKWStrings"], lexic.ItemType["ItemKWCondition"], item.GetType())
				}
			} else {
				p.errorf("Expected %s found %s.", lexic.ItemType["ItemOCurly"], item.GetType())
			}
		} else {
			p.errorf("Rule %s alredy defined.", item.GetValue())
		}
	} else {
		p.errorf("Expected %s found %s.", lexic.ItemType["ItemIdentifier"], item.GetType())
	}
}

func (p *Parser) processTags() []string {
	var tags []string
	item := p.nextItem()
	for !checkItemType(item, "__OPEN_CURLY__") {
		if checkItemType(item, "__IDENTIFIER__") {
			tags = append(tags, item.GetValue())
		} else {
			p.errorf("Expecting %s or %s and found %s.", lexic.ItemType["ItemIdentifier"], lexic.ItemType["ItemOCurly"], item.GetType())
		}
		item = p.nextItem()
	}
	return tags
}

func (p *Parser) processMeta() map[string]string {
	var key, value lexic.Item
	meta := make(map[string]string)
	item := p.nextItem()
	for !checkItemType(item, "__KW_STRINGS__") && !checkItemType(item, "__KW_CONDITION__") {
		key = item
		if checkItemType(item, "__IDENTIFIER__") {
			item = p.nextItem()
			if checkItemType(item, "__EQUAL__") {
				item = p.nextItem()
				value = item
				if checkItemType(item, "__STRING__") ||
					checkItemType(item, "__INT_NUMBER__") ||
					checkItemType(item, "__KW_TRUE__") || // Yara allows boolans as values
					checkItemType(item, "__KW_FLASE__") {
					p.log.Debugln("Meta: ", key, " = ", item)
					meta[key.GetValue()] = value.GetValue()
				} else {
					p.errorf("Expected %s or %s found %s", lexic.ItemType["ItemString"], lexic.ItemType["ItemIntNumber"], item.GetType())
				}
			} else {
				p.errorf("Expected %s found %s", lexic.ItemType["ItemEqual"], item.GetType())
			}
		} else {
			p.errorf("Expected %s found %s", lexic.ItemType["ItemIdentifier"], item.GetType())
		}
		item = p.nextItem()
	}
	return meta
}

func (p *Parser) processStrings() []StringDef {
	var key lexic.Item
	var value string
	var strings []StringDef
	stringTable := []string{}

	item := p.nextItem()
	for !checkItemType(item, "__KW_CONDITION__") {
		key = item
		if !stringDefined(stringTable, item) {
			stringTable = append(stringTable, item.GetValue())
			if checkItemType(item, "__VARIABLE__") {
				item = p.nextItem()
				if checkItemType(item, "__EQUAL__") {
					item = p.nextItem()
					if checkItemType(item, "__STRING__") ||
						checkItemType(item, "__KW_TRUE__") || // Yara allows boolans as values
						checkItemType(item, "__KW_FLASE__") {
						value, mods := p.processStringModifiers()
						strings = append(strings, StringDef{Name: key.GetValue(), Value: value, Modifiers: mods, Typ: StringString})
						p.log.Debugln("String: ", key, " = ", value, " ", mods)
					} else if checkItemType(item, "__REGEX__") {
						value, mods := p.processStringModifiers()
						strings = append(strings, StringDef{Name: key.GetValue(), Value: value, Modifiers: mods, Typ: StringRegex})
						p.log.Debugln("String: ", key, " = ", value, " ", mods)
					} else if checkItemType(item, "__OPEN_CURLY__") {
						value = p.processHexValues()
						strings = append(strings, StringDef{Name: key.GetValue(), Value: value, Typ: StringHex})
						p.log.Debugln("String: ", key, " = ", value)
					}
				} else {
					p.errorf("Expected %s found %s", lexic.ItemType["ItemEqual"], item.GetType())
				}
			} else {
				p.errorf("Expected %s found %s", lexic.ItemType["ItemVariable"], item.GetType())
			}
		} else {
			p.errorf("Duplicated string identifier %s", item.GetValue())
		}
		item = p.nextItem()
	}
	return strings
}

func (p *Parser) processStringModifiers() (string, []string) {
	value := p.LastItem
	var mods []string
	item := p.peek()
	if isStringModifier(item) {
		item = p.nextItem()
		for isStringModifier(item) {
			mods = append(mods, item.GetValue())
			item = p.nextItem()
		}
		p.backup()
	}
	return value.GetValue(), mods
}

func (p *Parser) processHexValues() string {
	value := "{"
	item := p.nextItem()
	for !checkItemType(item, "__CLOSE_CURLY__") {
		switch {
		case checkItemType(item, "__IDENTIFIER__"): // Hex value
			value = value + item.GetValue()
			break
		case checkItemType(item, "__INT_NUMBER__"): // Number
			value = value + item.GetValue()
			break
		case checkItemType(item, "__QMAKR__"): // ? Wildcard
			value = value + item.GetValue()
			break
		case checkItemType(item, "__OPEN_SQRT__"):
			value = value + p.processHexRange()
		case checkItemType(item, "__OPEN_BRACKET__"):
			value = value + p.preocessHexOption()
		}
		item = p.nextItem()
	}
	value = value + "}"
	return value
}

func (p *Parser) processHexRange() string {
	value := "["
	item := p.nextItem()
	if checkItemType(item, "__INT_NUMBER__") {
		numA := item.GetValue()
		nA, err := strconv.Atoi(numA)
		if err == nil {
			if nA < 0 {
				p.errorf("Expected >= 0 integer found %s", lexic.ItemType["ItemCSqrt"], item.GetType())
			}
		} else {
			p.errorf("Unable to convert %s", item.GetType())
		}
		item = p.nextItem()
		if checkItemType(item, "__DASH__") {
			dash := item.GetValue()
			item = p.nextItem()
			if checkItemType(item, "__INT_NUMBER__") {
				numB := item.GetValue()
				nB, err := strconv.Atoi(numB)
				if err == nil {
					if nB < nA {
						p.errorf("Expected %s to be grather than %s", numB, numA)
					}
				} else {
					p.errorf("Unable to convert %s", item.GetType())
				}
				item = p.nextItem()
				if checkItemType(item, "__CLOSE_SQRT__") {
					value = value + numA + dash + numB + item.GetValue()
					return value
				} else {
					p.errorf("Expected %s found %s", lexic.ItemType["ItemCSqrt"], item.GetType())
				}
			} else if checkItemType(item, "__CLOSE_SQRT__") {
				value = value + numA + dash + item.GetValue()
				return value
			} else {
				p.errorf("Expected %s or  found %s", lexic.ItemType["ItemIntNumber"], lexic.ItemType["ItemCSqrt"], item.GetType())
			}
		} else if checkItemType(item, "__CLOSE_SQRT__") {
			value = value + numA + item.GetValue()
			return value
		} else {
			p.errorf("Expected %s found %s", lexic.ItemType["ItemDash"], item.GetType())
		}
	} else if checkItemType(item, "__DASH__") {
		dash := item.GetValue()
		item = p.nextItem()
		if checkItemType(item, "__CLOSE_SQRT__") {
			value = value + dash + item.GetValue()
			return value
		} else {
			p.errorf("Expected %s found %s", lexic.ItemType["ItemCSqrt"], item.GetType())
		}
	} else {
		p.errorf("Expected %s or %s found %s", lexic.ItemType["ItemIntNumber"], lexic.ItemType["ItemDash"], item.GetType())
	}
	value = value + item.GetValue()
	return value
}

func (p *Parser) preocessHexOption() string {
	value := "("
	item := p.nextItem()
	for !checkItemType(item, "__CLOSE_BRACKET__") {
		switch {
		case checkItemType(item, "__IDENTIFIER__"): // Hex value
			value = value + item.GetValue()
		case checkItemType(item, "__INT_NUMBER__"): // Number
			value = value + item.GetValue()
		case checkItemType(item, "__QMAKR__"): // ? Wildcard
			value = value + item.GetValue()
		case checkItemType(item, "__PIPE__"): // |
			value = value + item.GetValue()
		case checkItemType(item, "__OPEN_SQRT__"): // [
			value = value + p.processHexRange()
		default:
			p.errorf("Expected one of %s, %s, %s, or %s and found %s", lexic.ItemType["ItemIdentifier"], lexic.ItemType["ItemIntNumber"], lexic.ItemType["ItemQMark"], lexic.ItemType["ItemPipe"], item.GetType())
		}
		item = p.nextItem()
	}
	value = value + item.GetValue()
	return value
}

func (p *Parser) processCondition() string {
	var last lexic.Item
	value := ""
	space := ""
	item := p.nextItem()
	last = item
	for !checkItemType(item, "__CLOSE_CURLY__") {
		if checkItemType(last, "__DOT__") || checkItemType(last, "__DOT_DOT__") {
			space = ""
		}
		if checkItemType(item, "__HASH__") {
			hash := item.GetValue()
			item = p.nextItem()
			if checkItemType(item, "__IDENTIFIER__") {
				value += space + hash + item.GetValue()
			} else {
				p.errorf("Expected %s and found %s", lexic.ItemType["ItemIdentifier"], item.GetType())
			}
		} else if checkItemType(item, "__IDENTIFIER__") {
			id := item.GetValue()
			if checkItemType(p.peek(), "__DOT__") {
				dot := p.nextItem()
				if checkItemType(p.peek(), "__IDENTIFIER__") {
					item := p.nextItem()
					value += space + id + dot.GetValue() + item.GetValue()
				} else {
					p.errorf("Expected %d and found %s", lexic.ItemType["ItemIdentifier"], item.GetType())
				}
			} else {
				value += space + id
			}
		} else if checkItemType(item, "__VARIABLE__") {
			v := item.GetValue()
			if checkItemType(p.peek(), "__STAR__") {
				star := p.nextItem()
				value += space + v + star.GetValue()
			} else {
				value += space + v
			}
		} else if checkItemType(item, "__DOT__") || checkItemType(item, "__DOT_DOT__") {
			value += item.GetValue()
		} else {
			value += space + item.GetValue()
		}
		space = " "
		last = p.LastItem
		item = p.nextItem()
	}
	p.log.Debugln("Condition: ", value)
	return value
}
