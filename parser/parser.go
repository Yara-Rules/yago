package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/Yara-Rules/yago/lexer"
)

func New(name string) *Parser {
	return &Parser{
		Name:      name,
		peekCount: 0,
		log:       logrus.New(),
	}
}

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

func (p *Parser) Parse(text string) {
	p.startParer(text)
	p.parse()
}

func (p *Parser) startParer(text string) {
	p.log.Debugln("Starting parser")
	p.Lex = lexer.Lex(p.Name, text)
}

func (p *Parser) errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)
	err["error"] = msg
	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}

func (p *Parser) warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)
	err["warning"] = msg
	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
}

// backup backs the input stream up one token.
func (p *Parser) backup() {
	p.peekCount++
}

// backup2 backs the input stream up two tokens.
// The zeroth token is already there.
func (p *Parser) backup2(i1 lexer.Item) {
	p.token[1] = i1
	p.peekCount = 2
}

// backup3 backs the input stream up three tokens
// The zeroth token is already there.
func (p *Parser) backup3(i2, i1 lexer.Item) { // Reverse order: we're pushing back.
	p.token[1] = i1
	p.token[2] = i2
	p.peekCount = 3
}

// peek returns but does not consume the next token.
func (p *Parser) peek() lexer.Item {
	if p.peekCount > 0 {
		return p.token[p.peekCount-1]
	}
	p.peekCount = 1
	p.token[0] = p.Lex.NextItem()
	return p.token[0]
}

func (p *Parser) nextItem() lexer.Item {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		p.token[0] = p.Lex.NextItem()
	}
	p.LastItem = p.token[p.peekCount]
	return p.LastItem
}

func (p *Parser) parse() {
	p.log.Debugln("Parsing")
	private := false
	global := false
	for item := range p.Lex.Items {
		switch {
		case checkItemType(item, lexer.ItemImport):
			p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemImport])
			p.processImport()
			break
		case checkItemType(item, lexer.ItemPrivate):
			p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemPrivate])
			private = true
			item = p.Lex.NextItem()
			if checkItemType(item, lexer.ItemGlobal) {
				p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemGlobal])
				global = true
				item = p.Lex.NextItem()
			}
			if checkItemType(item, lexer.ItemRule) {
				p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemRule])
				p.processRule(global, private)
				private = false
				global = false
			}
			break
		case checkItemType(item, lexer.ItemGlobal):
			p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemGlobal])
			global = true
			item = p.Lex.NextItem()
			if checkItemType(item, lexer.ItemPrivate) {
				p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemPrivate])
				private = true
				item = p.Lex.NextItem()
			}
			if checkItemType(item, lexer.ItemRule) {
				p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemRule])
				p.processRule(global, private)
				private = false
				global = false
			}
			break
		case checkItemType(item, lexer.ItemRule):
			p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemRule])
			p.processRule(global, private)
			break
		case checkItemType(item, lexer.Eof):
			break
		}
	}
}

func (p *Parser) processImport() {
	item := p.nextItem()
	if checkItemType(item, lexer.ItemString) {
		if !p.addImport(item) {
			p.warnf("WARN line %d: Module %s already imported.\n", item.Line, item.Val)
		}
	} else {
		p.errorf("ERROR: line %d expected %s and found.\n", item.Line, lexer.ItemNames[lexer.ItemString], lexer.ItemNames[item.Typ])
	}
}

func (p *Parser) processRule(global, private bool) {
	item := p.nextItem()
	if checkItemType(item, lexer.ItemIdentifier) {
		p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemIdentifier])
		ruleName := item.Val
		if !p.ruleAlreadyImported(ruleName) {
			newRule := RuleDef{
				Name:    ruleName,
				Global:  global,
				Private: private,
			}
			// fmt.Printf("RULE: %s\n", ruleName)
			item = p.nextItem()
			if checkItemType(p.LastItem, lexer.ItemColon) { // Tags comming
				newRule.Tags = p.processTags()
				p.log.Debugf("Tags %v", newRule.Tags)
			}
			if checkItemType(p.LastItem, lexer.ItemLeftCurly) {
				item = p.nextItem()
				if checkItemType(p.LastItem, lexer.ItemMeta) { // Meta comming
					p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemMeta])
					item = p.nextItem()
					if checkItemType(item, lexer.ItemColon) {
						newRule.Meta = p.processMeta()
						p.log.Debugf("Meta %v", newRule.Meta)
					} else {
						p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemColon], lexer.ItemNames[item.Typ])
					}
				}
				if checkItemType(p.LastItem, lexer.ItemStrings) { // Strings comming
					p.log.Debugf("Keyword found: %s\n", lexer.ItemNames[lexer.ItemStrings])
					item = p.nextItem()
					if checkItemType(item, lexer.ItemColon) {
						newRule.Strings = p.processStrings()
						p.log.Debugf("Strings %v", newRule.Strings)
					} else {
						p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemColon], lexer.ItemNames[item.Typ])
					}
				}
				if checkItemType(p.LastItem, lexer.ItemCondition) {
					item = p.nextItem()
					if checkItemType(item, lexer.ItemColon) {
						newRule.Condition = p.processCondition()
						p.addRule(newRule)
					} else {
						p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemColon], lexer.ItemNames[item.Typ])
					}
				} else {
					p.errorf("ERROR: line %d expected %s or %s or %s found %s", item.Line, lexer.ItemNames[lexer.ItemMeta], lexer.ItemNames[lexer.ItemStrings], lexer.ItemNames[lexer.ItemCondition], lexer.ItemNames[item.Typ])
				}
			} else {
				p.errorf("ERROR: line %d expected %s found %s.", item.Line, lexer.ItemNames[lexer.ItemLeftCurly], lexer.ItemNames[item.Typ])
			}
		} else {
			p.errorf("ERROR: line %d rule %s alredy defined.", item.Line, item.Val)
		}
	} else {
		p.errorf("ERROR: line %d expected %s found %s.", item.Line, lexer.ItemNames[lexer.ItemIdentifier], lexer.ItemNames[item.Typ])
	}
}

func (p *Parser) processTags() []string {
	var tags []string
	item := p.nextItem()
	for !checkItemType(item, lexer.ItemLeftCurly) {
		if checkItemType(item, lexer.ItemIdentifier) {
			tags = append(tags, item.Val)
		} else {
			p.errorf("ERROR: line %d expecting %s or %s and found %s.", item.Line, lexer.ItemNames[lexer.ItemIdentifier], lexer.ItemNames[lexer.ItemLeftCurly], lexer.ItemNames[item.Typ])
		}
		item = p.nextItem()
	}
	return tags
}

func (p *Parser) processMeta() map[string]string {
	var key, value lexer.Item
	meta := make(map[string]string)
	item := p.nextItem()
	for !checkItemType(item, lexer.ItemStrings) && !checkItemType(item, lexer.ItemCondition) {
		key = item
		if checkItemType(item, lexer.ItemIdentifier) {
			item = p.nextItem()
			if checkItemType(item, lexer.ItemEqual) {
				item = p.nextItem()
				value = item
				if checkItemType(item, lexer.ItemString) ||
					checkItemType(item, lexer.ItemNumber) ||
					checkItemType(item, lexer.ItemTrue) || // Yara allows boolans as values
					checkItemType(item, lexer.ItemFalse) {
					meta[key.Val] = value.Val
				} else {
					p.errorf("ERROR: line %d expected %s or %s found %s", item.Line, lexer.ItemNames[lexer.ItemString], lexer.ItemNames[lexer.ItemNumber], lexer.ItemNames[item.Typ])
				}
			} else {
				p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemEqual], lexer.ItemNames[item.Typ])
			}
		} else {
			p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemIdentifier], lexer.ItemNames[item.Typ])
		}
		item = p.nextItem()
	}
	return meta
}

func (p *Parser) processStrings() []StringDef {
	var key lexer.Item
	var value string
	var strings []StringDef

	item := p.nextItem()
	for !checkItemType(item, lexer.ItemCondition) {
		key = item
		if checkItemType(item, lexer.ItemVariable) {
			item = p.nextItem()
			if checkItemType(item, lexer.ItemEqual) {
				item = p.nextItem()
				if checkItemType(item, lexer.ItemString) ||
					checkItemType(item, lexer.ItemTrue) || // Yara allows boolans as values
					checkItemType(item, lexer.ItemFalse) {
					value, mods := p.processStringModifiers()
					strings = append(strings, StringDef{Name: key.Val, Value: value, Modifiers: mods})
				} else if checkItemType(item, lexer.ItemRegex) {
					value, mods := p.processStringModifiers()
					strings = append(strings, StringDef{Name: key.Val, Value: value, Modifiers: mods})
				} else if checkItemType(item, lexer.ItemLeftCurly) {
					value = p.processHexValues()
					strings = append(strings, StringDef{Name: key.Val, Value: value})
				}
			} else {
				p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemEqual], lexer.ItemNames[item.Typ])
			}
		} else {
			p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemVariable], lexer.ItemNames[item.Typ])
		}
		item = p.nextItem()
	}
	return strings
}

func (p *Parser) processStringModifiers() (string, []string) {
	value := p.LastItem
	var mods []string
	// item = p.nextItem()
	item := p.peek()
	if isStringModifier(item) {
		item = p.nextItem()
		for isStringModifier(item) {
			mods = append(mods, item.Val)
			item = p.nextItem()
		}
		p.backup()
	}
	return value.Val, mods
}

func (p *Parser) processHexValues() string {
	value := "{"
	item := p.nextItem()
	for !checkItemType(item, lexer.ItemRightCurly) {
		switch {
		case checkItemType(item, lexer.ItemIdentifier): // Hex value
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemNumber): // Number
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemWildCard): // ? Wildcard
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemLeftSqrt):
			value = value + p.processHexRange()
		case checkItemType(item, lexer.ItemLeftBra):
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
	if checkItemType(item, lexer.ItemNumber) {
		numA := item.Val
		nA, err := strconv.Atoi(numA)
		if err == nil {
			if nA < 0 {
				p.errorf("ERROR: line %d expected >= 0 integer found %s", item.Line, lexer.ItemNames[lexer.ItemRightSqrt], lexer.ItemNames[item.Typ])
			}
		} else {
			p.errorf("ERROR: line %d unable to convert %s", item.Line, lexer.ItemNames[item.Typ])
		}
		item = p.nextItem()
		if checkItemType(item, lexer.ItemDash) {
			dash := item.Val
			item = p.nextItem()
			if checkItemType(item, lexer.ItemNumber) {
				numB := item.Val
				nB, err := strconv.Atoi(numB)
				if err == nil {
					if nB < nA {
						p.errorf("ERROR: line %d expected %s to be grather than %s", item.Line, numB, numA)
					}
				} else {
					p.errorf("ERROR: line %d unable to convert %s", item.Line, lexer.ItemNames[item.Typ])
				}
				item = p.nextItem()
				if checkItemType(item, lexer.ItemRightSqrt) {
					value = value + numA + dash + numB + item.Val
					return value
				} else {
					p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemRightSqrt], lexer.ItemNames[item.Typ])
				}
			} else if checkItemType(item, lexer.ItemRightSqrt) {
				value = value + numA + dash + item.Val
				return value
			} else {
				p.errorf("ERROR: line %d expected %s or  found %s", item.Line, lexer.ItemNames[lexer.ItemNumber], lexer.ItemNames[lexer.ItemRightSqrt], lexer.ItemNames[item.Typ])
			}
		} else if checkItemType(item, lexer.ItemRightSqrt) {
			value = value + numA + item.Val
			return value
		} else {
			p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemDash], lexer.ItemNames[item.Typ])
		}
	} else if checkItemType(item, lexer.ItemDash) {
		dash := item.Val
		item = p.nextItem()
		if checkItemType(item, lexer.ItemRightSqrt) {
			value = value + dash + item.Val
			return value
		} else {
			p.errorf("ERROR: line %d expected %s found %s", item.Line, lexer.ItemNames[lexer.ItemRightSqrt], lexer.ItemNames[item.Typ])
		}
	} else {
		p.errorf("ERROR: line %d expected %s or %s found %s", item.Line, lexer.ItemNames[lexer.ItemNumber], lexer.ItemNames[lexer.ItemDash], lexer.ItemNames[item.Typ])
	}
	value = value + item.Val
	return value
}

func (p *Parser) preocessHexOption() string {
	value := "("
	item := p.nextItem()
	for !checkItemType(item, lexer.ItemRightBra) {
		switch {
		case checkItemType(item, lexer.ItemIdentifier): // Hex value
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemNumber): // Number
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemWildCard): // ? Wildcard
			value = value + item.Val
			break
		case checkItemType(item, lexer.ItemPipe): // ? Wildcard
			value = value + item.Val
			break
		}
		item = p.nextItem()
	}
	value = value + item.Val
	return value
}

func (p *Parser) processCondition() string {
	value := ""
	item := p.nextItem()
	for !checkItemType(item, lexer.ItemRightCurly) {
		value = value + item.Val + " "
		item = p.nextItem()
	}
	return strings.Trim(value, " ")
}
