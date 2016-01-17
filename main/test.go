package main

import (
	"fmt"
	"is_golang/mysamples"
)

func main()  {
	fmt.Println("start!")
	sum := mysamples.Sum(19, 56)
	fmt.Printf("%d\n", sum)
}
