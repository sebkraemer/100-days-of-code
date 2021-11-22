// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fints

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefintsListener is a complete listener for a parse tree produced by fintsParser.
type BasefintsListener struct{}

var _ fintsListener = &BasefintsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefintsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefintsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefintsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefintsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNachrichtenkopf is called when production nachrichtenkopf is entered.
func (s *BasefintsListener) EnterNachrichtenkopf(ctx *NachrichtenkopfContext) {}

// ExitNachrichtenkopf is called when production nachrichtenkopf is exited.
func (s *BasefintsListener) ExitNachrichtenkopf(ctx *NachrichtenkopfContext) {}

// EnterSegmentkopf is called when production segmentkopf is entered.
func (s *BasefintsListener) EnterSegmentkopf(ctx *SegmentkopfContext) {}

// ExitSegmentkopf is called when production segmentkopf is exited.
func (s *BasefintsListener) ExitSegmentkopf(ctx *SegmentkopfContext) {}

// EnterSegmentkennung is called when production segmentkennung is entered.
func (s *BasefintsListener) EnterSegmentkennung(ctx *SegmentkennungContext) {}

// ExitSegmentkennung is called when production segmentkennung is exited.
func (s *BasefintsListener) ExitSegmentkennung(ctx *SegmentkennungContext) {}

// EnterSegmentnummer is called when production segmentnummer is entered.
func (s *BasefintsListener) EnterSegmentnummer(ctx *SegmentnummerContext) {}

// ExitSegmentnummer is called when production segmentnummer is exited.
func (s *BasefintsListener) ExitSegmentnummer(ctx *SegmentnummerContext) {}

// EnterSegmentversion is called when production segmentversion is entered.
func (s *BasefintsListener) EnterSegmentversion(ctx *SegmentversionContext) {}

// ExitSegmentversion is called when production segmentversion is exited.
func (s *BasefintsListener) ExitSegmentversion(ctx *SegmentversionContext) {}

// EnterBezugssegment is called when production bezugssegment is entered.
func (s *BasefintsListener) EnterBezugssegment(ctx *BezugssegmentContext) {}

// ExitBezugssegment is called when production bezugssegment is exited.
func (s *BasefintsListener) ExitBezugssegment(ctx *BezugssegmentContext) {}

// EnterNachrichtengroesse is called when production nachrichtengroesse is entered.
func (s *BasefintsListener) EnterNachrichtengroesse(ctx *NachrichtengroesseContext) {}

// ExitNachrichtengroesse is called when production nachrichtengroesse is exited.
func (s *BasefintsListener) ExitNachrichtengroesse(ctx *NachrichtengroesseContext) {}

// EnterHbciVersion is called when production hbciVersion is entered.
func (s *BasefintsListener) EnterHbciVersion(ctx *HbciVersionContext) {}

// ExitHbciVersion is called when production hbciVersion is exited.
func (s *BasefintsListener) ExitHbciVersion(ctx *HbciVersionContext) {}

// EnterDialogId is called when production dialogId is entered.
func (s *BasefintsListener) EnterDialogId(ctx *DialogIdContext) {}

// ExitDialogId is called when production dialogId is exited.
func (s *BasefintsListener) ExitDialogId(ctx *DialogIdContext) {}

// EnterNachrichtennummer is called when production nachrichtennummer is entered.
func (s *BasefintsListener) EnterNachrichtennummer(ctx *NachrichtennummerContext) {}

// ExitNachrichtennummer is called when production nachrichtennummer is exited.
func (s *BasefintsListener) ExitNachrichtennummer(ctx *NachrichtennummerContext) {}

// EnterBezugsnachricht is called when production bezugsnachricht is entered.
func (s *BasefintsListener) EnterBezugsnachricht(ctx *BezugsnachrichtContext) {}

// ExitBezugsnachricht is called when production bezugsnachricht is exited.
func (s *BasefintsListener) ExitBezugsnachricht(ctx *BezugsnachrichtContext) {}

// EnterDe_sep is called when production de_sep is entered.
func (s *BasefintsListener) EnterDe_sep(ctx *De_sepContext) {}

// ExitDe_sep is called when production de_sep is exited.
func (s *BasefintsListener) ExitDe_sep(ctx *De_sepContext) {}

// EnterDeg_sep is called when production deg_sep is entered.
func (s *BasefintsListener) EnterDeg_sep(ctx *Deg_sepContext) {}

// ExitDeg_sep is called when production deg_sep is exited.
func (s *BasefintsListener) ExitDeg_sep(ctx *Deg_sepContext) {}
