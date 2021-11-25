lexer grammar FinTSLexerRules;

DT_AN:
	FRAG_ENTWERTET_SYNTAXZEICHEN
	| FRAG_ALPHANUMERISCH
	| FRAG_SYNTAXZEICHEN;

BINARY_START:
	'@' -> pushMode(BINARYMODE); // error lexical modes are only allowed in lexer

fragment FRAG_ALPHANUMERISCH: [A-Za-z0-9+:'?@];
fragment FRAG_SYNTAXZEICHEN:
	[!"#$%&()*+,\-./[\\\]^_]; // \x7B\x7D\xA7 fragment FRAG_ENTWERTET_SYNTAXZEICHEN: '?' [+:'?@];
fragment FRAG_ENTWERTET_SYNTAXZEICHEN: '?' [+:'?@];

mode BINARYMODE;
BINARY_LENGTH: [0-9]+? '@';
BINARY: '\u0000' ..'\u00FF'+ -> popMode;