// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 31, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 23, 10, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 2, 11, 2,
	13, 2, 3, 2, 5, 7, 2, 41, 41, 45, 45, 50, 60, 65, 92, 99, 124, 5, 2, 35,
	40, 42, 49, 93, 97, 6, 2, 41, 41, 45, 45, 60, 60, 65, 66, 2, 29, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 3, 15, 3, 2, 2, 2, 5, 17,
	3, 2, 2, 2, 7, 22, 3, 2, 2, 2, 9, 24, 3, 2, 2, 2, 11, 26, 3, 2, 2, 2, 13,
	28, 3, 2, 2, 2, 15, 16, 7, 45, 2, 2, 16, 4, 3, 2, 2, 2, 17, 18, 7, 60,
	2, 2, 18, 6, 3, 2, 2, 2, 19, 23, 5, 13, 7, 2, 20, 23, 5, 9, 5, 2, 21, 23,
	5, 11, 6, 2, 22, 19, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 21, 3, 2, 2, 2,
	23, 8, 3, 2, 2, 2, 24, 25, 9, 2, 2, 2, 25, 10, 3, 2, 2, 2, 26, 27, 9, 3,
	2, 2, 27, 12, 3, 2, 2, 2, 28, 29, 7, 65, 2, 2, 29, 30, 9, 4, 2, 2, 30,
	14, 3, 2, 2, 2, 4, 2, 22, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "DT_AN",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "DT_AN", "FRAG_ALPHANUMERISCH", "FRAG_SYNTAXZEICHEN", "FRAG_ENTWERTET_SYNTAXZEICHEN",
}

type fintsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewfintsLexer(input antlr.CharStream) *fintsLexer {

	l := new(fintsLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "fints.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fintsLexer tokens.
const (
	fintsLexerT__0  = 1
	fintsLexerT__1  = 2
	fintsLexerDT_AN = 3
)
