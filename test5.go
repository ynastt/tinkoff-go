// ok
package main

import "fmt"

func sum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}

func isNorm(slice []int) bool {
	n := len(slice)
	for l := 0; l < n - 1; l++ {
		for r := l + 1; r < n; r++ {
			arr := slice[l:r+1]
			if sum(arr) == 0 {
				return true
			}	
		}
	}
	return false
}

func calcNorm(n int, a []int) int {
	col := 0
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			arr := a[i:j+1]
			if isNorm(arr) {
				col++
			}	
		}
	}
	return col
}

func main() {
	var n, x int
	fmt.Scan(&n)
	a := make([]int,0, n)
	for i := 0; i < n ; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	// fmt.Println("a", a)
	fmt.Println(calcNorm(n, a))
}
