// ok
package main

import "fmt"

func main2() {
	var n, m, k uint64
	fmt.Scan(&n, &m, &k)
	res := (n * k + m - 1) / m
	fmt.Println(res)
}