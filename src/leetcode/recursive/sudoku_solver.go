/**
* Sudoku Solver - https://leetcode.com/problems/sudoku-solver/
*/

package main

import (
	"fmt"
)

const N = 9

// var rowBit [N][N + 1]byte
// var columnBit [N][N + 1]byte
// var squareBit [N][N + 1]byte

var (
	rowBit = make([][]byte , 9)
	columnBit = make([][]byte , 9)
	squareBit = make([][]byte , 9)
	solved = false
)



func init() {
	for i := 0 ; i < N ; i++ {
		rowBit[i] = make([]byte , N )
		columnBit[i] = make([]byte , N)
		squareBit[i] = make([]byte , N)
	}
}

func couldFill(value byte , row int , column int) bool {
	return rowBit[row][value - 1] == 0 && columnBit[column][value - 1] == 0 && squareBit[3 * (row / 3) + column / 3][value - 1] == 0
}

func placeCurrentValue(board [][]byte, value byte , row int , column int) {
	rowBit[row][value - 1]++
	columnBit[column][value - 1]++
	squareBit[3 * (row / 3) + column / 3][value - 1]++
	board[row][column] = value
	fmt.Println(row , column , board[row][column] , value)
}

func removeCurrentValue(board [][]byte, value byte , row int , column int) {
	rowBit[row][value - 1]--
	columnBit[column][value - 1]--
	squareBit[3 * (row / 3) + column / 3][value - 1]--
	board[row][column] = '.'
}

func placeNext(board [][]byte,row int , column int) {
	// end
	if row == N - 1 && column == N - 1 {
		solved = true
	}else if row == N - 1 {
		backtracking(board , row , column + 1)
	}else {
		backtracking(board , row + 1 , column)
	}
}


func backtracking(board [][]byte , row int , column int) {
	// fill if this gird value is '.' 
	if board[row][column] == '.' {
		var i byte = 1
		for ; i < 10 ; i ++ {
			// if value can fill the board , place and place next
			if couldFill(i , row , column) {
				placeCurrentValue(board , i , row , column)
				placeNext(board , row , column)
				// if sudoku is solved, there is no need to backtrack
          		// since the single unique solution is promised
				if !solved {
					removeCurrentValue(board , i, row, column)
				}
			}
		}
	}else {
		placeNext(board , row , column)
	}
}


func solveSudoku(board [][]byte)  {
	// backtracking(board , 0, 0)

}

func main() {

	board := [][]byte{
		{5,3,'.','.',7,'.','.','.','.'},
		{6,'.','.',1,9,5,'.','.','.'},
		{'.',9,8,'.','.','.','.',6,'.'},
		{8,'.','.','.',6,'.','.','.',3},
		{5,'.','.',8,'.',3,'.','.',1},
		{7,'.','.','.',2,'.','.','.',6},
		{'.',6,'.','.','.','.',2,8,'.'},
		{'.','.','.',4,1,9,'.','.',5},
		{'.','.','.','.',8,'.','.',7,9}}

	solveSudoku(board)

	fmt.Println(board)
}