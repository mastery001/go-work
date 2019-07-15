/**
* Substring with Concatenation of All Words - https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/

package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	var (
		n = len(s)
		m = len(words)
	)
	var res []int
	if m == 0 {
		return res
	}
	wordLen := len(words[0])
	if m * wordLen > n {
		return res
	}
	wordCount := make(map[string]int)
	for _,word := range words {
		_ , ok := wordCount[word]
		if ok {
			wordCount[word] += 1
		}else {
			wordCount[word] = 1
		}
	}
	for i := 0 ; i < n - m * wordLen + 1 ; i ++ {
		sCount := make(map[string]int)
		num := 0
		for num < m {
			temp := s[i + num * wordLen : i + (num + 1) * wordLen]
			if _ , ok := wordCount[temp]; ok {
				if _ , ok := sCount[temp]; ok {
					sCount[temp] += 1
				}else {
					sCount[temp] = 1
				}
				if wordCount[temp] < sCount[temp] {
					break
				}
			}else {
				break;
			}
			num ++
		}
		if num == m {
			res = append(res , i)
		}
	}
	return res
}

func main() {
	// s := "barfoothefoobarman"
 //  	words := []string{"foo","bar"}

  	s := "wordgoodgoodgoodbestword"
  	words := []string{"word","good","best","word"}
  	fmt.Println(findSubstring(s, words))
}