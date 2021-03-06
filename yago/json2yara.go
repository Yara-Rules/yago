package yago

import (
	"fmt"

	"github.com/Yara-Rules/yago/grammar"
)

type jsonCloak struct {
	Ruleset []*grammar.Parser
}

type unify struct {
	imports []string
	rules   []grammar.RuleDef
}

func (u *unify) addImport(imp string) {
	exist := false
	for _, i := range u.imports {
		if imp == i {
			exist = true
			break
		}
	}
	if !exist {
		u.imports = append(u.imports, imp)
	}
}

func (u *unify) addRule(rule grammar.RuleDef) {
	exist := false
	for _, rs := range u.rules {
		if rule.Name == rs.Name {
			exist = true
			break
		}
	}
	if !exist {
		u.rules = append(u.rules, rule)
	}

}

func (u *unify) String() string {
	r := ""
	for _, imp := range u.imports {
		r += fmt.Sprintf("import \"%s\"\n", imp)
	}

	if len(u.imports) > 0 {
		r += fmt.Sprintf("\n")
	}

	for _, rule := range u.rules {
		if rule.Private {
			r += fmt.Sprintf("private ")
		}
		if rule.Global {
			r += fmt.Sprintf("global ")
		}

		r += fmt.Sprintf("rule %s", rule.Name)
		if len(rule.Tags) > 0 {
			r += " :"
			for _, tag := range rule.Tags {
				r += fmt.Sprintf(" %s", tag)
			}
			r += " {\n"
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
				if str.Typ == grammar.StringString {
					r += fmt.Sprintf("\t\t%s = \"%s\"", str.Name, str.Value)
					if len(str.Modifiers) > 0 {
						for _, m := range str.Modifiers {
							r += fmt.Sprintf(" %s", m)
						}
					}
				} else if str.Typ == grammar.StringRegex {
					r += fmt.Sprintf("\t\t%s = %s", str.Name, str.Value)
				} else if str.Typ == grammar.StringHex {
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
