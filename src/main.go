package main

import (
	"lifecore"
	"render"
	"time"
)

func main() {
	row := 40
	column := 40
	lb := lifecore.LifeBoard{}
	lb.InitBoard(row, column)
	lb.SetCell(5, 20, true)
	lb.SetCell(5, 21, true)
	lb.SetCell(5, 22, true)
	lb.SetCell(4,20,true)
	lb.SetCell(3,21,true)

	render.InitScreen()
	lb.Sync()

	for i:=0; i<100; i++{
		render.Render(lb.CopyBoard,row,column)
		lb.NextGen()
		time.Sleep(120 * time.Millisecond)
	}



}
