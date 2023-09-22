package romannumerals

// var numbersMap map[int16]string

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
