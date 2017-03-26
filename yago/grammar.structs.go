package yago

import (
	"github.com/Sirupsen/logrus"
	"github.com/Yara-Rules/yago/lexic"
)

// Types of string variables
const (
	StringType = iota
	StringString
	StringRegex
	StringHex
)

// Parser represents the Yara rules
type Parser struct {
	Name      string         `json:"file_name"`
	Lex       *lexic.Lexer   `json:"-"`
	Item      lexic.Item     `json:"-"`
	LastItem  lexic.Item     `json:"-"`
	peekCount int            `json:"-"`
	token     [2]lexic.Item  `json:"-"` // two-token lookahead for parser.
	Imports   []string       `json:"imports"`
	Rules     []RuleDef      `json:"rules"`
	log       *logrus.Logger `json:"-"`
}

// StringDef defines a string variable
type StringDef struct {
	Name      string   `json:"name"`
	Value     string   `json:"value"`
	Modifiers []string `json:"modifers"`
	Typ       int      `json:"type"`
}

// RuleDef defines a yara rule
type RuleDef struct {
	Name      string            `json:"name"`
	Global    bool              `json:"global"`
	Private   bool              `json:"private"`
	Tags      []string          `json:"tags"`
	Meta      map[string]string `json:"meta"`
	Strings   []StringDef       `json:"strings"`
	Condition string            `json:"condition"`
}
