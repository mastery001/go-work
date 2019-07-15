/**
* Z字形序列 - https://leetcode.com/problems/zigzag-conversion/
*/
package main

import (
	// "fmt"
	"strings"
	"bytes"
)

func convert(s string, numRows int) string {
	n := len(s)
    if numRows > n {
    	return s
    }
    var res = make([]string , n)
    // 字符的下标
    var i = 0
    for i < n {
    	// 填充竖列
    	for j := 0 ; j < numRows ; j++ {
    		if i < n {
    			res[j] += string(s[i])
    			i++
    		}
    	}
    	// 填充斜线
    	// 斜线的长度
    	for j := numRows - 2 ; j > 0 ; j-- {
     		if i < n {
    			res[j] += string(s[i])
    			i++
    		}   		
    	}
    }
    return strings.Join(res, "")
}

func convertByBuffer(s string, numRows int) string {
    n := len(s)
    if numRows > n {
    	return s
    }
    var res = make([]bytes.Buffer , n)
    // 字符的下标
    var i = 0
    for i < n {
    	// 填充竖列
    	for j := 0 ; j < numRows ; j++ {
    		if i < n {
    			res[j].WriteByte(s[i])
    			i++
    		}
    	}
    	// 填充斜线
    	// 斜线的长度
    	for j := numRows - 2 ; j > 0 ; j-- {
     		if i < n {
     			res[j].WriteByte(s[i])
    			i++
    		}   		
    	}
    }
    var resBuilder bytes.Buffer
    for i := range res {
    	resBuilder.WriteString(res[i].String())
    }
    return resBuilder.String()
}

func main() {
	var s = "PAYPALISHIRING"
	for i := 0 ; i < 100000 ; i++ {
		s += "PAYPALISHIRING"
	}
	convert(s, 4)
	// res := convert(s, 4)
	// fmt.Println(res)
	// fmt.Println(res == "PINALSIGYAHRPI")
}