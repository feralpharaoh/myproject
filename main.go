package main

import (
	"fmt"

	"github.com/feralpharaoh/myproject/anagram"
	"github.com/feralpharaoh/myproject/palindrome"
)

func main() {
	//Calls upon the three separate packages with given inputs
	anagramTest := anagram.IsAnagram("feasts", "safest")
	palindromeTest = palindrome.IsPalindrome("racecar")
	//  finalText = strings.RemoveChar("HotPockets are the shiznit... super delicious!", "s")

	//Prints the results of the three separate tests
	fmt.Printf("Are the words 'feasts' and 'safest' an anagram of each other? --> %t\n", anagramTest)
	fmt.Println("Is the word 'racecar' a palindrome? --> %t", palindromeTest)
	//  fmt.Println("What is the resulting text after ommiting the given character? --> %t", finalText)
}
