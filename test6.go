package main

import (
	"fmt"
	"sort"
)

func sum(slice []int) int {
	s := 0
	for _, v := range slice {
		s += v
	}
	return s
}

func main6() {
	var n, s, x, y int
	fmt.Scan(&n, &s)
	l := make([]int, 0, n)
	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		l = append(l, x)
		r = append(r, y)
	}
	sort.Ints(l)
	sort.Ints(r)
	for sum(l) > s {
		for i := range l {
			l[i]--
		}
	}
	for sum(l) < s {
		for i := range l {
			l[i]++
		}
	}
	lmid := l[n / 2]
	for sum(r) > s {
		for i := range r {
			r[i]--
		}
	}
	for sum(r) < s {
		for i := range r {
			r[i]++
		}
	}
	rmid := r[n / 2]
	res := lmid
	if rmid > lmid {
        res = rmid
    } 
    fmt.Println(res)
}