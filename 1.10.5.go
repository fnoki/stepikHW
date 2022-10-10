package main

import "fmt"

func main() {
	var (
		d    = 0
		c    = 0
		n    = 0
		exit = 0
	)
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for m := 1; m <= n; m++ {
		if m%c == 0 && m%d != 0 {
			exit = m
			fmt.Println(exit)
			break
		} else {
			continue
		}
	}
}
