package decrypt

func Nimbus(str string) string {
	descryptedStr := ""

	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		descryptedStr += character

	}
	return descryptedStr
}
