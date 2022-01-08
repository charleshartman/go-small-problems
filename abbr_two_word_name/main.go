package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	// return strings.ToUpper(strings.Split(names[0], "")[0]) + "." + strings.ToUpper(strings.Split(names[1], "")[0])
	return strings.ToUpper(string(names[0][0])) + "." + strings.ToUpper(string(names[1][0]))
}

func EvenOrOdd(number int) string {
	parity := map[int]string{0: "Even", 1: "Odd"}
	return parity[number%2]
}

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
}
