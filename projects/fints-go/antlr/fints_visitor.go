// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fints

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by fintsParser.
type fintsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by fintsParser#nachrichtenkopf.
	VisitNachrichtenkopf(ctx *NachrichtenkopfContext) interface{}

	// Visit a parse tree produced by fintsParser#segmentkopf.
	VisitSegmentkopf(ctx *SegmentkopfContext) interface{}

	// Visit a parse tree produced by fintsParser#segmentkennung.
	VisitSegmentkennung(ctx *SegmentkennungContext) interface{}

	// Visit a parse tree produced by fintsParser#segmentnummer.
	VisitSegmentnummer(ctx *SegmentnummerContext) interface{}

	// Visit a parse tree produced by fintsParser#segmentversion.
	VisitSegmentversion(ctx *SegmentversionContext) interface{}

	// Visit a parse tree produced by fintsParser#bezugssegment.
	VisitBezugssegment(ctx *BezugssegmentContext) interface{}

	// Visit a parse tree produced by fintsParser#nachrichtengroesse.
	VisitNachrichtengroesse(ctx *NachrichtengroesseContext) interface{}

	// Visit a parse tree produced by fintsParser#hbciVersion.
	VisitHbciVersion(ctx *HbciVersionContext) interface{}

	// Visit a parse tree produced by fintsParser#dialogId.
	VisitDialogId(ctx *DialogIdContext) interface{}

	// Visit a parse tree produced by fintsParser#nachrichtennummer.
	VisitNachrichtennummer(ctx *NachrichtennummerContext) interface{}

	// Visit a parse tree produced by fintsParser#bezugsnachricht.
	VisitBezugsnachricht(ctx *BezugsnachrichtContext) interface{}

	// Visit a parse tree produced by fintsParser#de_sep.
	VisitDe_sep(ctx *De_sepContext) interface{}

	// Visit a parse tree produced by fintsParser#deg_sep.
	VisitDeg_sep(ctx *Deg_sepContext) interface{}
}
