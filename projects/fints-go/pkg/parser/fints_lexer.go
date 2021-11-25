// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/pkg/parser/FinTS.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 4, 11, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 3, 3, 3, 2, 2, 4, 3, 3, 5, 4,
	3, 2, 2, 2, 10, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 3, 7, 3, 2, 2, 2, 5,
	9, 3, 2, 2, 2, 7, 8, 7, 45, 2, 2, 8, 4, 3, 2, 2, 2, 9, 10, 7, 60, 2, 2,
	10, 6, 3, 2, 2, 2, 3, 2, 2,
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

var lexerSymbolicNames []string

var lexerRuleNames = []string{
	"T__0", "T__1",
}

type FinTSLexer struct {
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

func NewFinTSLexer(input antlr.CharStream) *FinTSLexer {

	l := new(FinTSLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FinTS.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FinTSLexer tokens.
const (
	FinTSLexerT__0 = 1
	FinTSLexerT__1 = 2
)
