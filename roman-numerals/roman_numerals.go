package main

import "strings"

type ArabicRoman struct {
	Arabic int
	Roman  string
}

func ArabicToRoman(a int) string {
	ars := []ArabicRoman{
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

	var result strings.Builder

	for _, ar := range ars {
		for a >= ar.Arabic {
			result.WriteString(ar.Roman)
			a -= ar.Arabic
		}
	}

	return result.String()
}

func RomanToArabic(r string) int {
	ras := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	rl := len(r)

	if rl == 0 {
		return 0
	}

	if rl == 1 {
		return ras[r[0]]
	}

	sum := ras[r[rl-1]]

	for i := rl - 2; i >= 0; i-- {
		if ras[r[i]] < ras[r[i+1]] {
			sum -= ras[r[i]]
		} else {
			sum += ras[r[i]]
		}
	}

	return sum
}
