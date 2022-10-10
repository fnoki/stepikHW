package main

import "fmt"

func main() {
	var n, b, sum int
	fmt.Scan(&n)
	for m := 0; m < n; m++ {
		fmt.Scan(&b)
		if b%8 == 0 && b < 100 && b > 9 {
			sum += b
		}
	}
	fmt.Println(sum)
}
