package gsm

// taken from http://www.unicode.org/Public/MAPPINGS/ETSI/GSM0338.TXT

var escDecode = map[byte]*utf8Enc{
	0x0A: {1, [3]byte{0x0c, 0x00, 0x00}}, // FORM FEED
	0x14: {1, [3]byte{0x5e, 0x00, 0x00}}, // CIRCUMFLEX ACCENT
	0x28: {1, [3]byte{0x7b, 0x00, 0x00}}, // LEFT CURLY BRACKET
	0x29: {1, [3]byte{0x7d, 0x00, 0x00}}, // RIGHT CURLY BRACKET
	0x2F: {1, [3]byte{0x5c, 0x00, 0x00}}, // REVERSE SOLIDUS
	0x3C: {1, [3]byte{0x5b, 0x00, 0x00}}, // LEFT SQUARE BRACKET
	0x3D: {1, [3]byte{0x7e, 0x00, 0x00}}, // TILDE
	0x3E: {1, [3]byte{0x5d, 0x00, 0x00}}, // RIGHT SQUARE BRACKET
	0x40: {1, [3]byte{0x7c, 0x00, 0x00}}, // VERTICAL LINE
	0x65: {3, [3]byte{0xe2, 0x82, 0xac}}, // EURO SIGN
}

var escEncode = [10]uint32{
	0x1B0A000C, // FORM FEED
	0x1B14005E, // CIRCUMFLEX ACCENT
	0x1B28007B, // LEFT CURLY BRACKET
	0x1B29007D, // RIGHT CURLY BRACKET
	0x1B2F005C, // REVERSE SOLIDUS
	0x1B3C005B, // LEFT SQUARE BRACKET
	0x1B3D007E, // TILDE
	0x1B3E005D, // RIGHT SQUARE BRACKET
	0x1B40007C, // VERTICAL LINE
	0x1B6520AC, // EURO SIGN
}

// utf8Enc holds a rune's UTF-8 encoding in data[:len].
type utf8Enc struct {
	len  uint8
	data [3]byte
}

var decode = [128]utf8Enc{
	{1, [3]byte{0x40, 0x00, 0x00}}, // COMMERCIAL AT
	{2, [3]byte{0xc2, 0xa3, 0x00}}, // POUND SIGN
	{1, [3]byte{0x24, 0x00, 0x00}}, // DOLLAR SIGN
	{2, [3]byte{0xc2, 0xa5, 0x00}}, // YEN SIGN
	{2, [3]byte{0xc3, 0xa8, 0x00}}, // LATIN SMALL LETTER E WITH GRAVE
	{2, [3]byte{0xc3, 0xa9, 0x00}}, // LATIN SMALL LETTER E WITH ACUTE
	{2, [3]byte{0xc3, 0xb9, 0x00}}, // LATIN SMALL LETTER U WITH GRAVE
	{2, [3]byte{0xc3, 0xac, 0x00}}, // LATIN SMALL LETTER I WITH GRAVE
	{2, [3]byte{0xc3, 0xb2, 0x00}}, // LATIN SMALL LETTER O WITH GRAVE
	{2, [3]byte{0xc3, 0xa7, 0x00}}, // LATIN SMALL LETTER C WITH CEDILLA
	{1, [3]byte{0x0a, 0x00, 0x00}}, // LINE FEED
	{2, [3]byte{0xc3, 0x98, 0x00}}, // LATIN CAPITAL LETTER O WITH STROKE
	{2, [3]byte{0xc3, 0xb8, 0x00}}, // LATIN SMALL LETTER O WITH STROKE
	{1, [3]byte{0x0d, 0x00, 0x00}}, // CARRIAGE RETURN
	{2, [3]byte{0xc3, 0x85, 0x00}}, // LATIN CAPITAL LETTER A WITH RING ABOVE
	{2, [3]byte{0xc3, 0xa5, 0x00}}, // LATIN SMALL LETTER A WITH RING ABOVE
	{2, [3]byte{0xce, 0x94, 0x00}}, // GREEK CAPITAL LETTER DELTA
	{1, [3]byte{0x5f, 0x00, 0x00}}, // LOW LINE
	{2, [3]byte{0xce, 0xa6, 0x00}}, // GREEK CAPITAL LETTER PHI
	{2, [3]byte{0xce, 0x93, 0x00}}, // GREEK CAPITAL LETTER GAMMA
	{2, [3]byte{0xce, 0x9b, 0x00}}, // GREEK CAPITAL LETTER LAMDA
	{2, [3]byte{0xce, 0xa9, 0x00}}, // GREEK CAPITAL LETTER OMEGA
	{2, [3]byte{0xce, 0xa0, 0x00}}, // GREEK CAPITAL LETTER PI
	{2, [3]byte{0xce, 0xa8, 0x00}}, // GREEK CAPITAL LETTER PSI
	{2, [3]byte{0xce, 0xa3, 0x00}}, // GREEK CAPITAL LETTER SIGMA
	{2, [3]byte{0xce, 0x98, 0x00}}, // GREEK CAPITAL LETTER THETA
	{2, [3]byte{0xce, 0x9e, 0x00}}, // GREEK CAPITAL LETTER XI
	{2, [3]byte{0xc2, 0xa0, 0x00}}, // ESCAPE TO EXTENSION TABLE (or displayed as NBSP, see note above)
	{2, [3]byte{0xc3, 0x86, 0x00}}, // LATIN CAPITAL LETTER AE
	{2, [3]byte{0xc3, 0xa6, 0x00}}, // LATIN SMALL LETTER AE
	{2, [3]byte{0xc3, 0x9f, 0x00}}, // LATIN SMALL LETTER SHARP S (German)
	{2, [3]byte{0xc3, 0x89, 0x00}}, // LATIN CAPITAL LETTER E WITH ACUTE
	{1, [3]byte{0x20, 0x00, 0x00}}, // SPACE
	{1, [3]byte{0x21, 0x00, 0x00}}, // EXCLAMATION MARK
	{1, [3]byte{0x22, 0x00, 0x00}}, // QUOTATION MARK
	{1, [3]byte{0x23, 0x00, 0x00}}, // NUMBER SIGN
	{2, [3]byte{0xc2, 0xa4, 0x00}}, // CURRENCY SIGN
	{1, [3]byte{0x25, 0x00, 0x00}}, // PERCENT SIGN
	{1, [3]byte{0x26, 0x00, 0x00}}, // AMPERSAND
	{1, [3]byte{0x27, 0x00, 0x00}}, // APOSTROPHE
	{1, [3]byte{0x28, 0x00, 0x00}}, // LEFT PARENTHESIS
	{1, [3]byte{0x29, 0x00, 0x00}}, // RIGHT PARENTHESIS
	{1, [3]byte{0x2a, 0x00, 0x00}}, // ASTERISK
	{1, [3]byte{0x2b, 0x00, 0x00}}, // PLUS SIGN
	{1, [3]byte{0x2c, 0x00, 0x00}}, // COMMA
	{1, [3]byte{0x2d, 0x00, 0x00}}, // HYPHEN-MINUS
	{1, [3]byte{0x2e, 0x00, 0x00}}, // FULL STOP
	{1, [3]byte{0x2f, 0x00, 0x00}}, // SOLIDUS
	{1, [3]byte{0x30, 0x00, 0x00}}, // DIGIT ZERO
	{1, [3]byte{0x31, 0x00, 0x00}}, // DIGIT ONE
	{1, [3]byte{0x32, 0x00, 0x00}}, // DIGIT TWO
	{1, [3]byte{0x33, 0x00, 0x00}}, // DIGIT THREE
	{1, [3]byte{0x34, 0x00, 0x00}}, // DIGIT FOUR
	{1, [3]byte{0x35, 0x00, 0x00}}, // DIGIT FIVE
	{1, [3]byte{0x36, 0x00, 0x00}}, // DIGIT SIX
	{1, [3]byte{0x37, 0x00, 0x00}}, // DIGIT SEVEN
	{1, [3]byte{0x38, 0x00, 0x00}}, // DIGIT EIGHT
	{1, [3]byte{0x39, 0x00, 0x00}}, // DIGIT NINE
	{1, [3]byte{0x3a, 0x00, 0x00}}, // COLON
	{1, [3]byte{0x3b, 0x00, 0x00}}, // SEMICOLON
	{1, [3]byte{0x3c, 0x00, 0x00}}, // LESS-THAN SIGN
	{1, [3]byte{0x3d, 0x00, 0x00}}, // EQUALS SIGN
	{1, [3]byte{0x3e, 0x00, 0x00}}, // GREATER-THAN SIGN
	{1, [3]byte{0x3f, 0x00, 0x00}}, // QUESTION MARK
	{2, [3]byte{0xc2, 0xa1, 0x00}}, // INVERTED EXCLAMATION MARK
	{1, [3]byte{0x41, 0x00, 0x00}}, // LATIN CAPITAL LETTER A
	{1, [3]byte{0x42, 0x00, 0x00}}, // LATIN CAPITAL LETTER B
	{1, [3]byte{0x43, 0x00, 0x00}}, // LATIN CAPITAL LETTER C
	{1, [3]byte{0x44, 0x00, 0x00}}, // LATIN CAPITAL LETTER D
	{1, [3]byte{0x45, 0x00, 0x00}}, // LATIN CAPITAL LETTER E
	{1, [3]byte{0x46, 0x00, 0x00}}, // LATIN CAPITAL LETTER F
	{1, [3]byte{0x47, 0x00, 0x00}}, // LATIN CAPITAL LETTER G
	{1, [3]byte{0x48, 0x00, 0x00}}, // LATIN CAPITAL LETTER H
	{1, [3]byte{0x49, 0x00, 0x00}}, // LATIN CAPITAL LETTER I
	{1, [3]byte{0x4a, 0x00, 0x00}}, // LATIN CAPITAL LETTER J
	{1, [3]byte{0x4b, 0x00, 0x00}}, // LATIN CAPITAL LETTER K
	{1, [3]byte{0x4c, 0x00, 0x00}}, // LATIN CAPITAL LETTER L
	{1, [3]byte{0x4d, 0x00, 0x00}}, // LATIN CAPITAL LETTER M
	{1, [3]byte{0x4e, 0x00, 0x00}}, // LATIN CAPITAL LETTER N
	{1, [3]byte{0x4f, 0x00, 0x00}}, // LATIN CAPITAL LETTER O
	{1, [3]byte{0x50, 0x00, 0x00}}, // LATIN CAPITAL LETTER P
	{1, [3]byte{0x51, 0x00, 0x00}}, // LATIN CAPITAL LETTER Q
	{1, [3]byte{0x52, 0x00, 0x00}}, // LATIN CAPITAL LETTER R
	{1, [3]byte{0x53, 0x00, 0x00}}, // LATIN CAPITAL LETTER S
	{1, [3]byte{0x54, 0x00, 0x00}}, // LATIN CAPITAL LETTER T
	{1, [3]byte{0x55, 0x00, 0x00}}, // LATIN CAPITAL LETTER U
	{1, [3]byte{0x56, 0x00, 0x00}}, // LATIN CAPITAL LETTER V
	{1, [3]byte{0x57, 0x00, 0x00}}, // LATIN CAPITAL LETTER W
	{1, [3]byte{0x58, 0x00, 0x00}}, // LATIN CAPITAL LETTER X
	{1, [3]byte{0x59, 0x00, 0x00}}, // LATIN CAPITAL LETTER Y
	{1, [3]byte{0x5a, 0x00, 0x00}}, // LATIN CAPITAL LETTER Z
	{2, [3]byte{0xc3, 0x84, 0x00}}, // LATIN CAPITAL LETTER A WITH DIAERESIS
	{2, [3]byte{0xc3, 0x96, 0x00}}, // LATIN CAPITAL LETTER O WITH DIAERESIS
	{2, [3]byte{0xc3, 0x91, 0x00}}, // LATIN CAPITAL LETTER N WITH TILDE
	{2, [3]byte{0xc3, 0x9c, 0x00}}, // LATIN CAPITAL LETTER U WITH DIAERESIS
	{2, [3]byte{0xc2, 0xa7, 0x00}}, // SECTION SIGN
	{2, [3]byte{0xc2, 0xbf, 0x00}}, // INVERTED QUESTION MARK
	{1, [3]byte{0x61, 0x00, 0x00}}, // LATIN SMALL LETTER A
	{1, [3]byte{0x62, 0x00, 0x00}}, // LATIN SMALL LETTER B
	{1, [3]byte{0x63, 0x00, 0x00}}, // LATIN SMALL LETTER C
	{1, [3]byte{0x64, 0x00, 0x00}}, // LATIN SMALL LETTER D
	{1, [3]byte{0x65, 0x00, 0x00}}, // LATIN SMALL LETTER E
	{1, [3]byte{0x66, 0x00, 0x00}}, // LATIN SMALL LETTER F
	{1, [3]byte{0x67, 0x00, 0x00}}, // LATIN SMALL LETTER G
	{1, [3]byte{0x68, 0x00, 0x00}}, // LATIN SMALL LETTER H
	{1, [3]byte{0x69, 0x00, 0x00}}, // LATIN SMALL LETTER I
	{1, [3]byte{0x6a, 0x00, 0x00}}, // LATIN SMALL LETTER J
	{1, [3]byte{0x6b, 0x00, 0x00}}, // LATIN SMALL LETTER K
	{1, [3]byte{0x6c, 0x00, 0x00}}, // LATIN SMALL LETTER L
	{1, [3]byte{0x6d, 0x00, 0x00}}, // LATIN SMALL LETTER M
	{1, [3]byte{0x6e, 0x00, 0x00}}, // LATIN SMALL LETTER N
	{1, [3]byte{0x6f, 0x00, 0x00}}, // LATIN SMALL LETTER O
	{1, [3]byte{0x70, 0x00, 0x00}}, // LATIN SMALL LETTER P
	{1, [3]byte{0x71, 0x00, 0x00}}, // LATIN SMALL LETTER Q
	{1, [3]byte{0x72, 0x00, 0x00}}, // LATIN SMALL LETTER R
	{1, [3]byte{0x73, 0x00, 0x00}}, // LATIN SMALL LETTER S
	{1, [3]byte{0x74, 0x00, 0x00}}, // LATIN SMALL LETTER T
	{1, [3]byte{0x75, 0x00, 0x00}}, // LATIN SMALL LETTER U
	{1, [3]byte{0x76, 0x00, 0x00}}, // LATIN SMALL LETTER V
	{1, [3]byte{0x77, 0x00, 0x00}}, // LATIN SMALL LETTER W
	{1, [3]byte{0x78, 0x00, 0x00}}, // LATIN SMALL LETTER X
	{1, [3]byte{0x79, 0x00, 0x00}}, // LATIN SMALL LETTER Y
	{1, [3]byte{0x7a, 0x00, 0x00}}, // LATIN SMALL LETTER Z
	{2, [3]byte{0xc3, 0xa4, 0x00}}, // LATIN SMALL LETTER A WITH DIAERESIS
	{2, [3]byte{0xc3, 0xb6, 0x00}}, // LATIN SMALL LETTER O WITH DIAERESIS
	{2, [3]byte{0xc3, 0xb1, 0x00}}, // LATIN SMALL LETTER N WITH TILDE
	{2, [3]byte{0xc3, 0xbc, 0x00}}, // LATIN SMALL LETTER U WITH DIAERESIS
	{2, [3]byte{0xc3, 0xa0, 0x00}}, // LATIN SMALL LETTER A WITH GRAVE

}

var encode = [128]uint32{
	0x0A00000A, // LINE FEED
	0x0D00000D, // CARRIAGE RETURN
	0x20000020, // SPACE
	0x21000021, // EXCLAMATION MARK
	0x22000022, // QUOTATION MARK
	0x23000023, // NUMBER SIGN
	0x02000024, // DOLLAR SIGN
	0x25000025, // PERCENT SIGN
	0x26000026, // AMPERSAND
	0x27000027, // APOSTROPHE
	0x28000028, // LEFT PARENTHESIS
	0x29000029, // RIGHT PARENTHESIS
	0x2A00002A, // ASTERISK
	0x2B00002B, // PLUS SIGN
	0x2C00002C, // COMMA
	0x2D00002D, // HYPHEN-MINUS
	0x2E00002E, // FULL STOP
	0x2F00002F, // SOLIDUS
	0x30000030, // DIGIT ZERO
	0x31000031, // DIGIT ONE
	0x32000032, // DIGIT TWO
	0x33000033, // DIGIT THREE
	0x34000034, // DIGIT FOUR
	0x35000035, // DIGIT FIVE
	0x36000036, // DIGIT SIX
	0x37000037, // DIGIT SEVEN
	0x38000038, // DIGIT EIGHT
	0x39000039, // DIGIT NINE
	0x3A00003A, // COLON
	0x3B00003B, // SEMICOLON
	0x3C00003C, // LESS-THAN SIGN
	0x3D00003D, // EQUALS SIGN
	0x3E00003E, // GREATER-THAN SIGN
	0x3F00003F, // QUESTION MARK
	0x00000040, // COMMERCIAL AT
	0x41000041, // LATIN CAPITAL LETTER A
	0x42000042, // LATIN CAPITAL LETTER B
	0x43000043, // LATIN CAPITAL LETTER C
	0x44000044, // LATIN CAPITAL LETTER D
	0x45000045, // LATIN CAPITAL LETTER E
	0x46000046, // LATIN CAPITAL LETTER F
	0x47000047, // LATIN CAPITAL LETTER G
	0x48000048, // LATIN CAPITAL LETTER H
	0x49000049, // LATIN CAPITAL LETTER I
	0x4A00004A, // LATIN CAPITAL LETTER J
	0x4B00004B, // LATIN CAPITAL LETTER K
	0x4C00004C, // LATIN CAPITAL LETTER L
	0x4D00004D, // LATIN CAPITAL LETTER M
	0x4E00004E, // LATIN CAPITAL LETTER N
	0x4F00004F, // LATIN CAPITAL LETTER O
	0x50000050, // LATIN CAPITAL LETTER P
	0x51000051, // LATIN CAPITAL LETTER Q
	0x52000052, // LATIN CAPITAL LETTER R
	0x53000053, // LATIN CAPITAL LETTER S
	0x54000054, // LATIN CAPITAL LETTER T
	0x55000055, // LATIN CAPITAL LETTER U
	0x56000056, // LATIN CAPITAL LETTER V
	0x57000057, // LATIN CAPITAL LETTER W
	0x58000058, // LATIN CAPITAL LETTER X
	0x59000059, // LATIN CAPITAL LETTER Y
	0x5A00005A, // LATIN CAPITAL LETTER Z
	0x1100005F, // LOW LINE
	0x61000061, // LATIN SMALL LETTER A
	0x62000062, // LATIN SMALL LETTER B
	0x63000063, // LATIN SMALL LETTER C
	0x64000064, // LATIN SMALL LETTER D
	0x65000065, // LATIN SMALL LETTER E
	0x66000066, // LATIN SMALL LETTER F
	0x67000067, // LATIN SMALL LETTER G
	0x68000068, // LATIN SMALL LETTER H
	0x69000069, // LATIN SMALL LETTER I
	0x6A00006A, // LATIN SMALL LETTER J
	0x6B00006B, // LATIN SMALL LETTER K
	0x6C00006C, // LATIN SMALL LETTER L
	0x6D00006D, // LATIN SMALL LETTER M
	0x6E00006E, // LATIN SMALL LETTER N
	0x6F00006F, // LATIN SMALL LETTER O
	0x70000070, // LATIN SMALL LETTER P
	0x71000071, // LATIN SMALL LETTER Q
	0x72000072, // LATIN SMALL LETTER R
	0x73000073, // LATIN SMALL LETTER S
	0x74000074, // LATIN SMALL LETTER T
	0x75000075, // LATIN SMALL LETTER U
	0x76000076, // LATIN SMALL LETTER V
	0x77000077, // LATIN SMALL LETTER W
	0x78000078, // LATIN SMALL LETTER X
	0x79000079, // LATIN SMALL LETTER Y
	0x7A00007A, // LATIN SMALL LETTER Z
	0x1B0000A0, // ESCAPE TO EXTENSION TABLE (or displayed as NBSP, see note above)
	0x400000A1, // INVERTED EXCLAMATION MARK
	0x010000A3, // POUND SIGN
	0x240000A4, // CURRENCY SIGN
	0x030000A5, // YEN SIGN
	0x5F0000A7, // SECTION SIGN
	0x600000BF, // INVERTED QUESTION MARK
	0x5B0000C4, // LATIN CAPITAL LETTER A WITH DIAERESIS
	0x0E0000C5, // LATIN CAPITAL LETTER A WITH RING ABOVE
	0x1C0000C6, // LATIN CAPITAL LETTER AE
	0x1F0000C9, // LATIN CAPITAL LETTER E WITH ACUTE
	0x5D0000D1, // LATIN CAPITAL LETTER N WITH TILDE
	0x5C0000D6, // LATIN CAPITAL LETTER O WITH DIAERESIS
	0x0B0000D8, // LATIN CAPITAL LETTER O WITH STROKE
	0x5E0000DC, // LATIN CAPITAL LETTER U WITH DIAERESIS
	0x1E0000DF, // LATIN SMALL LETTER SHARP S (German)
	0x7F0000E0, // LATIN SMALL LETTER A WITH GRAVE
	0x7B0000E4, // LATIN SMALL LETTER A WITH DIAERESIS
	0x0F0000E5, // LATIN SMALL LETTER A WITH RING ABOVE
	0x1D0000E6, // LATIN SMALL LETTER AE
	0x090000E7, // LATIN SMALL LETTER C WITH CEDILLA
	0x040000E8, // LATIN SMALL LETTER E WITH GRAVE
	0x050000E9, // LATIN SMALL LETTER E WITH ACUTE
	0x070000EC, // LATIN SMALL LETTER I WITH GRAVE
	0x7D0000F1, // LATIN SMALL LETTER N WITH TILDE
	0x080000F2, // LATIN SMALL LETTER O WITH GRAVE
	0x7C0000F6, // LATIN SMALL LETTER O WITH DIAERESIS
	0x0C0000F8, // LATIN SMALL LETTER O WITH STROKE
	0x060000F9, // LATIN SMALL LETTER U WITH GRAVE
	0x7E0000FC, // LATIN SMALL LETTER U WITH DIAERESIS
	0x13000393, // GREEK CAPITAL LETTER GAMMA
	0x10000394, // GREEK CAPITAL LETTER DELTA
	0x19000398, // GREEK CAPITAL LETTER THETA
	0x1400039B, // GREEK CAPITAL LETTER LAMDA
	0x1A00039E, // GREEK CAPITAL LETTER XI
	0x160003A0, // GREEK CAPITAL LETTER PI
	0x180003A3, // GREEK CAPITAL LETTER SIGMA
	0x120003A6, // GREEK CAPITAL LETTER PHI
	0x170003A8, // GREEK CAPITAL LETTER PSI
	0x150003A9, // GREEK CAPITAL LETTER OMEGA
}