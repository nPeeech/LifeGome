package render

import (
	"fmt"
)
func InitScreen(){
	fmt.Print("\033[2J")

}
func Render(board [][]bool, row int, column int){
	fmt.Print("\033[1;1H")
	for i:=0; i<row; i++{
		for j:=0; j<column; j++{
			if board[i][j]{
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Print("\n")

	}
}
