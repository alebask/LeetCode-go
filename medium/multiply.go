package medium

import "strings"

func multiply(num1 string, num2 string) string {

	if num1[0]-'0' == 0 || num2[0]-'0' == 0 {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)
	mult := make([]byte, len1+len2)

	k := 0
	var digit byte = 0
	for i := len2 - 1; i >= 0; i-- {
		digit = 0
		k = len2 - 1 - i
		for j := len1 - 1; j >= 0; j-- {
			n := (num2[i]-'0')*(num1[j]-'0') + digit
			digit = n / 10
			n = mult[k] + n%10
			digit += n / 10
			mult[k] = n % 10
			k++
		}
		if digit > 0 {
			mult[k] += digit
			k++
		}
	}

	var res strings.Builder
	for i := k - 1; i >= 0; i-- {
		res.WriteRune(rune(mult[i] + '0'))
	}

	return res.String()
}
