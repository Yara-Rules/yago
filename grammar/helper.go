package grammar

import (
	"reflect"

	"github.com/Yara-Rules/yago/lexic"
)

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
