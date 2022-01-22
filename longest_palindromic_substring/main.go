/* 5. Longest Palindromic Substring

PEDAC

input: string
output: string

rules:
- input string may include letters and digits
- return the longest possible palindromic substring from input string, in the event of 'ties' just return one of the 'winners'

clarifying questions:
- uppercase may be present in input string?
- are palindromes case sensitive (assume not, toLower before conditions)

mental model
- brute force [O(N^2)]
    - nested loop, find all substrings of given string and push into array
    - filter array to substrings that are palindromes
    - find the substring(s) within filtered array with longest length

- imagine better time complexity [O(N)]

*/

package main

import (
	"fmt"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func longestPalindrome(s string) (greatestSub string) {
	for i := 0; i < len(s)+1; i++ {
		for j := i; j < len(s)+1; j++ {
			sub := s[i:j]
			if sub == reverse(sub) && len(sub) > len(greatestSub) {
				greatestSub = sub
			}
		}
	}

	return greatestSub
}

func main() {
	fmt.Println(longestPalindrome("babad")) // "bab"
	fmt.Println(longestPalindrome("kayak")) // "kayak"
	fmt.Println(longestPalindrome("bill"))  // "ll"
	fmt.Println(longestPalindrome("ab"))    // "a"
	fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
