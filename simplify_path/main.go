/* Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, convert it to the simplified canonical path.

In a Unix-style file system, a period '.' refers to the current directory, a double period '..' refers to the directory up a level, and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. For this problem, any other format of periods such as '...' are treated as file/directory names.

The canonical path should have the following format:
The path starts with a single slash '/'.
Any two directories are separated by a single slash '/'.
The path does not end with a trailing '/'.
The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
Return the simplified canonical path.


Example 1:
Input: path = "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.

Example 2:
Input: path = "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: path = "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.


Constraints:

1 <= path.length <= 3000
path consists of English letters, digits, period '.', slash '/' or '_'.
path is a valid absolute Unix path.

UNDERSTAND THE PROBLEM
  Inputs
    - string
  Outputs
    - string
Rules/Implicit Requirements
  Explicit
    - string can be anywhere from one to 3000 characters in length
    - characters can be English letter, digits, period, slash, or underscore
    - path is a valid absolute Unix path
    - The path starts with a single slash '/'.
    - Any two directories are separated by a single slash '/'.
    - The path does not end with a trailing '/'.
    - The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
    - Return the simplified canonical path
    - period refers to the current directory
    - double period refers to one directory up
  Implicit
    -
Clarifying Questions
  -
Mental Model
  - Using a stack structure
  - Split the given string by slashes, which should leave only words or periods
  - Go through every element of the split string
    - If the element is a single period, do not add to stack
    - If the element is a double period
      - Do not add to stack AND
      - Pop element from stack
    - If the element is a "word" we just add to stack

  - We join elements of stack with slash
  - Prepend slash to joined elements

EXAMPLES/TEST CASES
  Normal Cases
    Example 1:
      Input: path = "/home/"
      Output: "/home"
      Explanation: Note that there is no trailing slash after the last directory name.

    Example 2:
      Input: path = "/../"
      Output: "/"
      Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.

    Example 3:
      Input: path = "/home//foo/"
      Output: "/home/foo"
      Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.

  Edge Cases
    -

DATA STRUCTURE
  Inputs
    -
  Rules
    -

ALGORITHM
  - Using a stack structure
  - Split the given string by slashes, which should leave only words or periods
  - Go through every element of the split string
    - If the element is a single period, do not add to stack
    - If the element is a double period
      - Do not add to stack AND
      - Pop element from stack
    - If the element is a "word" we just add to stack

  - We join elements of stack with slash
    - iterate until the stack is empty
      - pop element off the stack, concat popped element to new string
  - Prepend slash to joined element

CODE
*/

package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(i string) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() string {
	p := len(s.items) - 1
	if p >= 0 {
		popped := s.items[p]
		s.items = s.items[:p]
		return popped
	}
	return ""
}

func (s *Stack) Read() string {
	return s.items[len(s.items)-1]
}

func simplifyPath(path string) (result string) {
	stack := Stack{}

	for _, v := range strings.Split(path, "/") {
		switch v {
		case "":
			continue
		case ".":
			continue
		case "..":
			stack.Pop()
		default:
			stack.Push(v)
		}
	}

	if len(stack.items) == 0 {
		return "/"
	}

	for len(stack.items) > 0 {
		result = "/" + stack.Pop() + result
	}

	return result
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home/.."))
}

/*
/home/../.. -> Go down to home and then up one level, canonical: /home
cd .. -> Go up one level, canonical: /
cd ../.. -> Go up two levels, canonical: /

// Always start with a slash
"/home/" -> /home
"/../" -> /
"/home//foo/" -> /home/foo
*/
