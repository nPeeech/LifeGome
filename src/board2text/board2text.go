package board2text

func Convert(board [][]bool,row int,column int){
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
}
