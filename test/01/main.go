package main

import (
	"fmt"
	"time"
)

type NameCode struct {
	Name string
	Code string
}

func main() {
	var t time.Time
	fmt.Printf("start!\n")
	fmt.Printf("time : %v\n", t)

	arr := make([]NameCode, 0)
	fmt.Printf("arr : %p\n", arr)
	fmt.Printf("arr : %p\n", &arr)

	arr1 := make([]NameCode, 1)
	arr1[0] = NameCode{
		Name: "aa",
		Code: "code",
	}
	p := &arr1
	fmt.Printf("arr1 p : %v\n", p)
	fmt.Printf("arr1 *p : %v\n", *p)
	fmt.Printf("arr1 : %p\n", &arr1)
}
