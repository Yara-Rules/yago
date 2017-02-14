package lexer

const maxKeywordLength = 11

const (
	ItemError       ItemType = iota
	ItemNumber               // 1     // Simple number
	ItemComment              // 2     // Comment
	ItemIdentifier           // 3     // Identifier. It could be either rule name or meta field or hex value
	ItemString               // 4     // Quoted string
	ItemRegex                // 5     // Regex string
	ItemVariable             // 6     // Variable starting with '$', such as '$' or  '$1' or '$hello'
	ItemColon                // 7     // :
	ItemEqual                // 8     // =
	ItemLeftCurly            // 9     // {
	ItemRightCurly           // 10    // }
	ItemLeftSqrt             // 11    // [
	ItemRightSqrt            // 12    // ]
	ItemLeftBra              // 13    // (
	ItemRightBra             // 14    // )
	ItemPipe                 // 15    // |
	ItemSpace                // 16    // ' '
	ItemWildCard             // 17    // ?
	ItemDash                 // 18    // -
	ItemHash                 // 19    // #
	ItemDot                  // 20    // .
	ItemCaret                // 21    // ^
	ItemStar                 // 22    // *
	ItemPlusOrMinus          // 23    // +, -
	ItemComma                // 24    // ,
	ItemGratherThan          // 25    // >
	ItemLessThan             // 26    // <
	ItemAtSym                // 27    // @
	ItemEOF                  // 28    // EOF

	// Keywords appear after all the rest.
	ItemKeyword    // 13  // Used only to delimit the keywords
	ItemAll        // 14  //
	ItemAnd        // 15  //
	ItemAny        // 16  //
	ItemAscii      // 17  //
	ItemAt         // 18  //
	ItemCondition  // 19  //
	ItemContains   // 20  //
	ItemEntrypoint // 21  //
	ItemFalse      // 22  //
	ItemFilesize   // 23  //
	ItemFullword   // 24  //
	ItemFor        // 25  //
	ItemGlobal     // 26  //
	ItemIn         // 27  //
	ItemImport     // 28  //
	ItemInclude    // 29  //
	ItemInt8       // 30  //
	ItemInt16      // 31  //
	ItemInt32      // 32  //
	ItemInt8be     // 33  //
	ItemInt16be    // 34  //
	ItemInt32be    // 35  //
	ItemMatches    // 36  //
	ItemMeta       // 37  //
	ItemNocase     // 38  //
	ItemNot        // 39  //
	ItemOr         // 40  //
	ItemOf         // 41  //
	ItemPrivate    // 42  //
	ItemRule       // 43  //
	ItemStrings    // 44  //
	ItemThem       // 45  //
	ItemTrue       // 46  //
	ItemUint8      // 47  //
	ItemUint16     // 48  //
	ItemUint32     // 49  //
	ItemUint8be    // 50  //
	ItemUint16be   // 51  //
	ItemUint32be   // 52  //
	ItemWide       // 53  //
)

var keyword = map[string]ItemType{
	"all":        ItemAll,        //
	"and":        ItemAnd,        //
	"any":        ItemAny,        //
	"ascii":      ItemAscii,      //
	"at":         ItemAt,         //
	"condition":  ItemCondition,  //
	"contains":   ItemContains,   //
	"entrypoint": ItemEntrypoint, //
	"false":      ItemFalse,      //
	"filesize":   ItemFilesize,   //
	"fullword":   ItemFullword,   //
	"for":        ItemFor,        //
	"global":     ItemGlobal,     //
	"in":         ItemIn,         //
	"import":     ItemImport,     //
	"include":    ItemInclude,    //
	"int8":       ItemInt8,       //
	"int16":      ItemInt16,      //
	"int32":      ItemInt32,      //
	"int8be":     ItemInt8be,     //
	"int16be":    ItemInt16be,    //
	"int32be":    ItemInt32be,    //
	"matches":    ItemMatches,    //
	"meta":       ItemMeta,       //
	"nocase":     ItemNocase,     //
	"not":        ItemNot,        //
	"or":         ItemOr,         //
	"of":         ItemOf,         //
	"private":    ItemPrivate,    //
	"rule":       ItemRule,       //
	"strings":    ItemStrings,    //
	"them":       ItemThem,       //
	"true":       ItemTrue,       //
	"uint8":      ItemUint8,      //
	"uint16":     ItemUint16,     //
	"uint32":     ItemUint32,     //
	"uint8be":    ItemUint8be,    //
	"uint16be":   ItemUint16be,   //
	"uint32be":   ItemUint32be,   //
	"wide":       ItemWide,       //
}
