package strings

func RemoveChar(sourceText string, unwantedChar string) string {
	sourceRunes := []rune(sourceText)
	removeRunes := []rune(unwantedChar)
	remove := removeRunes[0]
	finalText := []rune{}

	//check each character from sourceText and add to result if not selected character
	for _, c := range sourceRunes {
		if c != remove { // if this character does not match character selected for removal, add it to the new slice
			finalText = append(finalText, c)
		}
	}
	return string(finalText)
}

func ReplaceChar(sourceText string, charToSwap rune, charToUse rune) string {
	sourceRunes := []rune(sourceText)
	for i, c := range sourceRunes {
		if c == charToSwap {
			sourceRunes[i] = charToUse
		}
	}
	return string(sourceRunes)
}
