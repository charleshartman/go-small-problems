package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	seen := map[int]bool{}

	for _, v := range nums1 {
		seen[v] = true
	}

	for _, v := range nums2 {
		if seen[v] {
			result = append(result, v)
			seen[v] = false
		}
	}

	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))                    // [2]
	fmt.Println(intersection([]int{9, 4, 9, 8, 4}, []int{4, 9, 5}))              // [9, 4]
	fmt.Println(intersection([]int{1, 2, 1, 32, 199}, []int{2, 2, 199, 23, 32})) // [2, 199, 32]
}

/* algorithm
- define a new array RESULT
- define a map SEEN
- iterate through the first array,
    - if the element is not in map
        - insert it (int key, bool val)
- iterate through second array
    - if element is present in SEEN
        - append it to RESULT
		- change map value to false
- return RESULT
*/
