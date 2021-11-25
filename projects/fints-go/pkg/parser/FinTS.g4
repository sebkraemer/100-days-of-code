grammar FinTS;

// options {
// 	tokenVocab = FinTSLexerRules;
// }

nachrichtenkopf:
	segmentkopf de_sep nachrichtengroesse de_sep hbciVersion de_sep dialogId de_sep
		nachrichtennummer (de_sep bezugsnachricht)?;

segmentkopf:
	segmentkennung deg_sep segmentnummer deg_sep segmentversion (
		deg_sep bezugssegment
	)?;

segmentkennung: (DT_AN)+?;
segmentnummer: (DT_AN)+?; // DT_num;
segmentversion: (DT_AN)+?; // DT_num;
bezugssegment: (DT_AN)+?; // DT_num;

nachrichtengroesse: (DT_AN)+?; // DT_num;
hbciVersion: (DT_AN)+?; // DT_num;
dialogId: (DT_AN)+?;
nachrichtennummer: (DT_AN)+?; // DT_num;
bezugsnachricht: dialogId deg_sep nachrichtennummer;

// H.1.1 Syntaxzeichen
de_sep: '+';
deg_sep: ':';
