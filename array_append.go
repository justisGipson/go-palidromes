// Array Appending — O(N) Time, O(N) Space
//
// We can use very similar logic to our first attempt while improving our time complexity. Instead of concatenating a
// string, we can append to an array. Using the programming language Go, this is generally a constant time operation!
//
// So first we initialize an empty byte array called result. Then we iterate through str in reverse order (don’t forget
// to initialize i to len(str)-1!) and append each character to result.
// At the end, we simply compare our initial parameter str with string(result) to test if it is a palindrome or not.

package main

func isPalindrome(str string) bool {
	result := []byte{}

	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}

	return str == string(result)
}

// Since we only have a single for loop over str and the append() operation in Go is a constant time operation, we have
// a resulting time complexity of O(N). We still store an array that is the length of str, so our space complexity is
// also O(N).
