package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\033[2J")
	for i := 0; i < 100000; i++ {
		fmt.Printf("\033[1;1H")
		fmt.Println(i)
		fmt.Println(i)
		fmt.Println(i)
	}
}
