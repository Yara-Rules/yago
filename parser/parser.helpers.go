package parser

import "github.com/Yara-Rules/yago/lexer"

func checkItemType(a lexer.Item, b lexer.ItemType) bool {
	return a.Typ == b
}

func isKeyword(a lexer.Item) bool {
	return a.Typ > lexer.ItemKeyword
}

func isStringModifier(a lexer.Item) bool {
	return a.Typ == lexer.ItemNocase || a.Typ == lexer.ItemAscii ||
		a.Typ == lexer.ItemWide || a.Typ == lexer.ItemFullword
}

func (p *Parser) alreadyImported(item lexer.Item) bool {
	for _, v := range p.Imports {
		if v == item.Val {
			return true
		}
	}
	return false
}

func (p *Parser) addImport(item lexer.Item) bool {
	if !p.alreadyImported(item) {
		p.Imports = append(p.Imports, item.Val)
		return true
	}
	return false
}

func (p *Parser) addRule(rule RuleDef) {
	p.Rules = append(p.Rules, rule)
}

func (p *Parser) ruleAlreadyImported(ruleName string) bool {
	for _, v := range p.Rules {
		if v.Name == ruleName {
			return true
		}
	}
	return false
}
