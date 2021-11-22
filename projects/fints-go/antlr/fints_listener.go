// Code generated from /Users/sebi/src/100-days-of-code/projects/fints-go/antlr/fints.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fints

import "github.com/antlr/antlr4/runtime/Go/antlr"

// fintsListener is a complete listener for a parse tree produced by fintsParser.
type fintsListener interface {
	antlr.ParseTreeListener

	// EnterNachrichtenkopf is called when entering the nachrichtenkopf production.
	EnterNachrichtenkopf(c *NachrichtenkopfContext)

	// EnterSegmentkopf is called when entering the segmentkopf production.
	EnterSegmentkopf(c *SegmentkopfContext)

	// EnterSegmentkennung is called when entering the segmentkennung production.
	EnterSegmentkennung(c *SegmentkennungContext)

	// EnterSegmentnummer is called when entering the segmentnummer production.
	EnterSegmentnummer(c *SegmentnummerContext)

	// EnterSegmentversion is called when entering the segmentversion production.
	EnterSegmentversion(c *SegmentversionContext)

	// EnterBezugssegment is called when entering the bezugssegment production.
	EnterBezugssegment(c *BezugssegmentContext)

	// EnterNachrichtengroesse is called when entering the nachrichtengroesse production.
	EnterNachrichtengroesse(c *NachrichtengroesseContext)

	// EnterHbciVersion is called when entering the hbciVersion production.
	EnterHbciVersion(c *HbciVersionContext)

	// EnterDialogId is called when entering the dialogId production.
	EnterDialogId(c *DialogIdContext)

	// EnterNachrichtennummer is called when entering the nachrichtennummer production.
	EnterNachrichtennummer(c *NachrichtennummerContext)

	// EnterBezugsnachricht is called when entering the bezugsnachricht production.
	EnterBezugsnachricht(c *BezugsnachrichtContext)

	// EnterDe_sep is called when entering the de_sep production.
	EnterDe_sep(c *De_sepContext)

	// EnterDeg_sep is called when entering the deg_sep production.
	EnterDeg_sep(c *Deg_sepContext)

	// ExitNachrichtenkopf is called when exiting the nachrichtenkopf production.
	ExitNachrichtenkopf(c *NachrichtenkopfContext)

	// ExitSegmentkopf is called when exiting the segmentkopf production.
	ExitSegmentkopf(c *SegmentkopfContext)

	// ExitSegmentkennung is called when exiting the segmentkennung production.
	ExitSegmentkennung(c *SegmentkennungContext)

	// ExitSegmentnummer is called when exiting the segmentnummer production.
	ExitSegmentnummer(c *SegmentnummerContext)

	// ExitSegmentversion is called when exiting the segmentversion production.
	ExitSegmentversion(c *SegmentversionContext)

	// ExitBezugssegment is called when exiting the bezugssegment production.
	ExitBezugssegment(c *BezugssegmentContext)

	// ExitNachrichtengroesse is called when exiting the nachrichtengroesse production.
	ExitNachrichtengroesse(c *NachrichtengroesseContext)

	// ExitHbciVersion is called when exiting the hbciVersion production.
	ExitHbciVersion(c *HbciVersionContext)

	// ExitDialogId is called when exiting the dialogId production.
	ExitDialogId(c *DialogIdContext)

	// ExitNachrichtennummer is called when exiting the nachrichtennummer production.
	ExitNachrichtennummer(c *NachrichtennummerContext)

	// ExitBezugsnachricht is called when exiting the bezugsnachricht production.
	ExitBezugsnachricht(c *BezugsnachrichtContext)

	// ExitDe_sep is called when exiting the de_sep production.
	ExitDe_sep(c *De_sepContext)

	// ExitDeg_sep is called when exiting the deg_sep production.
	ExitDeg_sep(c *Deg_sepContext)
}
