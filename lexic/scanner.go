package lexic

import (
	"fmt"
	"strconv"
	"strings"
)

// kwList contains all keywords
var kwList = []string{"all", "and", "any", "ascii", "at", "condition", "contains", "entrypoint", "false", "filesize", "fullword", "for", "global", "in", "import", "include", "int8", "int16", "int32", "int8be", "int16be", "int32be", "matches", "meta", "nocase", "not", "or", "of", "private", "rule", "strings", "them", "true", "uint8", "uint16", "uint32", "uint8be", "uint16be", "uint32be", "wide"}

const (
	// Keyword max length
	maxKeywordLength = 11
	// Characters to omit
	blankChars = " \t\n"
)

func scanCommentOrRegex(l *Lexer) stateFn {
	if l.peek() == '*' { // block comment
		end := false
		var r rune
		for !end {
			r = l.next()
			if r == '*' && l.peek() == '/' {
				r = l.next()
				l.emit("ItemComment")
				return lexText
			}
			if r == EOF {
				l.backup()
				l.errorf("Line %d: Expecting end of comment and found end of file.", l.Line)
			}
		}
	} else if l.peek() == '/' { // inline
		r := l.next()
		for !isEndOfLine(r) {
			r = l.next()
		}
		l.emit("ItemComment")
		return lexText
	} else {
		r := l.next()
		for !isSlash(r) {
			if isBackSlash(r) {
				if isBackSlash(l.peek()) {
					r = l.next()
				} else if isSlash(l.peek()) {
					r = l.next() // Read /
				}
			}
			r = l.next()
		}

		mod := ""
		r = l.next()
		fmt.Printf("%s -- '%s'", l.Name, string(r))
		if !isBlank(r) && !isValidRegexpMod(r) {
			l.errorf("Line %d: illegal regex modifier (%s)", l.Line, string(r))
		}
		for isValidRegexpMod(r) {
			mod += string(r)
			r = l.next()
		}
		l.backup()
		l.emit("ItemRegex")
		return lexText
	}
	return lexText
}

func scanVariable(l *Lexer) stateFn {
	var r rune

	if l.scanned() == "$" && isSpace(l.peek()) || isStar(l.peek()) || isEqual(l.peek()) { // $, $*
		l.emit("ItemVariable")
		return lexText
	}

	for r = l.peek(); !isBlank(r) && isAlphaNumeric(r); r = l.next() {
	}

	if isBlank(r) || !isAlphaNumeric(r) {
		l.backup()
	}

	l.emit("ItemVariable")
	return lexText
}

func scanQuote(l *Lexer) stateFn {
	r := l.next() // We've aready got the "
	for !isQuote(r) {
		if isBackSlash(r) {
			if isQuote(l.peek()) {
				r = l.next() // Read the "
			} else if l.peek() == '\\' {
				r = l.next() // Read the \
			} else if l.peek() == 't' {
				r = l.next() // Read the t
			} else if l.peek() == 'n' {
				r = l.next() // Read the n
			} else if l.peek() == 'x' {
				r = l.next() // Read the x
				// It needs a two hex characters
				if isHexChar(l.peek()) {
					r = l.next() // Read first hex char
					if isHexChar(l.peek()) {
						r = l.next() // Read first hex char
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
		r = l.next()
	}
	l.backup() // Remove the last "
	l.emit("ItemString")
	l.next()   // Consume the "
	l.ignore() // Just ignore it
	return lexText
}

func scanEqual(l *Lexer) stateFn {
	if isEqual(l.peek()) {
		_ = l.next()
		l.emit("ItemEqualEqual")
	} else {
		l.emit("ItemEqual")
	}
	return lexText
}

func scanNot(l *Lexer) stateFn {
	if isEqual(l.peek()) {
		_ = l.next()
		l.emit("ItemNotEqual")
	} else {
		l.emit("ItemNot")
	}
	return lexText
}

func scanKeyword(l *Lexer) stateFn {
	r := l.next()
	length := 1

	for !isBlank(r) && isAlphaNumeric(r) && length <= maxKeywordLength { // we still have content
		length++
		if isKeyword(l.scanned()) {
			if !isAlphaNumeric(l.peek()) {
				l.emit(fmt.Sprintf("ItemKW%s", strings.Title(l.scanned())))
				return lexText
			}
		}
		r = l.next()
	}

	if isBlank(r) || !isAlphaNumeric(r) {
		l.backup()
	}

	for !isBlank(l.peek()) && isAlphaNumeric(l.peek()) {
		l.next()
	}

	if _, err := strconv.Atoi(l.scanned()); err == nil {
		l.emit("ItemIntNumber")
		return lexText
	}

	l.emit("ItemIdentifier")
	return lexText
}

func scanColon(l *Lexer) stateFn {
	l.emit("ItemColon")
	return lexText
}

func scanOpenCurly(l *Lexer) stateFn {
	l.emit("ItemOCurly")
	return lexText
}

func scanCloseCurly(l *Lexer) stateFn {
	l.emit("ItemCCurly")
	return lexText
}

func scanOpenSqrt(l *Lexer) stateFn {
	l.emit("ItemOSqrt")
	return lexText
}

func scanCloseSqrt(l *Lexer) stateFn {
	l.emit("ItemCSqrt")
	return lexText
}

func scanOpenBra(l *Lexer) stateFn {
	l.emit("ItemOBracket")
	return lexText
}

func scanCloseBra(l *Lexer) stateFn {
	l.emit("ItemCBracket")
	return lexText
}

func scanQMark(l *Lexer) stateFn {
	l.emit("ItemQMark")
	return lexText
}

func scanDash(l *Lexer) stateFn {
	l.emit("ItemDash")
	return lexText
}

func scanPipe(l *Lexer) stateFn {
	l.emit("ItemPipe")
	return lexText
}

func scanHash(l *Lexer) stateFn {
	l.emit("ItemHash")
	return lexText
}

func scanDot(l *Lexer) stateFn {
	if isDot(l.peek()) {
		_ = l.next()
		l.emit("ItemDotDot")
	} else {
		l.emit("ItemDot")
	}
	return lexText
}

func scanStar(l *Lexer) stateFn {
	l.emit("ItemStar")
	return lexText
}

func scanCaret(l *Lexer) stateFn {
	l.emit("ItemCaret")
	return lexText
}

func scanPlusOrMinus(l *Lexer) stateFn {
	if l.scanned() == "-" {
		l.emit("ItemDash")
	} else {
		l.emit("ItemPlus")
	}
	return lexText
}

func scanComma(l *Lexer) stateFn {
	l.emit("ItemComma")
	return lexText
}

func scanAt(l *Lexer) stateFn {
	l.emit("ItemAt")
	return lexText
}

func scanGrater(l *Lexer) stateFn {
	if isEqual(l.peek()) {
		_ = l.next()
		l.emit("ItemGraterEqual")
	} else if isGrater(l.peek()) {
		_ = l.next()
		l.emit("ItemRightShift")
	} else {
		l.emit("ItemGrater")
	}
	return lexText
}

func scanLess(l *Lexer) stateFn {
	if isEqual(l.peek()) {
		_ = l.next()
		l.emit("ItemLessEqual")
	} else if isLess(l.peek()) {
		_ = l.next()
		l.emit("ItemLeftShift")
	} else {
		l.emit("ItemLess")
	}
	return lexText
}

func scanAnd(l *Lexer) stateFn {
	l.emit("ItemAnd")
	return lexText
}

func scanPercent(l *Lexer) stateFn {
	l.emit("ItemPercent")
	return lexText
}

func scanBitNot(l *Lexer) stateFn {
	l.emit("ItemBitNot")
	return lexText
}
