package whitespace



import (
	"testing"
)


func TestIsMandatoryBreak(t *testing.T) {

	tests := []struct {
		R rune
		Expected bool
	}{
		{' '  , false},

		{'\a' , false},  // bell (== \007)
		{'\f' ,  true},  // form feed (== \014)
		{'\t' ,  false}, // horizontal tab (== \011)
		{'\n' ,  true},  // newline (== \012)
		{'\r' ,  true},  // carriage return (== \015)
		{'\v' ,  true},  // vertical tab character (== \013)

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

		{'\u0009', false}, // horizontal tab
		{'\u000A', true},  // line feed
		{'\u000B', true},  // vertical tab
		{'\u000C', true},  // form feed
		{'\u000D', true},  // carriage return
		{'\u0020', false}, // space
		{'\u0085', true},  // next line
		{'\u00A0', false}, // no-break space
		{'\u1680', false}, // ogham space mark
		{'\u180E', false}, // mongolian vowel separator
		{'\u2000', false}, // en quad
		{'\u2001', false}, // em quad
		{'\u2002', false}, // en space
		{'\u2003', false}, // em space
		{'\u2004', false}, // three-per-em space
		{'\u2005', false}, // four-per-em space
		{'\u2006', false}, // six-per-em space
		{'\u2007', false}, // figure space
		{'\u2008', false}, // punctuation space
		{'\u2009', false}, // thin space
		{'\u200A', false}, // hair space
		{'\u2028', true},  // line separator
		{'\u2029', true},  // paragraph separator
		{'\u202F', false}, // narrow no-break space
		{'\u205F', false}, // medium mathematical space
		{'\u3000', false}, // ideographic space
	}


	for testNumber, datum := range tests {

		expected := datum.Expected

		actual := IsMandatoryBreak(datum.R)

		if expected != actual {
			t.Errorf("For test #%d, expected whitespace.IsMandatoryBreak(%q = %d) = %t, but actually got whitespace.IsMandatoryBreak(%q = %d) = %v.", testNumber, datum.R, datum.R, expected, datum.R, datum.R, actual)
		}

	}
}

