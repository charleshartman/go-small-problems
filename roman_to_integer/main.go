/* 13. Roman to Integer
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together.
12 is written as XII, which is simply X + II. The number 27 is written as XXVII,
which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written
as IV. Because the one is before the five we subtract it making four. The same
principle applies to the number nine, which is written as IX. There are six
instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999]. MMMCMXCIX

Add these cases:
IV   4
IX   9
XL   40
XC   90
CD   400
CM   900

UNDERSTAND THE PROBLEM
  Inputs
    - string
  Outputs
    - integer
Rules/Implicit Requirements
  Explicit
    - transform Roman numeral represented as string to the correct integer
    - all arguments will be valid Roman numerals
    - argument will be greater than 1 and less than 15 chars
  Implicit
    - all the normal rules regarding how Roman numerals are constructed (this
	  includes special preceding "subtraction" cases
    - this means we need to do extra analysis on I, X, and C and the character after

Clarifying Questions
  -
Mental Model (optional)
  -

EXAMPLES/TEST CASES
  Base Cases

  Input: s = "III"
  Output: 3
  Explanation: III = 3.

  Input: s = "LVIII"
  Output: 58
  Explanation: L = 50, V= 5, III = 3.

  Input: s = "MCMXCIV"
  Output: 1994
  Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Edge Cases
  Input: s = "MMMCMXCIX"
  Output: 3999
  Explanation: M * 3 = 3000
               CM = 900
               XC = 90
               IX = 9

  Input: s = "I"
  Output: 1

DATA STRUCTURE
  - create a map of numerals
  - input - string
    - convert to slice

ALGORITHM
  - define the map of numerals as a constant
  - declare RESULT variable and assign to 0
  - split the input string to form an array of chars (slice)
  - define a variable to use as a pointer to the end char
    - start at index pos 1
  - define a variable to use as a pointer to the starting char
    - start at index pos 0
  - if one character length, look up value, add to RESULT
  - iterate over the array of chars
    - if the pointers are out of bounds break
    - generate a slice from starting index to ending index + 1
    - Join the contents of the slice, which generates a string
    - Check if the string exists within the map of Romans
      - If it does:
        - get the value from Map and add to RESULT
        - move starting and end pointer + 2
      - If it does not:
        - get the value of the string at the current starting position pointer from Map and add to RESULT
        - move starting and end pointer + 1
  - return RESULT
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanNumeralToInt("LVIII"))     // 58
	fmt.Println(romanNumeralToInt("III"))       // 3
	fmt.Println(romanNumeralToInt("MMMCMXCIX")) // 3999
	fmt.Println(romanNumeralToInt("IV"))        // 4
	fmt.Println(romanNumeralToInt("I"))         // 1
}

func romanNumeralToInt(r string) int {
	romanNumerals := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4,
		"IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900,
	}

	result := 0

	chars := strings.Split(r, "")

	for i := 0; i < len(r); {
		if i == len(r)-1 {
			result += romanNumerals[chars[i]]
			break
		}

		currentValue := chars[i] + chars[i+1]
		if romanNumerals[currentValue] == 0 {
			currentValue = string(chars[i])
			i += 1
		} else {
			i += 2
		}

		result += romanNumerals[currentValue]
	}

	return result
}

// "MCM"
// An error will be raised when the start pointer (i) goes out of bounds
// - This gets taken care of by the loop conditional

// "MMM"
// An error will be raised when the end pointer (i + 2) goes out of bounds
// - Guard clause: If the current index is equal to the length of slice - 1,
//   that means there is only one character left, so we only need to add the
// 	 value associated with it to the RESULT and break out of the loop
