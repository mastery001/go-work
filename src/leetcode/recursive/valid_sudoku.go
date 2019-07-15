/**
* Valid Sudoku - https://leetcode.com/problems/valid-sudoku/
*/
package main

import (
	"fmt"
)

const N = 9

func isValidSudokuSquare(board [][]byte , x int , y int) bool {
    bitMap := make([]byte , N)
    for i := x ; i < x + 3 ; i++ {
    	for j := y ; j < y + 3 ; j++ {
    		e := board[i][j]
    		if e == '.' {
    			continue
    		}
    		index := e - 1
    		if bitMap[index] == 0 {
    			bitMap[index] = 1
    		}else {
    			// fmt.Println("it's me" , i , j , index)
    			return false
    		}
    	}
    }	
    return true
}

func isValidSudokuOneLoop(board [][]byte) bool {
	var (
		row = make([][]byte , 9)
		column = make([][]byte , 9)
		square = make([][]byte , 9)
	)
	for i := 0 ; i < N ; i++ {
		row[i] = make([]byte , 9)
    	column[i] = make([]byte , 9)
    	square[i] = make([]byte , 9)	
	}
	// valid row
    for i := 0 ; i < N ; i++ {
    	for j := 0 ; j < N ; j++ {
    		if board[i][j] != '.' {
    			index := board[i][j] - 1
    			row[i][index]++
    			column[j][index]++
    			square[3 * (i / 3) + j / 3][index]++
    			if row[i][index] == 2 || column[j][index] == 2 || square[3 * (i / 3) + j / 3][index] == 2 {
    				return false
    			}
    		}
    	}
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
	// valid row
    for i := 0 ; i < N ; i++ {
    	bitMap := make([]byte , N)
    	for j := 0 ; j < N ; j++ {
    		e := board[i][j]
    		if e == '.' {
    			continue
    		}
    		index := e - 1
    		if bitMap[index] == 0 {
    			bitMap[index] = 1
    		}else {
    			// fmt.Println("it's me row" , i , j , index)
    			return false
    		}
    	}
    }
    // valid column
    for i := 0 ; i < N ; i++ {
    	bitMap := make([]byte , N)
    	for j := 0 ; j < N ; j++ {
    		e := board[j][i]
    		if e == '.' {
    			continue
    		}
    		index := e - 1
    		if bitMap[index] == 0 {
    			bitMap[index] = 1
    		}else {
    			// fmt.Println("it's me c" , i , j , index)
    			return false
    		}
    	}
    }
    // valid 3 x 3
    x , y := 0 , 0
    for x < N {
    	for y < N {
    		if isValidSudokuSquare(board, x, y) {
    			y += 3
    		}else {
    			return false
    		}
    	}
    	x += 3
    }
    return true
}

func main() {
	board := [][]byte{
		{8,3,'.','.',7,'.','.','.','.'},
		{6,'.','.',1,9,5,'.','.','.'},
		{'.',9,8,'.','.','.','.',6,'.'},
		{8,'.','.','.',6,'.','.','.',3},
		{5,'.','.',8,'.',3,'.','.',1},
		{7,'.','.','.',2,'.','.','.',6},
		{'.',6,'.','.','.','.',2,8,'.'},
		{'.','.','.',4,1,9,'.','.',5},
		{'.','.','.','.',8,'.','.',7,9}}

	fmt.Println(isValidSudokuOneLoop(board))
}