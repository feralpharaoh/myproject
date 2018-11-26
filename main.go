package main

import (
	"fmt"

	"github.com/feralpharaoh/myproject/anagram"
)

func main() {
	//Calls upon the three separate packages with given inputs
	anagramTest := anagram.IsAnagram("feasts", "safest")
	//  palindromeTest = palindrome.IsPalindrome("racecar")
	//  finalText = strings.RemoveChar("HotPockets are the shiznit... super delicious!", "s")

	//Prints the results of the three separate tests
	fmt.Printf("Are these two words an anagram of each other? --> %t\n", anagramTest)
	//  fmt.Println("Is the provided word a palindrome? --> %t", palindromeTest)
	//  fmt.Println("What is the resulting text after ommiting the given character? --> %t", finalText)
}
