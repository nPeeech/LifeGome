package main

import (
	"fmt"
	"lifecore"
	"render"

	"math/rand"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	row := 100
	column := 100
	lb := lifecore.LifeBoard{}
	lb.InitBoard(row, column)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))

	for i:=0; i<row; i++{
		for j:=0; j<column; j++{
			rnd := rand.Intn(2)
			if rnd==1 {
				lb.SetCell(i,j,true)
			} else {
				lb.SetCell(i,j,false)
			}

		}
	}

	render.InitTerminal()
	lb.Sync()

	for i:=0; i<200; i++{
		render.Draw2LifeBoardInHtml(lb.CopyBoard,row,column)
		lb.NextGen()
		time.Sleep(200 * time.Millisecond)
	}



}
