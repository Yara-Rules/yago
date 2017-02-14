package lexer

type Item struct {
	Typ  ItemType
	Pos  Pos
	Val  string
	Line int
}

type Pos int

type ItemType int

const Eof = -1

var ItemNames = map[ItemType]string{
	ItemNumber:      "__NUMBER__",
	ItemComment:     "__COMMENT__",
	ItemIdentifier:  "__IDENTIFIER__",
	ItemString:      "__QUOTED_STRING__",
	ItemRegex:       "__REGEX__",
	ItemVariable:    "__VARIABLE__",
	ItemColon:       "__COLON__",
	ItemEqual:       "__EQUAL__",
	ItemLeftCurly:   "__LEFT_CURLY__",
	ItemRightCurly:  "__RIGHT_CURLY__",
	ItemLeftSqrt:    "__LEFT_SQUARE_BRACKET__",
	ItemRightSqrt:   "__RIGHT_SQUARE_BRACKET__",
	ItemLeftBra:     "__LEFT_BRACKET__",
	ItemRightBra:    "__RIGHT_BRACKET__",
	ItemPipe:        "__PIPE__",
	ItemSpace:       "__SPACE__",
	ItemWildCard:    "__QUESTION_MARK__",
	ItemDash:        "__DASH__",
	ItemHash:        "__HASH__",
	ItemDot:         "__DOT__",
	ItemCaret:       "__CARET__",
	ItemStar:        "__ASTERISC__",
	ItemPlusOrMinus: "__PLUS_OR_MINUS__",
	ItemComma:       "__COMMA__",
	ItemGratherThan: "__GRATHER_THAN__",
	ItemLessThan:    "__LESS_THAN__",
	ItemAtSym:       "__AT_SIGN__",
	ItemEOF:         "__EOF__",
	ItemAll:         "__ALL__",
	ItemAnd:         "__AND__",
	ItemAny:         "__ANY__",
	ItemAscii:       "__ASCII__",
	ItemAt:          "__AT__",
	ItemCondition:   "__CONDITION__",
	ItemContains:    "__CONTAINS__",
	ItemEntrypoint:  "__ENTRYPOINT__",
	ItemFalse:       "__FALSE__",
	ItemFilesize:    "__FILESIZE__",
	ItemFullword:    "__FULLWORD__",
	ItemFor:         "__FOR__",
	ItemGlobal:      "__GLOBAL__",
	ItemIn:          "__IN__",
	ItemImport:      "__IMPORT__",
	ItemInclude:     "__INCLUDE__",
	ItemInt8:        "__INT8__",
	ItemInt16:       "__INT16__",
	ItemInt32:       "__INT32__",
	ItemInt8be:      "__INT8BE__",
	ItemInt16be:     "__INT16BE__",
	ItemInt32be:     "__INT32BE__",
	ItemMatches:     "__MATCHES__",
	ItemMeta:        "__META__",
	ItemNocase:      "__NOCASE__",
	ItemNot:         "__NOT__",
	ItemOr:          "__OR__",
	ItemOf:          "__OF__",
	ItemPrivate:     "__PRIVATE__",
	ItemRule:        "__RULE__",
	ItemStrings:     "__STRINGS__",
	ItemThem:        "__THEN__",
	ItemTrue:        "__TRUE__",
	ItemUint8:       "__UINT8__",
	ItemUint16:      "__UINT16__",
	ItemUint32:      "__UINT32__",
	ItemUint8be:     "__UINT8BE__",
	ItemUint16be:    "__UINT16BE__",
	ItemUint32be:    "__UINT32BE__",
	ItemWide:        "__WIDEs__",
}
