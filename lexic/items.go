package lexic

import (
	"fmt"
	"os"
)

// Item interface for all Item in YaGo
type Item interface {
	String()
	GetValue() string
	GetType() string
	GetLine() int
	GetPos() int
}

type pos int

// ItemYaGo basic information for each item
type ItemYaGo struct {
	Value string
	Typ   string
	Pos   pos
	Line  int
}

// String prints on stdout the item value
func (it ItemYaGo) String() {
	os.Stdout.WriteString(fmt.Sprintf("%s", it.Value))
}

// GetValue return the item value
func (it ItemYaGo) GetValue() string {
	return string(it.Value)
}

// GetType returns the item type
func (it ItemYaGo) GetType() string {
	return string(it.Typ)
}

// GetLine returns where the item was found
func (it ItemYaGo) GetLine() int {
	return it.Line
}

// GetPos returns the item position inside the input
func (it ItemYaGo) GetPos() int {
	return int(it.Pos)
}

// ItemComment represents a comment Item
type ItemComment struct {
	ItemYaGo
}

// ItemEOF represents a EOF Item
type ItemEOF struct {
	ItemYaGo
}

// ItemIdentifier represents an identifier Item
type ItemIdentifier struct {
	ItemYaGo
}

// ItemString represents a string Item
type ItemString struct {
	ItemYaGo
}

// ItemRegex represents a identifier Item
type ItemRegex struct {
	ItemYaGo
}

// ItemIntNumber represents an integer number Item
type ItemIntNumber struct {
	ItemYaGo
	Value int
}

// String prints on stdout the item value
func (it ItemIntNumber) String() {
	os.Stdout.WriteString(fmt.Sprintf("%d", it.Value))
}

// ItemVariable represents a variable Item
type ItemVariable struct {
	ItemYaGo
}

// ItemColon represents a colon Item
type ItemColon struct {
	ItemYaGo
}

// ItemEqual represents a equal Item
type ItemEqual struct {
	ItemYaGo
}

// ItemOCurly represents an open curly bracket Item
type ItemOCurly struct {
	ItemYaGo
}

// ItemCCurly represents a close curly bracket Item
type ItemCCurly struct {
	ItemYaGo
}

// ItemOSqrt represents an open square bracket Item
type ItemOSqrt struct {
	ItemYaGo
}

// ItemCSqrt represents a close square bracket Item
type ItemCSqrt struct {
	ItemYaGo
}

// ItemOBracket represents an open bracket Item
type ItemOBracket struct {
	ItemYaGo
}

// ItemCBracket represents a close bracket Item
type ItemCBracket struct {
	ItemYaGo
}

// ItemPipe represents a pipe Item
type ItemPipe struct {
	ItemYaGo
}

// ItemSpace represents a space Item
type ItemSpace struct {
	ItemYaGo
}

// ItemQMark represents a question mark Item
type ItemQMark struct {
	ItemYaGo
}

// ItemDash represents a dash Item
type ItemDash struct {
	ItemYaGo
}

// ItemPlus represents a plus Item
type ItemPlus struct {
	ItemYaGo
}

// ItemHash represents a hash Item
type ItemHash struct {
	ItemYaGo
}

// ItemDot represents a dot Item
type ItemDot struct {
	ItemYaGo
}

// ItemDotDot represents a dotdot Item
type ItemDotDot struct {
	ItemYaGo
}

// ItemCaret represents a caret Item
type ItemCaret struct {
	ItemYaGo
}

// ItemStar represents a star Item
type ItemStar struct {
	ItemYaGo
}

// ItemSlash represents a slash Item
type ItemSlash struct {
	ItemYaGo
}

// ItemComma represents a comma Item
type ItemComma struct {
	ItemYaGo
}

// ItemGrater represents a grater than Item
type ItemGrater struct {
	ItemYaGo
}

// ItemLess represents a less than Item
type ItemLess struct {
	ItemYaGo
}

// ItemGraterEqual represents a grater equal Item
type ItemGraterEqual struct {
	ItemYaGo
}

// ItemLessEqual represents a less equal Item
type ItemLessEqual struct {
	ItemYaGo
}

// ItemEqualEqual represents a equal equal Item
type ItemEqualEqual struct {
	ItemYaGo
}

// ItemNotl represents a not Item
type ItemNot struct {
	ItemYaGo
}

// ItemNotlEqual represents a not equal Item
type ItemNotEqual struct {
	ItemYaGo
}

// ItemPercent represents a percentage Item
type ItemPercent struct {
	ItemYaGo
}

// ItemAnd represents a and Item
type ItemAnd struct {
	ItemYaGo
}

// ItemLeftShift represents a left shift Item
type ItemLeftShift struct {
	ItemYaGo
}

// ItemRightShift right shift
type ItemRightShift struct {
	ItemYaGo
}

// ItemBitNot represents a bitwise not Item
type ItemBitNot struct {
	ItemYaGo
}

// ItemAt represents an at (@) Item
type ItemAt struct {
	ItemYaGo
}

// ItemKWAll represents an all keyword Item
type ItemKWAll struct {
	ItemYaGo
}

// ItemKWAnd represents an and keyword Item
type ItemKWAnd struct {
	ItemYaGo
}

// ItemKWAny represents an any keyword Item
type ItemKWAny struct {
	ItemYaGo
}

// ItemKWAscii represents an ascii keyword Item
type ItemKWAscii struct {
	ItemYaGo
}

// ItemKWAt represents an at keyword Item
type ItemKWAt struct {
	ItemYaGo
}

// ItemKWCondition represents a condition keyword Item
type ItemKWCondition struct {
	ItemYaGo
}

// ItemKWContains represents a contains keyword Item
type ItemKWContains struct {
	ItemYaGo
}

// ItemKWEntrypoint represents a entrypoint keyword Item
type ItemKWEntrypoint struct {
	ItemYaGo
}

// ItemKWFalse represents a false keyword Item
type ItemKWFalse struct {
	ItemYaGo
}

// ItemKWFilesize represents a filesize keyword Item
type ItemKWFilesize struct {
	ItemYaGo
}

// ItemKWFullword represents a fullword keyword Item
type ItemKWFullword struct {
	ItemYaGo
}

// ItemKWFor represents a for keyword Item
type ItemKWFor struct {
	ItemYaGo
}

// ItemKWGlobal represents a global keyword Item
type ItemKWGlobal struct {
	ItemYaGo
}

// ItemKWIn represents an in keyword Item
type ItemKWIn struct {
	ItemYaGo
}

// ItemKWImport represents an import keyword Item
type ItemKWImport struct {
	ItemYaGo
}

// ItemKWInclude represents an include keyword Item
type ItemKWInclude struct {
	ItemYaGo
}

// ItemKWInt8 represents an int8 keyword Item
type ItemKWInt8 struct {
	ItemYaGo
}

// ItemKWInt16 represents an int16 keyword Item
type ItemKWInt16 struct {
	ItemYaGo
}

// ItemKWInt32 represents an int32 keyword Item
type ItemKWInt32 struct {
	ItemYaGo
}

// ItemKWInt8be represents an int8be keyword Item
type ItemKWInt8be struct {
	ItemYaGo
}

// ItemKWInt16be represents an int16be keyword Item
type ItemKWInt16be struct {
	ItemYaGo
}

// ItemKWInt32be represents an int32be keyword Item
type ItemKWInt32be struct {
	ItemYaGo
}

// ItemKWMatches represents a matches keyword Item
type ItemKWMatches struct {
	ItemYaGo
}

// ItemKWMeta represents a meta keyword Item
type ItemKWMeta struct {
	ItemYaGo
}

// ItemKWNocase represents a nocase keyword Item
type ItemKWNocase struct {
	ItemYaGo
}

// ItemKWNot represents a not keyword Item
type ItemKWNot struct {
	ItemYaGo
}

// ItemKWOr represents an or keyword Item
type ItemKWOr struct {
	ItemYaGo
}

// ItemKWOf represents an of keyword Item
type ItemKWOf struct {
	ItemYaGo
}

// ItemKWPrivate represents a private keyword Item
type ItemKWPrivate struct {
	ItemYaGo
}

// ItemKWRule represents a rule keyword Item
type ItemKWRule struct {
	ItemYaGo
}

// ItemKWStrings represents a strings keyword Item
type ItemKWStrings struct {
	ItemYaGo
}

// ItemKWThem represents a them keyword Item
type ItemKWThem struct {
	ItemYaGo
}

// ItemKWTrue represents a true keyword Item
type ItemKWTrue struct {
	ItemYaGo
}

// ItemKWUint8 represents an uint8 keyword Item
type ItemKWUint8 struct {
	ItemYaGo
}

// ItemKWUint16 represents an uint16 keyword Item
type ItemKWUint16 struct {
	ItemYaGo
}

// ItemKWUint32 represents an uint32 keyword Item
type ItemKWUint32 struct {
	ItemYaGo
}

// ItemKWUint8be represents an uint8be keyword Item
type ItemKWUint8be struct {
	ItemYaGo
}

// ItemKWUint16be represents an uint16be keyword Item
type ItemKWUint16be struct {
	ItemYaGo
}

// ItemKWUint32be represents an uint32be keyword Item
type ItemKWUint32be struct {
	ItemYaGo
}

// ItemKWWide represents a wide keyword Item
type ItemKWWide struct {
	ItemYaGo
}
