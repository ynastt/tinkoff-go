// ok
package main

import "fmt"

func fillMap(m map[int]int, slice string) map[int]int {
	nn := len(slice)
	for i := 0; i < nn; i++ {
		if(slice[i] == 'a') { m[0]++}
		if(slice[i] == 'b') { m[1]++}
		if(slice[i] == 'c') { m[2]++}
		if(slice[i] == 'd') { m[3]++}
	}
	return m
}

func mapIsGood(m map[int]int) bool {
	if m[0] >= 1 && m[1] >= 1 &&
		m[2] >= 1 && m[3] >= 1 {
			return true
		}
	return false
}

func lenOfBestSubstring(s string, m map[int]int) int {
	nn := len(s)
	mini := 999999 
	c := 0
	for i := 0; i < nn - 1; i++ {
		c = 0
		for j := i + 1; j < nn; j++ {
			c++
			arr := s[i:j+1]
			l := len(arr)
			m := fillMap(m, arr)
			// fmt.Println("c:", c)
			// fmt.Println("i:", i, "j:", j)
			if ( mapIsGood(m)) {
				// fmt.Println("map is good", "c:", c, l)
				if l < mini {
					mini = l
				}
				c = 0
				m[0], m[1], m[2], m[3] = 0, 0, 0, 0
				break
			}
		}
	}
	if mini != 999999{
		return mini
	}
	return -1
}

func main3() {
	var n int
	var s string
	//0 - a  1 - b  2 - c  3 -d
	var m = map[int]int {
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	fmt.Scan(&n)
	fmt.Scan(&s)
	if len(s) != n {
		panic(1)
	}
	fmt.Println()
	fmt.Println(lenOfBestSubstring(s, m))
}