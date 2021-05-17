// Opposite Indices — O(N) Time, O(1) Space
//
// We can do better! Let’s think for a moment. We’ve solved the problem with our first two attempts while creating\
// additional data structures of the str built in reverse order. Do we have to do this? What’s in the very nature of a
// palindrome? It’s that the reflection of the word is the word itself — even more specific, the first character of the
// string is equal to the last character of the string. The second character of the string is equal to the second to
// last character of the string, and so on…
//
// So instead of building the string in reverse order, why not just check if the very pattern of the palindrome exists
// in the string? We can do so by initializing two indices at the start and end of the string, and then incrementally
// check the string making sure the left and right sides of the string are equal!
// Here we just need to be careful with our indexing. We start a for loop as expected with i=0 and increment through
// the string. As we do so, we set another variable j=len(str)-1-i. So when i is 0, it is the first index and j will be
// the last index. When i is 1 it will be the second index, and j will be the second to last index, and so on.

package main

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// We have a single for loop that iterates through the string with constant time operations within. So our run-time
// complexity is O(N). Now that we don’t store any data structures, we have optimized our algorithm to have constant
// space complexity O(1).
//
// If you wanted one final optimization, you could terminate the for loop when i > (len(str)/2)+1, since we only need
// each half of the string to equal each other.
