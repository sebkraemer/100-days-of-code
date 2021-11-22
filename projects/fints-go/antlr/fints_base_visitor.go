// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fints

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasefintsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasefintsVisitor) VisitNachrichtenkopf(ctx *NachrichtenkopfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitSegmentkopf(ctx *SegmentkopfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitSegmentkennung(ctx *SegmentkennungContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitSegmentnummer(ctx *SegmentnummerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitSegmentversion(ctx *SegmentversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitBezugssegment(ctx *BezugssegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitNachrichtengroesse(ctx *NachrichtengroesseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitHbciVersion(ctx *HbciVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitDialogId(ctx *DialogIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitNachrichtennummer(ctx *NachrichtennummerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitBezugsnachricht(ctx *BezugsnachrichtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitDe_sep(ctx *De_sepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefintsVisitor) VisitDeg_sep(ctx *Deg_sepContext) interface{} {
	return v.VisitChildren(ctx)
}
