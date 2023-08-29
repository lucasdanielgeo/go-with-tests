package iteration

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}
	return repeated
}
