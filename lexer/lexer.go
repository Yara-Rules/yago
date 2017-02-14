package lexer

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stateFn func(*Lexer) stateFn

// lexer holds the state of the scanner.
type Lexer struct {
	Name    string    // the name of the input; used only for error reports
	Input   string    // the string being scanned
	State   stateFn   // the next lexing function to enter
	Pos     Pos       // current position in the input
	Start   Pos       // start position of this Item
	Width   Pos       // width of last rune read from input
	LastPos Pos       // position of most recent Item returned by nextItem
	Items   chan Item // channel of scanned Items
	Line    int       // 1+number of newlines seen
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.Pos) >= len(l.Input) {
		l.Width = 0
		return Eof
	}
	r, w := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Width = Pos(w)
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
func (l *Lexer) emit(t ItemType) {
	l.Items <- Item{t, l.Start, l.Input[l.Start:l.Pos], l.Line}
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

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *Lexer) errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := make(map[string]string)
	err["error"] = msg
	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}

// nextItem returns the next Item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextItem() Item {
	Item := <-l.Items
	l.LastPos = Item.Pos
	return Item
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) drain() {
	for range l.Items {
	}
}

// scanned
func (l *Lexer) scanned() string {
	return l.Input[l.Start:l.Pos]
}

// lex creates a new scanner for the input string.
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

const (
	inlineComment = "//"
	leftComment   = "/*"
	rightComment  = "*/"
	blankChars    = " \t\n"
)

// lexText scans until an opening action delimiter, "{{".
func lexText(l *Lexer) stateFn {
	// start collecting Token text
	r := l.next()
	for !isEOF(r) {
		if isEOF(r) {
			l.emit(ItemEOF)
			return nil
		} else if isBlank(r) {
			l.acceptRun(blankChars)
			l.ignore()
		} else if isSlash(r) {
			return scanCommentOrRegex
		} else if isDolar(r) {
			return scanVariable
		} else if isQuote(r) {
			l.ignore() // Removes the first "
			return scanQuote
		} else if isEqual(r) {
			return scanEqual
		} else if isAlphaNumeric(r) {
			return scanKeyword
		} else if isColon(r) {
			return scanColon
		} else if isLeftCurly(r) {
			return scanLeftCurly
		} else if isRigthCurly(r) {
			return scanRightCurly
		} else if isLeftSqrt(r) {
			return scanLeftSqrt
		} else if isRigthSqrt(r) {
			return scanRightSqrt
		} else if isLeftBra(r) {
			return scanLeftBra
		} else if isRigthBra(r) {
			return scanRightBra
		} else if isWildCard(r) {
			return scanWildCard
		} else if isDash(r) {
			return scanDash
		} else if isPipe(r) {
			return scanPipe
		} else if isHash(r) {
			return scanHash
		} else if isDot(r) {
			return scanDot
		} else if isStar(r) {
			return scanStar
		} else if isCaret(r) {
			return scanCaret
		} else if isPlusOrMinus(r) {
			return scanPlusOrMinus
		} else if isComma(r) {
			return scanComma
		} else if isAt(r) {
			return scanAt
		} else if isGratherThan(r) {
			return scanGratherThan
		} else if isLessThan(r) {
			return scanLessThan
		}

		r = l.next()
	}
	l.emit(ItemEOF)
	return nil
}

func scanGratherThan(l *Lexer) stateFn {
	l.emit(ItemGratherThan)
	return lexText
}

func scanLessThan(l *Lexer) stateFn {
	l.emit(ItemLessThan)
	return lexText
}

func scanAt(l *Lexer) stateFn {
	l.emit(ItemAtSym)
	return lexText
}

func scanComma(l *Lexer) stateFn {
	l.emit(ItemComma)
	return lexText
}

func scanPlusOrMinus(l *Lexer) stateFn {
	l.emit(ItemPlusOrMinus)
	return lexText
}

func scanCaret(l *Lexer) stateFn {
	l.emit(ItemCaret)
	return lexText
}

func scanStar(l *Lexer) stateFn {
	l.emit(ItemStar)
	return lexText
}

func scanDot(l *Lexer) stateFn {
	l.emit(ItemDot)
	return lexText
}

func scanHash(l *Lexer) stateFn {
	l.emit(ItemHash)
	return lexText
}

func scanPipe(l *Lexer) stateFn {
	l.emit(ItemPipe)
	return lexText
}

func scanDash(l *Lexer) stateFn {
	l.emit(ItemDash)
	return lexText
}

func scanWildCard(l *Lexer) stateFn {
	l.emit(ItemWildCard)
	return lexText
}

func scanLeftBra(l *Lexer) stateFn {
	l.emit(ItemLeftBra)
	return lexText
}

func scanRightBra(l *Lexer) stateFn {
	l.emit(ItemRightBra)
	return lexText
}

func scanLeftSqrt(l *Lexer) stateFn {
	l.emit(ItemLeftSqrt)
	return lexText
}

func scanRightSqrt(l *Lexer) stateFn {
	l.emit(ItemRightSqrt)
	return lexText
}

func scanLeftCurly(l *Lexer) stateFn {
	l.emit(ItemLeftCurly)
	return lexText
}

func scanRightCurly(l *Lexer) stateFn {
	l.emit(ItemRightCurly)
	return lexText
}

func scanQuote(l *Lexer) stateFn {
	r := l.next() // We've aready got the "
	for !isQuote(r) {
		r = l.next()
		if r == '\\' {
			if l.peek() == '"' {
				r = l.next() // Read the "
				r = l.next() // Prepare r for the next iteration
			} else if l.peek() == '\\' {
				r = l.next() // Read the \
				r = l.next() // Prepare r for the next iteration
			} else if l.peek() == 't' {
				r = l.next() // Read the t
				r = l.next() // Prepare r for the next iteration
			} else if l.peek() == 'n' {
				r = l.next() // Read the n
				r = l.next() // Prepare r for the next iteration
			} else if l.peek() == 'x' {
				r = l.next() // Read the x
				// It needs a two hex characters
				if isHexChar(l.peek()) {
					r = l.next() // Read first hex char
					if isHexChar(l.peek()) {
						r = l.next() // Read first hex char
						r = l.next() // Prepare r for the next iteration
					} else {
						l.errorf("Line %d: illegal escape sequence", l.Line)
					}
				} else {
					l.errorf("Line %d: illegal escape sequence", l.Line)
				}
			} else {
				l.errorf("Line %d: illegal escape sequence", l.Line)
			}
		}
	}
	l.backup() // Remove the last "
	l.emit(ItemString)
	l.next()   // Consume the "
	l.ignore() // Just ignore it
	return lexText
}

func scanEqual(l *Lexer) stateFn {
	l.emit(ItemEqual)
	return lexText
}

func scanVariable(l *Lexer) stateFn {
	var r rune

	if l.peek() == ' ' && l.scanned() == "$" {
		l.emit(ItemVariable)
		return lexText
	}

	for r = l.peek(); !isBlank(r) && isAlphaNumeric(r); r = l.next() {
	}

	if isBlank(r) || !isAlphaNumeric(r) {
		l.backup()
	}

	l.emit(ItemVariable)
	return lexText
}

func scanColon(l *Lexer) stateFn {
	l.emit(ItemColon)
	return lexText
}

func scanKeyword(l *Lexer) stateFn {
	r := l.next()
	length := 1

	for !isBlank(r) && isAlphaNumeric(r) && length <= maxKeywordLength { // we still have content
		length++
		if keyword[l.scanned()] > ItemKeyword {
			l.emit(keyword[l.scanned()])
			return lexText
		}
		r = l.next()
	}

	if isBlank(r) || !isAlphaNumeric(r) {
		l.backup()
	}

	for !isBlank(l.peek()) && isAlphaNumeric(r) {
		l.next()
	}

	if _, err := strconv.Atoi(l.scanned()); err == nil {
		l.emit(ItemNumber)
		return lexText
	}
	l.emit(ItemIdentifier)
	return lexText
}

func scanCommentOrRegex(l *Lexer) stateFn {
	if l.peek() == '*' { // block comment
		end := false
		var r rune
		for !end {
			r = l.next()
			if r == '*' && l.peek() == '/' {
				r = l.next()
				l.emit(ItemComment)
				return lexText
			}
			if r == Eof {
				l.backup()
				l.errorf("Line %d: Expecting end of comment and found end of file.", l.Line)
			}
		}
	} else if l.peek() == '/' { // inline
		r := l.next()
		for !isEndOfLine(r) {
			r = l.next()
		}
		l.emit(ItemComment)
		return lexText
	} else {
		r := l.next()
		for ; !isSlash(r); r = l.next() {
		}
		l.emit(ItemRegex)
		return lexText
	}
	return lexText
}

/*
   Helpers
*/

// isColon reports whether r is a colon
func isColon(r rune) bool {
	return r == ':'
}

// isSlash reports whether r is a slash
func isSlash(r rune) bool {
	return r == '/'
}

func isDolar(r rune) bool {
	return r == '$'
}

func isQuote(r rune) bool {
	return r == '"'
}

func isHexChar(r rune) bool {
	return r == '0' || r == '1' || r == '2' || r == '3' || r == '4' ||
		r == '5' || r == '6' || r == '7' || r == '8' || r == '9' ||
		r == 'a' || r == 'b' || r == 'c' || r == 'd' || r == 'e' ||
		r == 'f' || r == 'A' || r == 'B' || r == 'C' || r == 'D' ||
		r == 'E' || r == 'F'
}

func isEqual(r rune) bool {
	return r == '='
}

func isNewLine(r rune) bool {
	return r == '\n'
}

func isLeftCurly(r rune) bool {
	return r == '{'
}

func isRigthCurly(r rune) bool {
	return r == '}'
}

func isLeftSqrt(r rune) bool {
	return r == '['
}

func isRigthSqrt(r rune) bool {
	return r == ']'
}

func isLeftBra(r rune) bool {
	return r == '('
}

func isRigthBra(r rune) bool {
	return r == ')'
}

func isWildCard(r rune) bool {
	return r == '?'
}

func isDash(r rune) bool {
	return r == '-'
}

func isPipe(r rune) bool {
	return r == '|'
}

func isHash(r rune) bool {
	return r == '#'
}

func isStar(r rune) bool {
	return r == '*'
}

func isCaret(r rune) bool {
	return r == '^'
}

func isPlusOrMinus(r rune) bool {
	return r == '+' || r == '-'
}

func isDot(r rune) bool {
	return r == '.'
}

func isComma(r rune) bool {
	return r == ','
}

func isAt(r rune) bool {
	return r == '@'
}

func isGratherThan(r rune) bool {
	return r == '>'
}

func isLessThan(r rune) bool {
	return r == '<'
}

// isSpace reports whether r is a space character.
func isBlank(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEOF(r rune) bool {
	return r == Eof
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
