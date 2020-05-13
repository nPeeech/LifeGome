package main

import (
	"lifecore"
	"render"
	"time"
	"math/rand"
)

func main() {
	row := 80
	column := 180
	lb := lifecore.LifeBoard{}
	lb.InitBoard(row, column)
	lb.SetCell(5, 50, true)
	lb.SetCell(5, 51, true)
	lb.SetCell(5, 52, true)
	lb.SetCell(4,50,true)
	lb.SetCell(3,51,true)


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

	render.InitScreen()
	lb.Sync()

	for i:=0; i<1000; i++{
		render.Render(lb.CopyBoard,row,column)
		lb.NextGen()
		time.Sleep(100 * time.Millisecond)
	}



}
