// String Concatenation — O(N²) Time, O(N) Space
//
// A common brute force solution would be to recreate our given string in reverse order, and then compare the two
// strings to see if they are equal to each other. These initial solutions are actually a great first step in coding
// questions and show your interviewer that you think of the basics first before jumping to optimization.
//
// First we initialize an empty string we call reversedStr. We will then execute a for loop where we initialize the
// index i to the end of the array minus one (array of length 30 has max index of 29). As we move i from the end of the
// string to the beginning in reverse order, we concatenate the current character str[i] with reversedStr.
// Note: Accessing an element of an array with an index in Go will give you a byte object, so you have to convert it to
// a string first string(str[i]).
//
// Last we have another for loop in which we use the function range() to have our index i iterate through the length of
// the string str. As we do so, we compare each character if str and reversedS and check if they are equal. If they
// aren’t at any point in time, we return false. If we make it through the entire string and all of the characters are
// equal, then the string is in fact a palindrome and we return true!

package main

func isPalindrome(str string) bool {
	reversedString := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}

	for i := range str {
		if str[i] != reversedString[i] {
			return false
		}
	}
	return true
}
