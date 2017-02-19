package parser

import (
	"github.com/Sirupsen/logrus"
	"github.com/Yara-Rules/yago/lexer"
)

type Parser struct {
	Name      string         `json:"file_name"  bson:"file_name"`
	Lex       *lexer.Lexer   `json:"-"          bson:"-"`
	LastItem  lexer.Item     `json:"-"          bson:"-"`
	peekCount int            `json:"-"          bson:"-"`
	token     [3]lexer.Item  `json:"-"          bson:"-"` // three-token lookahead for parser.
	Imports   []string       `json:"imports"    bson:"imports"`
	Rules     []RuleDef      `json:"rules"      bson:"rules"`
	log       *logrus.Logger `json:"-"          bson:"-"`
}

type StringDef struct {
	Name      string   `json:"name"     bson:"name"`
	Value     string   `json:"value"    bson:"value"`
	Modifiers []string `json:"modifers" bson:"modifiers"`
}

type RuleDef struct {
	Name      string            `json:"name"      bson:"name"`
	Global    bool              `json:"global"    bson:"global"`
	Private   bool              `json:"private"   bson:"private"`
	Tags      []string          `json:"tags"      bson:"tags"`
	Meta      map[string]string `json:"meta"      bson:"meta"`
	Strings   []StringDef       `json:"strings"   bson:"strings"`
	Condition string            `json:"condition" bson:"condition"`
}
