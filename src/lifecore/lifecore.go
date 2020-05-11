package lifecore

type LifeBoard struct {
	board     [][]bool
	nextBoard [][]bool
	row       int
	column    int
}

var (
	iikanjiX = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	iikanjiY = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
)

func (l *LifeBoard) InitBoard(row, culumn int) {
	l.board = make([][]bool, row)
	for i := 0; i < row; i++ {
		l.board[i] = make([]bool, culumn)
	}
	l.nextBoard = l.board
	l.row = row
	l.column = culumn
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
	l.nextBoard[y][x] = cel
}
func (l *LifeBoard) NextGen() {
	for i := 0; i < l.row; i++ {
		for j := 0; j < l.column; j++ {
			cnt := 0
			for k := 0; k < 8; k++ {
				if l.GetCell(i+iikanjiY[k], j+iikanjiX[k]) {
					cnt++
				}
			}
			if !l.GetCell(i, j) && cnt == 3 {
				l.setNextCell(i, j, true)
			} else if l.GetCell(i, j) && (cnt == 2 || cnt == 3) {
				l.setNextCell(i, j, true)
			} else if l.GetCell(i, j) && cnt >= 1 {
				l.setNextCell(i, j, false)
			} else if l.GetCell(i, j) && cnt >= 4 {
				l.setNextCell(i, j, false)
			}
		}
	}
	l.board = l.nextBoard
}
