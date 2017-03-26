package yago

import (
	"reflect"

	"github.com/Yara-Rules/yago/lexic"
)

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

func isStringModifier(a lexic.Item) bool {
	return a.GetType() == "__KW_NOCASE__" || a.GetType() == "__KW_ASCII__" ||
		a.GetType() == "__KW_WIDE__" || a.GetType() == "__KW_FULLWORD__"
}

func checkItemType(item lexic.Item, itemType string) bool {
	return lexic.ItemType[reflect.TypeOf(item).Name()] == itemType
}

func stringDefined(stringTable []string, item lexic.Item) bool {
	if item.GetValue() == "$" {
		return false
	}
	for _, str := range stringTable {
		if str == item.GetValue() {
			return true
		}
	}
	return false
}
