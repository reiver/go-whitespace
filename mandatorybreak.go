package whitespace


// IsMandatoryBreak is a function tells you if a rune is a mandatory break character or not; it accounts for all 7 UNICODE mandatory break characters.
func IsMandatoryBreak(r rune) bool {

	result := false

	switch (r) {
		case
			'\u000A', // line feed
			'\u000B', // vertical tab
			'\u000C', // form feed
			'\u000D', // carriage return
			'\u0085', // next line
			'\u2028', // line separator
			'\u2029': // paragraph separator
			result = true
		default:
			result = false
	}

	return result
}
