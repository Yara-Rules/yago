package lexic

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type stateFn func(*Lexer) stateFn

// EOF represents a End Of File
const EOF = -1

// Lexer holds the state of the scanner.
type Lexer struct {
	Name    string    // the name of the input; used only for error reports
	Input   string    // the string being scanned
	State   stateFn   // the next lexing function to enter
	Pos     pos       // current position in the input
	Start   pos       // start position of this Item
	Width   pos       // width of last rune read from input
	LastPos pos       // position of most recent Item returned by nextItem
	Items   chan Item // channel of scanned Items
	Line    int       // 1+number of newlines seen
}

// Lex creates a new scanner for the input string.
func Lex(name, input string) *Lexer {
	l := &Lexer{
		Name:  name,
		Input: input,
		Items: make(chan Item),
		Line:  1,
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *Lexer) run() {
	for l.State = lexText; l.State != nil; {
		l.State = l.State(l)
	}
	close(l.Items)
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.Pos) >= len(l.Input) {
		l.Width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.Input[l.Pos:])

	l.Width = pos(w)
	l.Pos += l.Width
	if r == '\n' {
		l.Line++
	}
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.Pos -= l.Width
	// Correct newline count.
	if l.Width == 1 && l.Input[l.Pos] == '\n' {
		l.Line--
	}
}

// emit passes an Item back to the client.
func (l *Lexer) emit(itemName string) {
	var it Item

	switch itemName {
	case "ItemComment":
		it = NewItemComment(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemEOF":
		it = NewItemEOF(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemIdentifier":
		it = NewItemIdentifier(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemString":
		it = NewItemString(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemRegex":
		it = NewItemRegex(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemIntNumber":
		it = NewItemIntNumber(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemVariable":
		it = NewItemVariable(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemColon":
		it = NewItemColon(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemEqual":
		it = NewItemEqual(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemOCurly":
		it = NewItemOCurly(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemCCurly":
		it = NewItemCCurly(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemOSqrt":
		it = NewItemOSqrt(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemCSqrt":
		it = NewItemCSqrt(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemOBracket":
		it = NewItemOBracket(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemCBracket":
		it = NewItemCBracket(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemPipe":
		it = NewItemPipe(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemSpace":
		it = NewItemSpace(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemQMark":
		it = NewItemQMark(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemDash":
		it = NewItemDash(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemPlus":
		it = NewItemPlus(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemHash":
		it = NewItemHash(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemDot":
		it = NewItemDot(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemDotDot":
		it = NewItemDotDot(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemCaret":
		it = NewItemCaret(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemStar":
		it = NewItemStar(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemSlash":
		it = NewItemSlash(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemComma":
		it = NewItemComma(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemGrater":
		it = NewItemGrater(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemLess":
		it = NewItemLess(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemGraterEqual":
		it = NewItemGraterEqual(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemLessEqual":
		it = NewItemLessEqual(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemEqualEqual":
		it = NewItemEqualEqual(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemNotEqual":
		it = NewItemNotEqual(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemNot":
		it = NewItemNot(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemAnd":
		it = NewItemAnd(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemRightShift":
		it = NewItemRightShift(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemLeftShift":
		it = NewItemLeftShift(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemAt":
		it = NewItemAt(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemPercent":
		it = NewItemPercent(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWAll":
		it = NewItemKWAll(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWAnd":
		it = NewItemKWAnd(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWAny":
		it = NewItemKWAny(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWAscii":
		it = NewItemKWAscii(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWAt":
		it = NewItemKWAt(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWCondition":
		it = NewItemKWCondition(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWContains":
		it = NewItemKWContains(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWEntrypoint":
		it = NewItemKWEntrypoint(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWFalse":
		it = NewItemKWFalse(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWFilesize":
		it = NewItemKWFilesize(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWFullword":
		it = NewItemKWFullword(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWFor":
		it = NewItemKWFor(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWGlobal":
		it = NewItemKWGlobal(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWIn":
		it = NewItemKWIn(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWImport":
		it = NewItemKWImport(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInclude":
		it = NewItemKWInclude(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt8":
		it = NewItemKWInt8(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt16":
		it = NewItemKWInt16(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt32":
		it = NewItemKWInt32(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt8be":
		it = NewItemKWInt8be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt16be":
		it = NewItemKWInt16be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWInt32be":
		it = NewItemKWInt32be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWMatches":
		it = NewItemKWMatches(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWMeta":
		it = NewItemKWMeta(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWNocase":
		it = NewItemKWNocase(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWNot":
		it = NewItemKWNot(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWOr":
		it = NewItemKWOr(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWOf":
		it = NewItemKWOf(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWPrivate":
		it = NewItemKWPrivate(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWRule":
		it = NewItemKWRule(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWStrings":
		it = NewItemKWStrings(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWThem":
		it = NewItemKWThem(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWTrue":
		it = NewItemKWTrue(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint8":
		it = NewItemKWUint8(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint16":
		it = NewItemKWUint16(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint32":
		it = NewItemKWUint32(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint8be":
		it = NewItemKWUint8be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint16be":
		it = NewItemKWUint16be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWUint32be":
		it = NewItemKWUint32be(l.Input[l.Start:l.Pos], l.Start, l.Line)
	case "ItemKWWide":
		it = NewItemKWWide(l.Input[l.Start:l.Pos], l.Start, l.Line)
		// default:
		// 	panic(fmt.Sprintf("ERROR: %s is not a valid Item.", itemName))
	}
	l.Items <- it
	l.Start = l.Pos
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.Start = l.Pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *Lexer) scanned() string {
	return l.Input[l.Start:l.Pos]
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *Lexer) errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)

	err["lexical"] = "error"
	err["line"] = fmt.Sprintf("%d", l.Line)
	err["msg"] = msg

	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}

// NextItem returns the next Item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextItem() Item {
	item := <-l.Items
	return item
}

// lexText scans until an opening action delimiter, "{{".
func lexText(l *Lexer) stateFn {
	r := l.next()
	for !isEOF(r) {
		switch {
		case isEOF(r):
			l.emit("ItemEOF")
			return nil
		case isBlank(r):
			l.acceptRun(blankChars)
			l.ignore()
		case isSlash(r):
			return scanCommentOrRegex
		case isDollar(r):
			return scanVariable
		case isQuote(r):
			l.ignore() // Removes the first "
			return scanQuote
		case isEqual(r):
			return scanEqual
		case isAlphaNumeric(r):
			return scanKeyword
		case isColon(r):
			return scanColon
		case isOpenCurly(r):
			return scanOpenCurly
		case isCloseCurly(r):
			return scanCloseCurly
		case isOpenSqrt(r):
			return scanOpenSqrt
		case isCloseSqrt(r):
			return scanCloseSqrt
		case isOpenBra(r):
			return scanOpenBra
		case isCloseBra(r):
			return scanCloseBra
		case isQMark(r):
			return scanQMark
		case isDash(r):
			return scanDash
		case isPipe(r):
			return scanPipe
		case isHash(r):
			return scanHash
		case isDot(r):
			return scanDot
		case isStar(r):
			return scanStar
		case isCaret(r):
			return scanCaret
		case isPlusOrDash(r):
			return scanPlusOrMinus
		case isComma(r):
			return scanComma
		case isAt(r):
			return scanAt
		case isGrater(r):
			return scanGrater
		case isLess(r):
			return scanLess
			r = l.next()
		case isNot(r):
			return scanNot
			r = l.next()
		case isAnd(r):
			return scanAnd
			r = l.next()
		case isPercent(r):
			return scanPercent
			r = l.next()
		case isBitNot(r):
			return scanBitNot
			r = l.next()
		}
		r = l.next()
	}
	if isEOF(r) {
		l.emit("ItemEOF")
	} else {
		l.errorf("Not EOF Found")
	}
	return nil
}
