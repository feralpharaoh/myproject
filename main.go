package main

import (
	"fmt"

	"github.com/feralpharaoh/myproject/anagram"
	"github.com/feralpharaoh/myproject/palindrome"
	"github.com/feralpharaoh/myproject/strings"
)

func main() {
	//Calls upon the three separate packages with given inputs
	anagramTest := anagram.IsAnagram("feasts", "safest")
	palindromeTest := palindrome.IsPalindrome("racecar")
	finalText := strings.RemoveChar("HotPockets are the shiznit... super delicious!", "s")

	//Prints the results of the three separate tests
	fmt.Printf("Are the words 'feasts' and 'safest' an anagram of each other? --> %t\n", anagramTest)
	fmt.Printf("Is the word 'racecar' a palindrome? --> %t\n", palindromeTest)
	fmt.Printf("What is the resulting text after ommiting the given character? --> %s", finalText)
}
