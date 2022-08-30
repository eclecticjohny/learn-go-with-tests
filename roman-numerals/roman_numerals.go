package main

import "strings"

type ArabicRoman struct {
	Arabic int
	Roman  string
}

var rns = []ArabicRoman{
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

	for _, ar := range rns {
		for a >= ar.Arabic {
			result.WriteString(ar.Roman)
			a -= ar.Arabic
		}
	}

	return result.String()
}
