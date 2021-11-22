// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fints

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 101, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 41, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 51, 10, 3, 3, 4, 6, 4, 54, 10, 4, 13, 4, 14, 4, 55, 3,
	5, 6, 5, 59, 10, 5, 13, 5, 14, 5, 60, 3, 6, 6, 6, 64, 10, 6, 13, 6, 14,
	6, 65, 3, 7, 6, 7, 69, 10, 7, 13, 7, 14, 7, 70, 3, 8, 6, 8, 74, 10, 8,
	13, 8, 14, 8, 75, 3, 9, 6, 9, 79, 10, 9, 13, 9, 14, 9, 80, 3, 10, 6, 10,
	84, 10, 10, 13, 10, 14, 10, 85, 3, 11, 6, 11, 89, 10, 11, 13, 11, 14, 11,
	90, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 10,
	55, 60, 65, 70, 75, 80, 85, 90, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 2, 2, 2, 97, 2, 28, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 53,
	3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 68, 3, 2, 2, 2,
	14, 73, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20, 88, 3,
	2, 2, 2, 22, 92, 3, 2, 2, 2, 24, 96, 3, 2, 2, 2, 26, 98, 3, 2, 2, 2, 28,
	29, 5, 4, 3, 2, 29, 30, 5, 24, 13, 2, 30, 31, 5, 14, 8, 2, 31, 32, 5, 24,
	13, 2, 32, 33, 5, 16, 9, 2, 33, 34, 5, 24, 13, 2, 34, 35, 5, 18, 10, 2,
	35, 36, 5, 24, 13, 2, 36, 40, 5, 20, 11, 2, 37, 38, 5, 24, 13, 2, 38, 39,
	5, 22, 12, 2, 39, 41, 3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 40, 41, 3, 2, 2,
	2, 41, 3, 3, 2, 2, 2, 42, 43, 5, 6, 4, 2, 43, 44, 5, 26, 14, 2, 44, 45,
	5, 8, 5, 2, 45, 46, 5, 26, 14, 2, 46, 50, 5, 10, 6, 2, 47, 48, 5, 26, 14,
	2, 48, 49, 5, 12, 7, 2, 49, 51, 3, 2, 2, 2, 50, 47, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 5, 3, 2, 2, 2, 52, 54, 7, 5, 2, 2, 53, 52, 3, 2, 2, 2,
	54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 7, 3, 2,
	2, 2, 57, 59, 7, 5, 2, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 64, 7, 5, 2, 2,
	63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 65, 63, 3,
	2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 69, 7, 5, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	70, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 13, 3, 2, 2,
	2, 72, 74, 7, 5, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 15, 3, 2, 2, 2, 77, 79, 7, 5, 2, 2,
	78, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 80, 78, 3,
	2, 2, 2, 81, 17, 3, 2, 2, 2, 82, 84, 7, 5, 2, 2, 83, 82, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 19, 3, 2, 2,
	2, 87, 89, 7, 5, 2, 2, 88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 21, 3, 2, 2, 2, 92, 93, 5, 18, 10,
	2, 93, 94, 5, 26, 14, 2, 94, 95, 5, 20, 11, 2, 95, 23, 3, 2, 2, 2, 96,
	97, 7, 3, 2, 2, 97, 25, 3, 2, 2, 2, 98, 99, 7, 4, 2, 2, 99, 27, 3, 2, 2,
	2, 12, 40, 50, 55, 60, 65, 70, 75, 80, 85, 90,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "':'",
}
var symbolicNames = []string{
	"", "", "", "DT_AN",
}

var ruleNames = []string{
	"nachrichtenkopf", "segmentkopf", "segmentkennung", "segmentnummer", "segmentversion",
	"bezugssegment", "nachrichtengroesse", "hbciVersion", "dialogId", "nachrichtennummer",
	"bezugsnachricht", "de_sep", "deg_sep",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type fintsParser struct {
	*antlr.BaseParser
}

func NewfintsParser(input antlr.TokenStream) *fintsParser {
	this := new(fintsParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "fints.g4"

	return this
}

// fintsParser tokens.
const (
	fintsParserEOF   = antlr.TokenEOF
	fintsParserT__0  = 1
	fintsParserT__1  = 2
	fintsParserDT_AN = 3
)

// fintsParser rules.
const (
	fintsParserRULE_nachrichtenkopf    = 0
	fintsParserRULE_segmentkopf        = 1
	fintsParserRULE_segmentkennung     = 2
	fintsParserRULE_segmentnummer      = 3
	fintsParserRULE_segmentversion     = 4
	fintsParserRULE_bezugssegment      = 5
	fintsParserRULE_nachrichtengroesse = 6
	fintsParserRULE_hbciVersion        = 7
	fintsParserRULE_dialogId           = 8
	fintsParserRULE_nachrichtennummer  = 9
	fintsParserRULE_bezugsnachricht    = 10
	fintsParserRULE_de_sep             = 11
	fintsParserRULE_deg_sep            = 12
)

// INachrichtenkopfContext is an interface to support dynamic dispatch.
type INachrichtenkopfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNachrichtenkopfContext differentiates from other interfaces.
	IsNachrichtenkopfContext()
}

type NachrichtenkopfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNachrichtenkopfContext() *NachrichtenkopfContext {
	var p = new(NachrichtenkopfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_nachrichtenkopf
	return p
}

func (*NachrichtenkopfContext) IsNachrichtenkopfContext() {}

func NewNachrichtenkopfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NachrichtenkopfContext {
	var p = new(NachrichtenkopfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_nachrichtenkopf

	return p
}

func (s *NachrichtenkopfContext) GetParser() antlr.Parser { return s.parser }

func (s *NachrichtenkopfContext) Segmentkopf() ISegmentkopfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentkopfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentkopfContext)
}

func (s *NachrichtenkopfContext) AllDe_sep() []IDe_sepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDe_sepContext)(nil)).Elem())
	var tst = make([]IDe_sepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDe_sepContext)
		}
	}

	return tst
}

func (s *NachrichtenkopfContext) De_sep(i int) IDe_sepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDe_sepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDe_sepContext)
}

func (s *NachrichtenkopfContext) Nachrichtengroesse() INachrichtengroesseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INachrichtengroesseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INachrichtengroesseContext)
}

func (s *NachrichtenkopfContext) HbciVersion() IHbciVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHbciVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHbciVersionContext)
}

func (s *NachrichtenkopfContext) DialogId() IDialogIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDialogIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDialogIdContext)
}

func (s *NachrichtenkopfContext) Nachrichtennummer() INachrichtennummerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INachrichtennummerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INachrichtennummerContext)
}

func (s *NachrichtenkopfContext) Bezugsnachricht() IBezugsnachrichtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBezugsnachrichtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBezugsnachrichtContext)
}

func (s *NachrichtenkopfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NachrichtenkopfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NachrichtenkopfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterNachrichtenkopf(s)
	}
}

func (s *NachrichtenkopfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitNachrichtenkopf(s)
	}
}

func (s *NachrichtenkopfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitNachrichtenkopf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Nachrichtenkopf() (localctx INachrichtenkopfContext) {
	localctx = NewNachrichtenkopfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, fintsParserRULE_nachrichtenkopf)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Segmentkopf()
	}
	{
		p.SetState(27)
		p.De_sep()
	}
	{
		p.SetState(28)
		p.Nachrichtengroesse()
	}
	{
		p.SetState(29)
		p.De_sep()
	}
	{
		p.SetState(30)
		p.HbciVersion()
	}
	{
		p.SetState(31)
		p.De_sep()
	}
	{
		p.SetState(32)
		p.DialogId()
	}
	{
		p.SetState(33)
		p.De_sep()
	}
	{
		p.SetState(34)
		p.Nachrichtennummer()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == fintsParserT__0 {
		{
			p.SetState(35)
			p.De_sep()
		}
		{
			p.SetState(36)
			p.Bezugsnachricht()
		}

	}

	return localctx
}

// ISegmentkopfContext is an interface to support dynamic dispatch.
type ISegmentkopfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentkopfContext differentiates from other interfaces.
	IsSegmentkopfContext()
}

type SegmentkopfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentkopfContext() *SegmentkopfContext {
	var p = new(SegmentkopfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_segmentkopf
	return p
}

func (*SegmentkopfContext) IsSegmentkopfContext() {}

func NewSegmentkopfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentkopfContext {
	var p = new(SegmentkopfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_segmentkopf

	return p
}

func (s *SegmentkopfContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentkopfContext) Segmentkennung() ISegmentkennungContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentkennungContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentkennungContext)
}

func (s *SegmentkopfContext) AllDeg_sep() []IDeg_sepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeg_sepContext)(nil)).Elem())
	var tst = make([]IDeg_sepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeg_sepContext)
		}
	}

	return tst
}

func (s *SegmentkopfContext) Deg_sep(i int) IDeg_sepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeg_sepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeg_sepContext)
}

func (s *SegmentkopfContext) Segmentnummer() ISegmentnummerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentnummerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentnummerContext)
}

func (s *SegmentkopfContext) Segmentversion() ISegmentversionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentversionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentversionContext)
}

func (s *SegmentkopfContext) Bezugssegment() IBezugssegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBezugssegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBezugssegmentContext)
}

func (s *SegmentkopfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentkopfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentkopfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterSegmentkopf(s)
	}
}

func (s *SegmentkopfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitSegmentkopf(s)
	}
}

func (s *SegmentkopfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitSegmentkopf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Segmentkopf() (localctx ISegmentkopfContext) {
	localctx = NewSegmentkopfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, fintsParserRULE_segmentkopf)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Segmentkennung()
	}
	{
		p.SetState(41)
		p.Deg_sep()
	}
	{
		p.SetState(42)
		p.Segmentnummer()
	}
	{
		p.SetState(43)
		p.Deg_sep()
	}
	{
		p.SetState(44)
		p.Segmentversion()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == fintsParserT__1 {
		{
			p.SetState(45)
			p.Deg_sep()
		}
		{
			p.SetState(46)
			p.Bezugssegment()
		}

	}

	return localctx
}

// ISegmentkennungContext is an interface to support dynamic dispatch.
type ISegmentkennungContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentkennungContext differentiates from other interfaces.
	IsSegmentkennungContext()
}

type SegmentkennungContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentkennungContext() *SegmentkennungContext {
	var p = new(SegmentkennungContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_segmentkennung
	return p
}

func (*SegmentkennungContext) IsSegmentkennungContext() {}

func NewSegmentkennungContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentkennungContext {
	var p = new(SegmentkennungContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_segmentkennung

	return p
}

func (s *SegmentkennungContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentkennungContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *SegmentkennungContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *SegmentkennungContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentkennungContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentkennungContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterSegmentkennung(s)
	}
}

func (s *SegmentkennungContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitSegmentkennung(s)
	}
}

func (s *SegmentkennungContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitSegmentkennung(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Segmentkennung() (localctx ISegmentkennungContext) {
	localctx = NewSegmentkennungContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fintsParserRULE_segmentkennung)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(50)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ISegmentnummerContext is an interface to support dynamic dispatch.
type ISegmentnummerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentnummerContext differentiates from other interfaces.
	IsSegmentnummerContext()
}

type SegmentnummerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentnummerContext() *SegmentnummerContext {
	var p = new(SegmentnummerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_segmentnummer
	return p
}

func (*SegmentnummerContext) IsSegmentnummerContext() {}

func NewSegmentnummerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentnummerContext {
	var p = new(SegmentnummerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_segmentnummer

	return p
}

func (s *SegmentnummerContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentnummerContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *SegmentnummerContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *SegmentnummerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentnummerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentnummerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterSegmentnummer(s)
	}
}

func (s *SegmentnummerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitSegmentnummer(s)
	}
}

func (s *SegmentnummerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitSegmentnummer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Segmentnummer() (localctx ISegmentnummerContext) {
	localctx = NewSegmentnummerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fintsParserRULE_segmentnummer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(55)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// ISegmentversionContext is an interface to support dynamic dispatch.
type ISegmentversionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentversionContext differentiates from other interfaces.
	IsSegmentversionContext()
}

type SegmentversionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentversionContext() *SegmentversionContext {
	var p = new(SegmentversionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_segmentversion
	return p
}

func (*SegmentversionContext) IsSegmentversionContext() {}

func NewSegmentversionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentversionContext {
	var p = new(SegmentversionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_segmentversion

	return p
}

func (s *SegmentversionContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentversionContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *SegmentversionContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *SegmentversionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentversionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentversionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterSegmentversion(s)
	}
}

func (s *SegmentversionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitSegmentversion(s)
	}
}

func (s *SegmentversionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitSegmentversion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Segmentversion() (localctx ISegmentversionContext) {
	localctx = NewSegmentversionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fintsParserRULE_segmentversion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(60)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IBezugssegmentContext is an interface to support dynamic dispatch.
type IBezugssegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBezugssegmentContext differentiates from other interfaces.
	IsBezugssegmentContext()
}

type BezugssegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBezugssegmentContext() *BezugssegmentContext {
	var p = new(BezugssegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_bezugssegment
	return p
}

func (*BezugssegmentContext) IsBezugssegmentContext() {}

func NewBezugssegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BezugssegmentContext {
	var p = new(BezugssegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_bezugssegment

	return p
}

func (s *BezugssegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *BezugssegmentContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *BezugssegmentContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *BezugssegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BezugssegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BezugssegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterBezugssegment(s)
	}
}

func (s *BezugssegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitBezugssegment(s)
	}
}

func (s *BezugssegmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitBezugssegment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Bezugssegment() (localctx IBezugssegmentContext) {
	localctx = NewBezugssegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, fintsParserRULE_bezugssegment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(65)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// INachrichtengroesseContext is an interface to support dynamic dispatch.
type INachrichtengroesseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNachrichtengroesseContext differentiates from other interfaces.
	IsNachrichtengroesseContext()
}

type NachrichtengroesseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNachrichtengroesseContext() *NachrichtengroesseContext {
	var p = new(NachrichtengroesseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_nachrichtengroesse
	return p
}

func (*NachrichtengroesseContext) IsNachrichtengroesseContext() {}

func NewNachrichtengroesseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NachrichtengroesseContext {
	var p = new(NachrichtengroesseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_nachrichtengroesse

	return p
}

func (s *NachrichtengroesseContext) GetParser() antlr.Parser { return s.parser }

func (s *NachrichtengroesseContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *NachrichtengroesseContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *NachrichtengroesseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NachrichtengroesseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NachrichtengroesseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterNachrichtengroesse(s)
	}
}

func (s *NachrichtengroesseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitNachrichtengroesse(s)
	}
}

func (s *NachrichtengroesseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitNachrichtengroesse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Nachrichtengroesse() (localctx INachrichtengroesseContext) {
	localctx = NewNachrichtengroesseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, fintsParserRULE_nachrichtengroesse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(70)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IHbciVersionContext is an interface to support dynamic dispatch.
type IHbciVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHbciVersionContext differentiates from other interfaces.
	IsHbciVersionContext()
}

type HbciVersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHbciVersionContext() *HbciVersionContext {
	var p = new(HbciVersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_hbciVersion
	return p
}

func (*HbciVersionContext) IsHbciVersionContext() {}

func NewHbciVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HbciVersionContext {
	var p = new(HbciVersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_hbciVersion

	return p
}

func (s *HbciVersionContext) GetParser() antlr.Parser { return s.parser }

func (s *HbciVersionContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *HbciVersionContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *HbciVersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HbciVersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HbciVersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterHbciVersion(s)
	}
}

func (s *HbciVersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitHbciVersion(s)
	}
}

func (s *HbciVersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitHbciVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) HbciVersion() (localctx IHbciVersionContext) {
	localctx = NewHbciVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fintsParserRULE_hbciVersion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(75)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IDialogIdContext is an interface to support dynamic dispatch.
type IDialogIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDialogIdContext differentiates from other interfaces.
	IsDialogIdContext()
}

type DialogIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDialogIdContext() *DialogIdContext {
	var p = new(DialogIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_dialogId
	return p
}

func (*DialogIdContext) IsDialogIdContext() {}

func NewDialogIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DialogIdContext {
	var p = new(DialogIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_dialogId

	return p
}

func (s *DialogIdContext) GetParser() antlr.Parser { return s.parser }

func (s *DialogIdContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *DialogIdContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *DialogIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DialogIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DialogIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterDialogId(s)
	}
}

func (s *DialogIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitDialogId(s)
	}
}

func (s *DialogIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitDialogId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) DialogId() (localctx IDialogIdContext) {
	localctx = NewDialogIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fintsParserRULE_dialogId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(80)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// INachrichtennummerContext is an interface to support dynamic dispatch.
type INachrichtennummerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNachrichtennummerContext differentiates from other interfaces.
	IsNachrichtennummerContext()
}

type NachrichtennummerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNachrichtennummerContext() *NachrichtennummerContext {
	var p = new(NachrichtennummerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_nachrichtennummer
	return p
}

func (*NachrichtennummerContext) IsNachrichtennummerContext() {}

func NewNachrichtennummerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NachrichtennummerContext {
	var p = new(NachrichtennummerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_nachrichtennummer

	return p
}

func (s *NachrichtennummerContext) GetParser() antlr.Parser { return s.parser }

func (s *NachrichtennummerContext) AllDT_AN() []antlr.TerminalNode {
	return s.GetTokens(fintsParserDT_AN)
}

func (s *NachrichtennummerContext) DT_AN(i int) antlr.TerminalNode {
	return s.GetToken(fintsParserDT_AN, i)
}

func (s *NachrichtennummerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NachrichtennummerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NachrichtennummerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterNachrichtennummer(s)
	}
}

func (s *NachrichtennummerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitNachrichtennummer(s)
	}
}

func (s *NachrichtennummerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitNachrichtennummer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Nachrichtennummer() (localctx INachrichtennummerContext) {
	localctx = NewNachrichtennummerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fintsParserRULE_nachrichtennummer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(85)
				p.Match(fintsParserDT_AN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IBezugsnachrichtContext is an interface to support dynamic dispatch.
type IBezugsnachrichtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBezugsnachrichtContext differentiates from other interfaces.
	IsBezugsnachrichtContext()
}

type BezugsnachrichtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBezugsnachrichtContext() *BezugsnachrichtContext {
	var p = new(BezugsnachrichtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_bezugsnachricht
	return p
}

func (*BezugsnachrichtContext) IsBezugsnachrichtContext() {}

func NewBezugsnachrichtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BezugsnachrichtContext {
	var p = new(BezugsnachrichtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_bezugsnachricht

	return p
}

func (s *BezugsnachrichtContext) GetParser() antlr.Parser { return s.parser }

func (s *BezugsnachrichtContext) DialogId() IDialogIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDialogIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDialogIdContext)
}

func (s *BezugsnachrichtContext) Deg_sep() IDeg_sepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeg_sepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeg_sepContext)
}

func (s *BezugsnachrichtContext) Nachrichtennummer() INachrichtennummerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INachrichtennummerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INachrichtennummerContext)
}

func (s *BezugsnachrichtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BezugsnachrichtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BezugsnachrichtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterBezugsnachricht(s)
	}
}

func (s *BezugsnachrichtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitBezugsnachricht(s)
	}
}

func (s *BezugsnachrichtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitBezugsnachricht(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Bezugsnachricht() (localctx IBezugsnachrichtContext) {
	localctx = NewBezugsnachrichtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, fintsParserRULE_bezugsnachricht)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.DialogId()
	}
	{
		p.SetState(91)
		p.Deg_sep()
	}
	{
		p.SetState(92)
		p.Nachrichtennummer()
	}

	return localctx
}

// IDe_sepContext is an interface to support dynamic dispatch.
type IDe_sepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDe_sepContext differentiates from other interfaces.
	IsDe_sepContext()
}

type De_sepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDe_sepContext() *De_sepContext {
	var p = new(De_sepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_de_sep
	return p
}

func (*De_sepContext) IsDe_sepContext() {}

func NewDe_sepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *De_sepContext {
	var p = new(De_sepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_de_sep

	return p
}

func (s *De_sepContext) GetParser() antlr.Parser { return s.parser }
func (s *De_sepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *De_sepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *De_sepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterDe_sep(s)
	}
}

func (s *De_sepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitDe_sep(s)
	}
}

func (s *De_sepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitDe_sep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) De_sep() (localctx IDe_sepContext) {
	localctx = NewDe_sepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, fintsParserRULE_de_sep)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(fintsParserT__0)
	}

	return localctx
}

// IDeg_sepContext is an interface to support dynamic dispatch.
type IDeg_sepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeg_sepContext differentiates from other interfaces.
	IsDeg_sepContext()
}

type Deg_sepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeg_sepContext() *Deg_sepContext {
	var p = new(Deg_sepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fintsParserRULE_deg_sep
	return p
}

func (*Deg_sepContext) IsDeg_sepContext() {}

func NewDeg_sepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Deg_sepContext {
	var p = new(Deg_sepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fintsParserRULE_deg_sep

	return p
}

func (s *Deg_sepContext) GetParser() antlr.Parser { return s.parser }
func (s *Deg_sepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Deg_sepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Deg_sepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.EnterDeg_sep(s)
	}
}

func (s *Deg_sepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fintsListener); ok {
		listenerT.ExitDeg_sep(s)
	}
}

func (s *Deg_sepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case fintsVisitor:
		return t.VisitDeg_sep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *fintsParser) Deg_sep() (localctx IDeg_sepContext) {
	localctx = NewDeg_sepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, fintsParserRULE_deg_sep)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(fintsParserT__1)
	}

	return localctx
}
