package main

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{9, "IX"},
		{27, "XXVII"},
		{48, "XLVIII"},
		{59, "LIX"},
		{93, "XCIII"},
		{141, "CXLI"},
		{163, "CLXIII"},
		{402, "CDII"},
		{575, "DLXXV"},
		{911, "CMXI"},
		{1024, "MXXIV"},
		{3000, "MMM"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{4000, "MMMM"},
		{4500, "MMMMD"},
		{5000, "MMMMM"},
		{2014, "MMXIV"},
		{2022, "MMXXII"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ArabicToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}
