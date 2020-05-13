package render

import (
	"fmt"
	"syscall/js"
)
func InitTerminal(){
	fmt.Print("\033[2J")

}
func Draw2Terminal(board [][]bool, row int, column int){
	fmt.Print("\033[1;1H")
	for i:=0; i<row; i++{
		for j:=0; j<column; j++{
			if board[i][j]{
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")

	}
}

func Draw2LifeBoardInHtml(board [][]bool,row int, column int){
	document := js.Global().Get("document")
	htmlBoard := document.Call("getElementById","LifeBoard")
	text := board2Text(board,row,column)
	htmlBoard.Set("innerHTML",text)
}

func board2Text(board [][]bool,row int,column int) string{
	var text string
	for i:=0; i<row; i++{
		for j:=0; j<column; j++{
			if(board[i][j]){
				text+="■"
			} else {
				text+="□"
			}
		}
		text+="<br>"
	}
	return text
}