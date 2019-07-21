/**
* Regular Expression Matching:https://leetcode.com/problems/regular-expression-matching/
*/

package main

import "fmt"


func findNext(s string , i int , elements []string , elementsSize int , position int) bool {
	tempPosition := position + 1
	for tempPosition < elementsSize {
		temp := elements[tempPosition]
		l := len(temp)
		if temp[0] == '.' {
			return false
		}
		if l == 2 && temp[1] == '*' {
			if s[i] != temp[0] {
				// 寻找下一个
				tempPosition++
				continue
			}else if i + 1 < len(s) && findNext(s, i + 1, elements, elementsSize, position - 1){
				// 寻找下一个
				tempPosition++
				continue
			}
			
		}
		if s[i] == temp[0] {
			return true
		}
		return false
		
	}
	return false
}

func compile(p string) []string {
	var elements []string
	var i = 0
	var n = len(p)
	for i < n {
		if p[i] == '.' {
			if i + 1 < n && p[i + 1] == '*' {
				elements = append(elements , p[i : i+2])
				i += 2
			}else {
				elements = append(elements , p[i : i+1])
				i++
			}
		}else if i + 1 < n && p[i + 1] == '*' {
			j := i + 2
			for j < n && p[j] == p[i] {
				j++
			}     
			elements = append(elements , p[i : i + 2])
			i = j
		}else {
			// 纯字符情况，例如ababa
			var j = i
			for j < n && j + 1 < n && p[j + 1] != '*' && p[j + 1] != '.' {
				j++
			}
			if j != i {
				elements = append(elements , p[i : j])
				i = j
			}else {
				elements = append(elements , p[i : j + 1])
				i++
			}
		}

	}

	return elements
}

func match(elements []string , elementsSize int , position int , s string , i int) bool {
	for ; position < elementsSize ; position ++ {
		e := elements[position]
		l := len(e)
		if l == 2 && e[1] == '*' &&  s[i] != e[0] {
			continue
		}
		if e[0] == '.' && elements[position - 1][0] == '.' && position == elementsSize - 1 {
			return true
		}
		if e[0] == '.' && i + 1 < len(s) && position + 1 < elementsSize && 
			match(elements , elementsSize , position + 1 , s , i + 1) {
			return true
		}
		return s[i] == e[0]
	}
	return false
}

func isMatchImpl(s string, p string) bool {
	elements := compile(p)
	fmt.Println(elements)
	elementsSize := len(elements)
	n := len(s)
	var (
		i = 0
	)
	// for _,e := range elements {
	for position := range elements {
		e := elements[position]
		len := len(e)
		if i >= n {
			if len == 2 && e[1] == '*' {
				continue
			}
			return false
		}
		if len == 1 {
			if e == "." || e == string(s[i]) {
				i++
			}else {
				return false
			}
		}else if len == 2 && e[1] == '*'{
			if e[0] == '.' {
				for i < n {
					// 如果有下个匹配的规则需要中断
					if match(elements , elementsSize , position + 1 , s , i) {
						break
					}
					i++
					// if i < n && findNext(s , i, elements , elementsSize , position) {
					// 	break;
					// }
				}
			}else {
				var j = i
				// 找到最后一个匹配的数字
				for j < n && s[j] == e[0] {
					j ++
				}
				if j > i + 1 {
					// 然后拿后续的规则和最后一个数字做匹配，如果有匹配的规则，则需要后退
					if match(elements , elementsSize , position + 1 , s , j - 1) {
						i = j -1
					}else {
						i = j
					}
				}else {
					i = j
				}
				// for i < n && s[i] == e[0] {
				// 	i++
				// 	if i < n && findNext(s , i, elements , elementsSize , position) {
				// 		break;
				// 	}
				// }
				// fmt.Println(e , s , i , string(s[i]))
			}
		}else {
			if i + len <= n && s[i : i + len] == e {
				i = i + len
			}else {
				return false
			}
		}
	}
	if i < n {
		fmt.Println("me" , i , n)
		return false
	}
	return true	
}

func recursiveMatch(s string , p string) bool {
	sn := len(s)
	pn := len(p)
	if pn == 0 {
		return sn == 0
	}
	firstMatch := sn > 0 && (s[0:1] == p[0:1] || p[0:1] == ".")
	if pn >= 2 && p[1:2] == "*" {
		return recursiveMatch(s, p[2:]) || (firstMatch && recursiveMatch(s[1:], p))
	}else {
		return firstMatch && recursiveMatch(s[1:], p[1:])
	}
}

func isMatch(s string, p string) bool {
	return dpMatch(s, p)
}

func dpMatch(s string , p string) bool {
	sl := len(s)
	pl := len(p)
	memory := make([][]int , sl + 1)
	for i := 0 ; i < sl + 1 ; i++ {
		memory[i] = make([]int , pl + 1)
		for j := range memory[i] {
			memory[i][j] = -1
		}
	}
	return dp(memory , 0 , 0  , s , p)
}

func dp(memory [][]int , si int , pi int , s string , p string) bool {
	if memory[si][pi] != -1 {
		return memory[si][pi] == 1
	}
	var res bool = false
	if pi == len(p) {
		res = si == len(s)
	}else {
		firstMatch := si < len(s) && (s[si:si+1] == p[pi:pi+1] || p[pi:pi+1] == ".")
		if pi + 1 < len(p) && p[pi + 1 : pi + 2] == "*" {
			res = dp(memory , si , pi + 2 , s , p ) || (firstMatch && dp(memory , si + 1 , pi , s , p))
		}else {
			res = firstMatch && dp(memory , si + 1 , pi + 1, s , p)
		}
	}
	if res {
		memory[si][pi] = 1
	}else {
		memory[si][pi] = 0
	}
	
	return res
}

func main() {
	fmt.Println(isMatch("aaa" , "a.a"))
	fmt.Println(isMatch("mississippi" , "mis*is*ip*."))
	fmt.Println(isMatch("aaa" , "ab*a*c*a"))
	fmt.Println(isMatch("aaa" , "a*a"))
	fmt.Println(isMatch("aaa" , "a*"))
	fmt.Println(isMatch("aaa" , "aaa"))
	fmt.Println(isMatch("aaa" , "aaa"))
	fmt.Println(isMatch("aaca" , "ab*a*c*a"))
	fmt.Println(isMatch("a" , "ab*"))
	fmt.Println(isMatch("bbbba" , ".*a*a"))
	fmt.Println(isMatch("ab" , ".*.."))
	fmt.Println(isMatch("ab" , ".*..c*"))
}