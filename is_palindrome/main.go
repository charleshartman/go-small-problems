/*

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.

*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
base case is when length of string <= 1
  return true
  empty string and string of 1 char is a palindrome
base case where the length is string == 2
  return str[0] == str[1]
if str[0] == str[len(str) - 1]
  isPalindrome(str[1:len(str)-1])

zyxyz
zyxxyz  str[1:len(string)-1]

*/

// func isPalindrome(str string) bool {
//   regStr := strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(str), -1),"")

//   if len(regStr) <= 1 {
//     return true
//   }

//   if len(regStr) == 2 {
//     return regStr[0] == regStr[1]
//   }

//   if regStr[0] == regStr[len(regStr) - 1] {
//     return isPalindrome(regStr[1:len(regStr) - 1])
//   }

//   return false
// }

func isPalindrome(str string) bool {
	regStr := strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(str), -1), "")

	s, e := 0, len(regStr)-1

	for s <= e {
		if regStr[s] != regStr[e] {
			return false
		}

		s++
		e--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("ae"))
	fmt.Println(isPalindrome(""))
}
