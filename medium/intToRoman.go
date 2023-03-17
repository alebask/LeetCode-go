package medium

import "strings"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func intToRoman(num int) string {

	rd := []struct {
		dec int
		rom string
	}{
		{dec: 1000, rom: "M"}, {dec: 900, rom: "CM"}, {dec: 500, rom: "D"},
		{dec: 400, rom: "CD"}, {dec: 100, rom: "C"}, {dec: 90, rom: "XC"},
		{dec: 50, rom: "L"}, {dec: 40, rom: "XL"}, {dec: 10, rom: "X"},
		{dec: 9, rom: "IX"}, {dec: 5, rom: "V"}, {dec: 4, rom: "IV"},
		{dec: 1, rom: "I"},
	}

	var res strings.Builder
	for _, p := range rd {
		for i := 0; i < num/p.dec; i++ {
			res.WriteString(p.rom)
		}
		num = num % p.dec
	}

	return res.String()
}
