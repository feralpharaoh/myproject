package anagram

func IsAnagram(sourceText string, compText string) bool {
	sourceRunes := []rune(sourceText)
	compRunes := []rune(compText)

	if len(sourceText) != len(compText) {
		return false
	}

	sourceMap := make(map[rune]int)
	compMap := make(map[rune]int)

	for _, c := range sourceRunes {
		sourceMap[c]++
	}

	for _, c := range compRunes {
		compMap[c]++
	}

	for k := range sourceMap {
		if sourceMap[k] != compMap[k] {
			return false
		}
	}

	return true
}
