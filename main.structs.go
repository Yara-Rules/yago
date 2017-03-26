package main

import (
	"fmt"

	"dev.jau.me/YaGo2/yago"
)

type jsonCloak struct {
	Ruleset []*yago.Parser
}

type unify struct {
	imports []string
	rules   []yago.RuleDef
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

func (u *unify) addRule(rule yago.RuleDef) {
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
				if str.Typ == yago.StringString {
					r += fmt.Sprintf("\t\t%s = \"%s\"", str.Name, str.Value)
					if len(str.Modifiers) > 0 {
						for _, m := range str.Modifiers {
							r += fmt.Sprintf(" %s", m)
						}
					}
				} else if str.Typ == yago.StringRegex {
					r += fmt.Sprintf("\t\t%s = %s", str.Name, str.Value)
				} else if str.Typ == yago.StringHex {
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
