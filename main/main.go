package main

import (
	"fmt"
	"strings"
)

var message = "Bonjour"

func main() {
	greeting, name := "Hello", "ishi"
	fmt.Printf("%s %s!!\n", greeting, strings.ToUpper(name))
	fmt.Printf("%s %s!!\n", message, strings.ToUpper(name))

	var n int

	fmt.Printf("number : %d\n", n)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n--------------------------")
	printOdd()
}

func printOdd() {
	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n%2 != 0 {
			fmt.Println(n)
		} else {
			continue
		}
	}
}