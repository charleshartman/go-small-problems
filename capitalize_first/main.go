package main

import (
	"fmt"
	"strings"
)

func CapFirst(str string) string {
	letters := strings.Split(str, "")
	letters[0] = strings.ToUpper(letters[0])
	return strings.Join(letters, "")
}

func UpperEach(str string) string {
	result := []string{}
	words := strings.Split(str, " ")
	for _, w := range words {
		w = CapFirst(w)
		result = append(result, w)
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(CapFirst("aren't"))
	fmt.Println(UpperEach("How can mirrors be real if our eyes aren't real") ==
		"How Can Mirrors Be Real If Our Eyes Aren't Real")
}
