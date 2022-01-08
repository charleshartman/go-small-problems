package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	letters := strings.Split(arr[0], "")
	return strings.Join(letters, "***")
}

func main() {
	s1 := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(s1) == "b***i***t***c***o***i***n")
	s2 := []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
	fmt.Println(TwoSort(s2) == "L***e***t***s")
}
