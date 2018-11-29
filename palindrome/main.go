package palindrome

func IsPalindrome(possiblePalindrome string) bool {
	sourceRunes := []rune(possiblePalindrome)
	reverseRunes := make([]rune, len(sourceRunes))
	// create reverse of input string
	for i := len(sourceRunes) - 1; i >= 0; i-- {
		reverseRunes[len(sourceRunes)-1-i] = sourceRunes[i]
	}
	// compare source array to reverse array
	for j := range sourceRunes {
		if sourceRunes[j] != reverseRunes[j] {
			return false
		}
	}
	return true
}
