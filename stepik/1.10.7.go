package main

import "fmt"

func main() {
	var (
		x, p, y, year int
	)
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for year = 0; x <= y; year++ {
		if x <= y {
			x = x + x*p/100
		}
	}
	fmt.Println(year)
}
