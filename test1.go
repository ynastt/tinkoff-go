// ok
package main

import "fmt"

func main() {
	var a [4]int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i])
	}

	f := false // is ubiv
	ok := true
	if a[0] > a[1] {
		f = true
	}
	for i := 1; i < 3; i++ {
		if f == true && a[i] < a[i+1] {
			fmt.Println("NO")
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("YES")
	}
}
