package main

import "fmt"

func minMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, val := range array {
        if max < val {
            max = val
        }
        if val < min {
            min = val
        }
    }
    return min, max
}

func minMaxOfSet(set map[int]bool) (int, int) {
	keys := make([]int, 0)
	for k := range set {
    	keys = append(keys, k)
	}
	return minMax(keys)
}

func count(el int, slice []int) int{
	c := 0
	for i := range slice {
		if el == slice[i] { c++ }
	}
	return c
}

func isBoring(slice []int) bool {
	myset := make(map[int]bool)
	for i := range slice {
		myset[slice[i]] = true
	}
	// for k, v := range myset {         
	// 	fmt.Println("k:",k, "v:", v)
	// }
	if len(myset) == 2 {
		mmin, mmax := minMaxOfSet(myset)
		if count(mmin,slice) == 1{
			return true
		}
		if count(mmax,slice) == 1{
			return true
		}
	}
	if count(1, slice) == len(slice) {
		return true
	}
	return false
}


func MaxBoringPrefix(n int, a []int) int {
	res := 1
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		k := a[i]
		m[k]++
		vals := make([]int, 0)
		for _, v := range m {
			vals = append(vals, v)
		}
		if isBoring(vals) { res = i + 1}
	}
	return res
}

func main() {
	var n, x int
	fmt.Scan(&n)
	a := make([]int,0, n)
	for i := 0; i < n ; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	fmt.Println(MaxBoringPrefix(n, a))
}