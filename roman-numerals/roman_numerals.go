package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var rns = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ArabicToRoman(a int) string {
	var result strings.Builder

	for _, rn := range rns {
		for a >= rn.Value {
			result.WriteString(rn.Symbol)
			a -= rn.Value
		}
	}

	return result.String()
}
