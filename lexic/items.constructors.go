package lexic

import "strconv"

// NewItemComment comment item constructor
func NewItemComment(value string, pos pos, line int) Item {
	return ItemComment{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemComment"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemEOF EOF item constructor
func NewItemEOF(value string, pos pos, line int) Item {
	return ItemEOF{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemEOF"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemIdentifier identifier item constructor
func NewItemIdentifier(value string, pos pos, line int) Item {
	return ItemIdentifier{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemIdentifier"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemString string item constructor
func NewItemString(value string, pos pos, line int) Item {
	return ItemString{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemString"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemRegex regex item constructor
func NewItemRegex(value string, pos pos, line int) Item {
	return ItemRegex{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemRegex"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemIntNumber int number item constructor
func NewItemIntNumber(value string, pos pos, line int) Item {
	val, err := strconv.Atoi(value)
	checkErr(err)
	return ItemIntNumber{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemIntNumber"],
			Pos:   pos,
			Line:  line,
		},
		val,
	}
}

// NewItemVariable strings' variable item constructor
func NewItemVariable(value string, pos pos, line int) Item {
	return ItemVariable{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemVariable"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemColon colon item constructor
func NewItemColon(value string, pos pos, line int) Item {
	return ItemColon{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemColon"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemEqual equal item constructor
func NewItemEqual(value string, pos pos, line int) Item {
	return ItemEqual{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemEqual"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemOCurly open curly item constructor
func NewItemOCurly(value string, pos pos, line int) Item {
	return ItemOCurly{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemOCurly"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemCCurly close curly item constructor
func NewItemCCurly(value string, pos pos, line int) Item {
	return ItemCCurly{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemCCurly"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemOSqrt open square bracket item constructor
func NewItemOSqrt(value string, pos pos, line int) Item {
	return ItemOSqrt{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemOSqrt"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemCSqrt close square bracket item constructor
func NewItemCSqrt(value string, pos pos, line int) Item {
	return ItemCSqrt{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemCSqrt"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemOBracket open bracket item constructor
func NewItemOBracket(value string, pos pos, line int) Item {
	return ItemOBracket{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemOBracket"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemCBracket close bracket item constructor
func NewItemCBracket(value string, pos pos, line int) Item {
	return ItemCBracket{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemCBracket"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemPipe pipe item constructor
func NewItemPipe(value string, pos pos, line int) Item {
	return ItemPipe{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemPipe"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemSpace space item constructor
func NewItemSpace(value string, pos pos, line int) Item {
	return ItemSpace{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemSpace"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemQMark question mark item constructor
func NewItemQMark(value string, pos pos, line int) Item {
	return ItemQMark{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemQMark"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemDash dash item constructor
func NewItemDash(value string, pos pos, line int) Item {
	return ItemDash{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemDash"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemPlus plus item constructor
func NewItemPlus(value string, pos pos, line int) Item {
	return ItemPlus{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemPlus"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemHash hash item constructor
func NewItemHash(value string, pos pos, line int) Item {
	return ItemHash{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemHash"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemDot dot item constructor
func NewItemDot(value string, pos pos, line int) Item {
	return ItemDot{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemDot"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemDotDot dot dot (:) item constructor
func NewItemDotDot(value string, pos pos, line int) Item {
	return ItemDotDot{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemDotDot"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemCaret caret item constructor
func NewItemCaret(value string, pos pos, line int) Item {
	return ItemCaret{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemCaret"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemStar star item constructor
func NewItemStar(value string, pos pos, line int) Item {
	return ItemStar{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemStar"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemSlash slash item constructor
func NewItemSlash(value string, pos pos, line int) Item {
	return ItemSlash{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemSlash"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemComma comma item constructor
func NewItemComma(value string, pos pos, line int) Item {
	return ItemComma{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemComma"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemGrater grater than item constructor
func NewItemGrater(value string, pos pos, line int) Item {
	return ItemGrater{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemGrater"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemLess less than item constructor
func NewItemLess(value string, pos pos, line int) Item {
	return ItemLess{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemLess"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemGraterEqual grater equal item constructor
func NewItemGraterEqual(value string, pos pos, line int) Item {
	return ItemGraterEqual{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemGraterEqual"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemLessEqual less equal item constructor
func NewItemLessEqual(value string, pos pos, line int) Item {
	return ItemLessEqual{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemLessEqual"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemEqualEqual equal equal (==) item constructor
func NewItemEqualEqual(value string, pos pos, line int) Item {
	return ItemEqualEqual{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemEqualEqual"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemNot not item constructor
func NewItemNot(value string, pos pos, line int) Item {
	return ItemNot{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemNot"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemNotEqual not equal (!=) item constructor
func NewItemNotEqual(value string, pos pos, line int) Item {
	return ItemNotEqual{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemNotEqual"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemPercent percentage item constructor
func NewItemPercent(value string, pos pos, line int) Item {
	return ItemPercent{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemPercent"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemAnd and item constructor
func NewItemAnd(value string, pos pos, line int) Item {
	return ItemAnd{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemAnd"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemLeftShift left shift item constructor
func NewItemLeftShift(value string, pos pos, line int) Item {
	return ItemLeftShift{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemLeftShift"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemRightShift right shift item constructor
func NewItemRightShift(value string, pos pos, line int) Item {
	return ItemRightShift{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemRightShift"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemBitNot bitwise not item constructor
func NewItemBitNot(value string, pos pos, line int) Item {
	return ItemBitNot{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemBitNot"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemAt at (@) item constructor
func NewItemAt(value string, pos pos, line int) Item {
	return ItemAt{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemAt"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWAll all keyword item constructor
func NewItemKWAll(value string, pos pos, line int) Item {
	return ItemKWAll{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWAll"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWAnd and keyword item constructor
func NewItemKWAnd(value string, pos pos, line int) Item {
	return ItemKWAnd{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWAnd"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWAny any keyword item constructor
func NewItemKWAny(value string, pos pos, line int) Item {
	return ItemKWAny{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWAny"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWAscii ascii keyword item constructor
func NewItemKWAscii(value string, pos pos, line int) Item {
	return ItemKWAscii{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWAscii"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWAt at keyword item constructor
func NewItemKWAt(value string, pos pos, line int) Item {
	return ItemKWAt{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWAt"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWCondition condition keyword item constructor
func NewItemKWCondition(value string, pos pos, line int) Item {
	return ItemKWCondition{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWCondition"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWContains contains keyword item constructor
func NewItemKWContains(value string, pos pos, line int) Item {
	return ItemKWContains{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWContains"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWEntrypoint entrypoint keyword item constructor
func NewItemKWEntrypoint(value string, pos pos, line int) Item {
	return ItemKWEntrypoint{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWEntrypoint"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWFalse false keyword item constructor
func NewItemKWFalse(value string, pos pos, line int) Item {
	return ItemKWFalse{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWFalse"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWFilesize filesize keyword item constructor
func NewItemKWFilesize(value string, pos pos, line int) Item {
	return ItemKWFilesize{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWFilesize"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWFullword fullword keyword item constructor
func NewItemKWFullword(value string, pos pos, line int) Item {
	return ItemKWFullword{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWFullword"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWFor for keyword item constructor
func NewItemKWFor(value string, pos pos, line int) Item {
	return ItemKWFor{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWFor"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWGlobal global keyword item constructor
func NewItemKWGlobal(value string, pos pos, line int) Item {
	return ItemKWGlobal{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWGlobal"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWIn in keyword item constructor
func NewItemKWIn(value string, pos pos, line int) Item {
	return ItemKWIn{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWIn"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWImport import keyword item constructor
func NewItemKWImport(value string, pos pos, line int) Item {
	return ItemKWImport{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWImport"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInclude include keyword item constructor
func NewItemKWInclude(value string, pos pos, line int) Item {
	return ItemKWInclude{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInclude"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt8 int8 keyword item constructor
func NewItemKWInt8(value string, pos pos, line int) Item {
	return ItemKWInt8{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt8"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt16 int16 keyword item constructor
func NewItemKWInt16(value string, pos pos, line int) Item {
	return ItemKWInt16{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt16"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt32 int32 keyword item constructor
func NewItemKWInt32(value string, pos pos, line int) Item {
	return ItemKWInt32{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt32"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt8be int8be keyword item constructor
func NewItemKWInt8be(value string, pos pos, line int) Item {
	return ItemKWInt8be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt8be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt16be int16be keyword item constructor
func NewItemKWInt16be(value string, pos pos, line int) Item {
	return ItemKWInt16be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt16be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWInt32be int32be keyword item constructor
func NewItemKWInt32be(value string, pos pos, line int) Item {
	return ItemKWInt32be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWInt32be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWMatches matches keyword item constructor
func NewItemKWMatches(value string, pos pos, line int) Item {
	return ItemKWMatches{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWMatches"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWMeta meta keyword item constructor
func NewItemKWMeta(value string, pos pos, line int) Item {
	return ItemKWMeta{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWMeta"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWNocase nocase keyword item constructor
func NewItemKWNocase(value string, pos pos, line int) Item {
	return ItemKWNocase{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWNocase"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWNot not keyword item constructor
func NewItemKWNot(value string, pos pos, line int) Item {
	return ItemKWNot{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWNot"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWOr or keyword item constructor
func NewItemKWOr(value string, pos pos, line int) Item {
	return ItemKWOr{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWOr"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWOf of keyword item constructor
func NewItemKWOf(value string, pos pos, line int) Item {
	return ItemKWOf{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWOf"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWPrivate private keyword item constructor
func NewItemKWPrivate(value string, pos pos, line int) Item {
	return ItemKWPrivate{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWPrivate"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWRule rule keyword item constructor
func NewItemKWRule(value string, pos pos, line int) Item {
	return ItemKWRule{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWRule"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWStrings strings keyword item constructor
func NewItemKWStrings(value string, pos pos, line int) Item {
	return ItemKWStrings{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWStrings"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWThem them keyword item constructor
func NewItemKWThem(value string, pos pos, line int) Item {
	return ItemKWThem{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWThem"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWTrue true keyword item constructor
func NewItemKWTrue(value string, pos pos, line int) Item {
	return ItemKWTrue{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWTrue"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint8 uint8 keyword item constructor
func NewItemKWUint8(value string, pos pos, line int) Item {
	return ItemKWUint8{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint8"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint16 uint16 keyword item constructor
func NewItemKWUint16(value string, pos pos, line int) Item {
	return ItemKWUint16{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint16"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint32 uint32 keyword item constructor
func NewItemKWUint32(value string, pos pos, line int) Item {
	return ItemKWUint32{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint32"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint8be uint8be keyword item constructor
func NewItemKWUint8be(value string, pos pos, line int) Item {
	return ItemKWUint8be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint8be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint16be uint16be keyword item constructor
func NewItemKWUint16be(value string, pos pos, line int) Item {
	return ItemKWUint16be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint16be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWUint32be uint32be keyword item constructor
func NewItemKWUint32be(value string, pos pos, line int) Item {
	return ItemKWUint32be{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWUint32be"],
			Pos:   pos,
			Line:  line,
		},
	}
}

// NewItemKWWide wide keyword item constructor
func NewItemKWWide(value string, pos pos, line int) Item {
	return ItemKWWide{
		ItemYaGo{
			Value: value,
			Typ:   ItemType["ItemKWWide"],
			Pos:   pos,
			Line:  line,
		},
	}
}
