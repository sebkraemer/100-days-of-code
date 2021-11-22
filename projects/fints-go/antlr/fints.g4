grammar fints;


nachrichtenkopf:
	segmentkopf de_sep nachrichtengroesse de_sep hbciVersion de_sep dialogId de_sep
		nachrichtennummer (de_sep bezugsnachricht)?;

segmentkopf:
	segmentkennung deg_sep segmentnummer deg_sep segmentversion (
		deg_sep bezugssegment
	)?;

DT_AN:
	FRAG_ENTWERTET_SYNTAXZEICHEN
	| FRAG_ALPHANUMERISCH
	| FRAG_SYNTAXZEICHEN;

fragment FRAG_ALPHANUMERISCH: [A-Za-z0-9+:'?@];
fragment FRAG_SYNTAXZEICHEN:
	[!"#$%&()*+,\-./[\\\]^_]; // \x7B\x7D\xA7 fragment FRAG_ENTWERTET_SYNTAXZEICHEN: '?' [+:'?@];
fragment FRAG_ENTWERTET_SYNTAXZEICHEN: '?' [+:'?@];

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
