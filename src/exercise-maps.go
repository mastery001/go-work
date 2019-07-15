package main

import (
	"golang.org/x/tour/wc"
	"strings"
	// "fmt"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	fields := strings.Fields(s)
	for _, value := range fields {
		v := m[value]
		if v != 0 {
			m[value] = m[value] + 1
		}else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}