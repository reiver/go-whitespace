package whitespace



import (
	"testing"
)


func TestWhitespace(t *testing.T) {

	tests := []struct {
		R rune
		Expected bool
	}{
		{' '  , true},

		{'\a' , false}, // bell (== \007)
		{'\f' ,  true}, // form feed (== \014)
		{'\t' ,  true}, // horizontal tab (== \011)
		{'\n' ,  true}, // newline (== \012)
		{'\r' ,  true}, // carriage return (== \015)
		{'\v' ,  true}, // vertical tab character (== \013)

		{'a'  , false},
		{'A'  , false},
		{'x'  , false},
		{'X'  , false},
		{'z'  , false},
		{'Z'  , false},

		{'!'  , false},
		{'@'  , false},
		{'#'  , false},
		{'$'  , false},
		{'%'  , false},
		{'^'  , false},
		{'&'  , false},
		{'*'  , false},
		{'('  , false},
		{')'  , false},
		{'_'  , false},
		{'+'  , false},

		{'1'  , false},
		{'2'  , false},
		{'3'  , false},
		{'4'  , false},
		{'5'  , false},
		{'6'  , false},
		{'7'  , false},
		{'8'  , false},
		{'9'  , false},
		{'0'  , false},
		{'-'  , false},
		{'='  , false},

		{'\u0009', true}, // horizontal tab
		{'\u000A', true}, // line feed
		{'\u000B', true}, // vertical tab
		{'\u000C', true}, // form feed
		{'\u000D', true}, // carriage return
		{'\u0020', true}, // space
		{'\u0085', true}, // next line
		{'\u00A0', true}, // no-break space
		{'\u1680', true}, // ogham space mark
		{'\u180E', true}, // mongolian vowel separator
		{'\u2000', true}, // en quad
		{'\u2001', true}, // em quad
		{'\u2002', true}, // en space
		{'\u2003', true}, // em space
		{'\u2004', true}, // three-per-em space
		{'\u2005', true}, // four-per-em space
		{'\u2006', true}, // six-per-em space
		{'\u2007', true}, // figure space
		{'\u2008', true}, // punctuation space
		{'\u2009', true}, // thin space
		{'\u200A', true}, // hair space
		{'\u2028', true}, // line separator
		{'\u2029', true}, // paragraph separator
		{'\u202F', true}, // narrow no-break space
		{'\u205F', true}, // medium mathematical space
		{'\u3000', true}, // ideographic space

	}


	for testNumber, datum := range tests {

		expected := datum.Expected

		actual := IsWhitespace(datum.R)

		if expected != actual {
			t.Errorf("For test #%d, expected whitespace.IsWhitespace(%q = %d) = %t, but actually got whitespace.IsWhitespace(%q = %d) = %v.", testNumber, datum.R, datum.R, expected, datum.R, datum.R, actual)
		}

	}
}

