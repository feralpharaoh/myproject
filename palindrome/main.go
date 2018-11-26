package palindrome

func IsPalindrome(possiblePalindrome string) bool {
	testResult := true
	sourceRunes := []rune(possiblePalindrome)
	reverseRunes := []rune("")
	// create reverse of input string
	for i := len(sourceRunes) - 1; i >= 0; i-- {
		reverseRunes = append(reverseRunes, sourceRunes[i])
	}
	// compare source array to reverse array
	for j := 0; j < len(sourceRunes); j++ {
		if sourceRunes[j] != reverseRunes[j] {
			testResult = false
			break
		}
	}
	return testResult
}
