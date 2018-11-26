package strings

func RemoveChar(sourceText string, unwantedChar string) string {
	sourceRunes := []rune(sourceText)
	removeRunes := []rune(unwantedChar)
	remove := removeRunes[0]
	finalText := []rune(nil)

	//check each character from sourceText and add to result if not selected character
	for i := 0; i < len(sourceText); i++ {
		currChar := sourceRunes[i]
		if currChar != remove { // if this character does not match character selected for removal, add it to the new slice
			finalText = append(finalText, currChar)
		}
	}
	return string(finalText)
}
