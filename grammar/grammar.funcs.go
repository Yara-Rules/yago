package grammar

import (
	"fmt"

	"dev.jau.me/YaGo2/lexic"
)

// String returns a string representation of the Parser
func (p *Parser) String() string {
	r := ""
	for _, imp := range p.Imports {
		r += fmt.Sprintf("import \"%s\"\n", imp)
	}

	if len(p.Imports) > 0 {
		r += fmt.Sprintf("\n")
	}

	for _, rule := range p.Rules {
		if rule.Private {
			r += fmt.Sprintf("private ")
		}
		if rule.Global {
			r += fmt.Sprintf("global ")
		}

		r += fmt.Sprintf("rule %s", rule.Name)
		if len(rule.Tags) > 0 {
			r += fmt.Sprintf(" :")
			for i := 0; i < len(rule.Tags)-1; i++ {
				r += fmt.Sprintf(" %s", rule.Tags[i])
			}
			r += fmt.Sprintf("%s {\n", rule.Tags[len(rule.Tags)-1])
		} else {
			r += fmt.Sprintf(" {\n")
		}
		if len(rule.Meta) > 0 {
			r += fmt.Sprintf("\tmeta:\n")
			for k, v := range rule.Meta {
				r += fmt.Sprintf("\t\t%s = \"%s\"\n", k, v)
			}
		}
		if len(rule.Strings) > 0 {
			r += fmt.Sprintf("\tstrings:\n")
			for _, str := range rule.Strings {
				if str.Typ == StringString {
					r += fmt.Sprintf("\t\t%s = \"%s\"", str.Name, str.Value)
					if len(str.Modifiers) > 0 {
						for _, m := range str.Modifiers {
							r += fmt.Sprintf(" %s", m)
						}
					}
				} else if str.Typ == StringRegex {
					r += fmt.Sprintf("\t\t%s = %s", str.Name, str.Value)
				} else if str.Typ == StringHex {
					r += fmt.Sprintf("\t\t%s = %s", str.Name, str.Value)
				}
				r += fmt.Sprintf("\n")
			}
		}
		r += fmt.Sprintf("\tcondition:\n")
		r += fmt.Sprintf("\t\t%s\n", rule.Condition)
		r += fmt.Sprintf("}\n\n")
	}
	return r
}

func (p *Parser) addImport(item lexic.Item) bool {
	if !p.moduleAlreadyImported(item) {
		p.Imports = append(p.Imports, item.GetValue())
		return true
	}
	return false
}

func (p *Parser) moduleAlreadyImported(item lexic.Item) bool {
	for _, v := range p.Imports {
		if v == item.GetValue() {
			return true
		}
	}
	return false
}

func (p *Parser) ruleAlreadyImported(ruleName string) bool {
	for _, v := range p.Rules {
		if v.Name == ruleName {
			return true
		}
	}
	return false
}

func (p *Parser) addRule(rule RuleDef) {
	p.Rules = append(p.Rules, rule)
}
