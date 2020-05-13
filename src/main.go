package main

import (
	"lifecore"
	"render"
	"time"
	"math/rand"
)

func main() {
	row := 10
	column := 10
	lb := lifecore.LifeBoard{}
	lb.InitBoard(row, column)
	lb.SetCell(5, 5, true)
	lb.SetCell(5, 6, true)
	lb.SetCell(5, 7, true)
	lb.SetCell(4,5,true)
	lb.SetCell(3,6,true)


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

	for i:=0; i<80; i++{
		render.Render(lb.CopyBoard,row,column)
		lb.NextGen()
		time.Sleep(100 * time.Millisecond)
	}



}
