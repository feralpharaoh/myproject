package anagram

func IsAnagram(sourceText string, compText string) bool {

	testResult := true
	sourceRunes := []rune(sourceText)
	compRunes := []rune(compText)
	//Compares length of each input strings
	lengthCheck := len(sourceText) == len(compText)

	//If strings are equal length, create maps of characters and compare strings' key values
	if lengthCheck == true {

		//initialize source and comparison string maps
		sourceMap := make(map[rune]int)
		compMap := make(map[rune]int)

		// create map of sourceRunes with number of each character
		for i := 0; i < len(sourceRunes); i++ { //iterate characters over the string
			currentChar := sourceRunes[i] //pull current character
			exist := sourceMap[currentChar]
			if exist == 0 { //if this character is not already represented in the map, count how many exist and add key to map
				// count how many times this character shows up in the string
				count := 0
				for j := 0; j < len(sourceRunes); j++ {
					if sourceRunes[j] == currentChar {
						count++
					}
				}
				// add key and value to map
				sourceMap[currentChar] = count
			}
		}

		// while creating comparison Map, check for comparison to source Map
		for i := 0; i < len(compRunes); i++ { //iterate characters over the comparison string
			currentChar := compRunes[i] //pull current character
			existInSource := sourceMap[currentChar]
			if existInSource != 0 { // if this character exists in sourceMap, proceed with test
				exist := compMap[currentChar]
				if exist == 0 { //if this character is not already represented in the map, count how many exist and add key to map
					// count how many times this character shows up in the string
					count := 0
					for j := 0; j < len(compRunes); j++ {
						if compRunes[j] == currentChar {
							count++
						}
					}
					// add key and value to map if value of character matches sourceMap, else false and break
					sourceVal := sourceMap[currentChar]
					if sourceVal == count {
						compMap[currentChar] = count
					} else {
						testResult = false
						break
					}
				}

			} else { // if current character does not exist in sourceMap, test results in false and break
				testResult = false
				break
			}
		}
	} else { //If strings are not equal length, return false
		testResult = false
	}

	return testResult
}
