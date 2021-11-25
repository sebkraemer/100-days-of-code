// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/pkg/parser/FinTS.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FinTS
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FinTSParser.
type FinTSVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FinTSParser#nachrichtenkopf.
	VisitNachrichtenkopf(ctx *NachrichtenkopfContext) interface{}

	// Visit a parse tree produced by FinTSParser#segmentkopf.
	VisitSegmentkopf(ctx *SegmentkopfContext) interface{}

	// Visit a parse tree produced by FinTSParser#segmentkennung.
	VisitSegmentkennung(ctx *SegmentkennungContext) interface{}

	// Visit a parse tree produced by FinTSParser#segmentnummer.
	VisitSegmentnummer(ctx *SegmentnummerContext) interface{}

	// Visit a parse tree produced by FinTSParser#segmentversion.
	VisitSegmentversion(ctx *SegmentversionContext) interface{}

	// Visit a parse tree produced by FinTSParser#bezugssegment.
	VisitBezugssegment(ctx *BezugssegmentContext) interface{}

	// Visit a parse tree produced by FinTSParser#nachrichtengroesse.
	VisitNachrichtengroesse(ctx *NachrichtengroesseContext) interface{}

	// Visit a parse tree produced by FinTSParser#hbciVersion.
	VisitHbciVersion(ctx *HbciVersionContext) interface{}

	// Visit a parse tree produced by FinTSParser#dialogId.
	VisitDialogId(ctx *DialogIdContext) interface{}

	// Visit a parse tree produced by FinTSParser#nachrichtennummer.
	VisitNachrichtennummer(ctx *NachrichtennummerContext) interface{}

	// Visit a parse tree produced by FinTSParser#bezugsnachricht.
	VisitBezugsnachricht(ctx *BezugsnachrichtContext) interface{}

	// Visit a parse tree produced by FinTSParser#de_sep.
	VisitDe_sep(ctx *De_sepContext) interface{}

	// Visit a parse tree produced by FinTSParser#deg_sep.
	VisitDeg_sep(ctx *Deg_sepContext) interface{}
}
