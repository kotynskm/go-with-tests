package iteration

func RepeatChar(char string, num int) string {
	var repeatedString string
	for i := 0; i < num; i++ {
		repeatedString += char
	}
	return repeatedString
}