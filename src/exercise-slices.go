package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var slices [][]uint8
	for i := 0 ; i < dy ; i++ {
		slicesX := make([]uint8 , 0 , dx)
		for j := 0 ; j < dx ; j++ {
			slicesX = append(slicesX , uint8(i * j))
		}
		slices = append(slices , slicesX)
	}	
	return slices
}

func main() {
	pic.Show(Pic)
}
