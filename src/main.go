package main

import (
	"fmt"
	"lifecore"
)

func main() {
	fmt.Printf("\033[2J")
	for i := 0; i < 1000; i++ {
		fmt.Printf("\033[1;1H")
		fmt.Println(i)
		fmt.Println(i)
		fmt.Printf("%d", i)
	}
	fmt.Print("\n")

	//var arrr [][]int
	//arrr = make([][]int,3,3)
	//arrr[0][0] = 3
	//fmt.Println(arrr[1][1])

	lb := lifecore.LifeBoard{}
	lb.InitBoard(3, 4)
	lb.SetCell(2, 0, true)
	lb.SetCell(2, 1, true)
	lb.SetCell(1, 0, true)

	lb.NextGen()

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if lb.GetCell(i, j) {
				fmt.Print("O")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}

}
