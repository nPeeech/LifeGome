package lifecore

type LifeBoard struct {
	board     [][]bool
	CopyBoard [][]bool
	row       int
	column    int
}

var (
	iikanjiX = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	iikanjiY = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
)

func (l *LifeBoard) InitBoard(row, column int) {
	l.board = make([][]bool, row)
	for i := 0; i < row; i++ {
		l.board[i] = make([]bool, column)
	}
	//fmt.Println("board:     ",&l.board[0][0])

	l.CopyBoard = make([][]bool,row)
	for i := 0; i < row; i++ {
		l.CopyBoard[i] = make([]bool,column)
	}
	//fmt.Println("NextBoard: ",&l.NextBoard[0][0])

	l.Sync()

	//fmt.Println("board:     ",&l.board[0][0])
	//fmt.Println("NextBoard: ",&l.NextBoard[0][0])
	l.row = row
	l.column = column
}

func (l *LifeBoard) GetCell(y, x int) bool {
	if x < 0 || y < 0 || y >= l.row || x >= l.column {
		return false
	} else {
		return l.board[y][x]
	}
}
func (l *LifeBoard) SetCell(y int, x int, cel bool) {
	l.board[y][x] = cel
}
func (l *LifeBoard) setNextCell(y int, x int, cel bool) {
	l.CopyBoard[y][x] = cel
}

func (l *LifeBoard) NextGen() {
	for i := 0; i < l.row; i++ {
		for j := 0; j < l.column; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				if l.GetCell(i+iikanjiY[k], j+iikanjiX[k]) {
					//fmt.Printf("Trueee i= %d, j= %d \n",i+iikanjiY[k],j+iikanjiX[k])
					cnt++
				}
			}
			if !l.GetCell(i, j) && cnt == 3 {
				l.setNextCell(i, j, true)
				//fmt.Printf("i= %d, j= %d, cnt= %d,誕生\n",i,j,cnt)
			} else if l.GetCell(i, j) && (cnt == 2 || cnt == 3) {
				l.setNextCell(i, j, true)
				//fmt.Printf("i= %d, j= %d, cnt= %d,生存\n",i,j,cnt)
			} else if l.GetCell(i, j) && cnt <= 1 {
				l.setNextCell(i, j, false)
				//fmt.Printf("i= %d, j= %d, cnt= %d,過疎\n",i,j,cnt)
			} else if l.GetCell(i, j) && cnt >= 4 {
				l.setNextCell(i, j, false)
				//fmt.Printf("i= %d, j= %d, cnt= %d,過密\n",i,j,cnt)
			} else {

				//fmt.Printf("i= %d, j= %d, cnt= %d,else\n",i,j,cnt)
			}
			//fmt.Println(l.board)
			//fmt.Print("\n")
		}
	}
	for i := 0; i < l.row; i++{
		copy(l.board[i],l.CopyBoard[i])
	}
}

func (l *LifeBoard) Sync(){
	for i := 0; i < l.row; i++{
		copy(l.CopyBoard[i],l.board[i])
	}
}
