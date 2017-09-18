package lexic

import (
	"fmt"
	"os"
	"unicode"
)

// checkErr checks if error has happend
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// isKeyword reports wheter w is a yara keyword
func isKeyword(w string) bool {
	for _, kw := range kwList {
		if kw == w {
			return true
		}
	}
	return false
}

// isColon reports whether r is a colon
func isColon(r rune) bool {
	return r == ':'
}

// isSlash reports whether r is a slash
func isSlash(r rune) bool {
	return r == '/'
}

// isBackSlash reports whether r is a back slash
func isBackSlash(r rune) bool {
	return r == '\\'
}

// isDollar reports whether r is a dollar
func isDollar(r rune) bool {
	return r == '$'
}

// isQuote reports whether r is a quote
func isQuote(r rune) bool {
	return r == '"'
}

// isHexChar reports whether r is a hex character
func isHexChar(r rune) bool {
	return r == '0' || r == '1' || r == '2' || r == '3' || r == '4' ||
		r == '5' || r == '6' || r == '7' || r == '8' || r == '9' ||
		r == 'a' || r == 'b' || r == 'c' || r == 'd' || r == 'e' ||
		r == 'f' || r == 'A' || r == 'B' || r == 'C' || r == 'D' ||
		r == 'E' || r == 'F'
}

// isEqual reports whether r is equal
func isEqual(r rune) bool {
	return r == '='
}

// isNot  reports whether r is not
func isNot(r rune) bool {
	return r == '!'
}

// isNewLine reports whether r is new line
func isNewLine(r rune) bool {
	return r == '\n'
}

// isOpenCurly reports whether r is an open curly bracket
func isOpenCurly(r rune) bool {
	return r == '{'
}

// isCloseCurly reports whether r is a close curly bracket
func isCloseCurly(r rune) bool {
	return r == '}'
}

// isOpenSqrt reports whether r is an open square bracket
func isOpenSqrt(r rune) bool {
	return r == '['
}

// isCloseSqrt reports whether r is a close square bracket
func isCloseSqrt(r rune) bool {
	return r == ']'
}

// isOpenBra reports whether r is an open bracket
func isOpenBra(r rune) bool {
	return r == '('
}

// isCloseBra reports whether r is a close bracket
func isCloseBra(r rune) bool {
	return r == ')'
}

// isQMark reports whether r is a question mark
func isQMark(r rune) bool {
	return r == '?'
}

// isDash reports whether r is a dash
func isDash(r rune) bool {
	return r == '-'
}

// isPlus reports whether r is a plus
func isPlus(r rune) bool {
	return r == '+'
}

// isPipe reports whether r is a pipe
func isPipe(r rune) bool {
	return r == '|'
}

// isHash reports whether r is a hash
func isHash(r rune) bool {
	return r == '#'
}

// isStar reports whether r is a asterisk
func isStar(r rune) bool {
	return r == '*'
}

// isCaret reports whether r is a caret
func isCaret(r rune) bool {
	return r == '^'
}

// isPlusOrDash reports whether r is a plus o minus
func isPlusOrDash(r rune) bool {
	return r == '+' || r == '-'
}

// isDot reports whether r is a dot
func isDot(r rune) bool {
	return r == '.'
}

// isComma reports whether r is a comma
func isComma(r rune) bool {
	return r == ','
}

// isAt reports whether r is a @
func isAt(r rune) bool {
	return r == '@'
}

// isGrater reports whether r is a grater than
func isGrater(r rune) bool {
	return r == '>'
}

// isLess reports whether r is a less than
func isLess(r rune) bool {
	return r == '<'
}

// isAnd reports whether r is a logical and
func isAnd(r rune) bool {
	return r == '&'
}

// isBitNot reports whether r is a bitwise not
func isBitNot(r rune) bool {
	return r == '~'
}

// isPercent reports whether r is a percentage
func isPercent(r rune) bool {
	return r == '%'
}

// isBlank reports whether r is a black character
func isBlank(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' '
}

func isValidRegexpMod(r rune) bool {
	return r == 'i' || r == 's'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEOF(r rune) bool {
	return r == EOF
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
