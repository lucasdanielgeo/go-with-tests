package romannumerals

import "strings"

// var numbersMap map[int16]string

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}
	return result.String()
}
