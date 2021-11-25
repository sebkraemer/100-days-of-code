// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/pkg/parser/FinTS.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FinTS
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFinTSVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFinTSVisitor) VisitNachrichtenkopf(ctx *NachrichtenkopfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitSegmentkopf(ctx *SegmentkopfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitSegmentkennung(ctx *SegmentkennungContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitSegmentnummer(ctx *SegmentnummerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitSegmentversion(ctx *SegmentversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitBezugssegment(ctx *BezugssegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitNachrichtengroesse(ctx *NachrichtengroesseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitHbciVersion(ctx *HbciVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitDialogId(ctx *DialogIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitNachrichtennummer(ctx *NachrichtennummerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitBezugsnachricht(ctx *BezugsnachrichtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitDe_sep(ctx *De_sepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFinTSVisitor) VisitDeg_sep(ctx *Deg_sepContext) interface{} {
	return v.VisitChildren(ctx)
}
