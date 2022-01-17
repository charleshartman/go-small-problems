/* 345. Reverse Vowels of a String

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.

Example 1:
Input: s = "hello"
Output: "holle"

Example 2:
Input: s = "leetcode"
Output: "leotcede"

Constraints:
1 <= s.length <= 3 * 105
s consist of printable ASCII characters.

*/

/* algorithm
- define key arr of aeiouAEIOU
- convert arg string to slice, split on ""
- define start and end pointers (0, len(sl) -1)
- outer loop, loop until start >= end
	- inner loop
		- move start until it encounters vowel s++
		- move end until it encounters vowel e--
		- swap elements at start and end indices
- return slice joined to str on ""
*/

package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	letters := strings.Split(s, "")
	start, end := 0, len(letters)-1

	for start <= end {
		for !strings.Contains(vowels, letters[start]) && start < end {
			start++
		}
		for !strings.Contains(vowels, letters[end]) && start < end {
			end--
		}
		letters[start], letters[end] = letters[end], letters[start]
		start++
		end--
	}

	return strings.Join(letters, "")
}

func main() {
	fmt.Println(reverseVowels("hello"))    // "holle"
	fmt.Println(reverseVowels("leetcode")) // "leotcede"
}
