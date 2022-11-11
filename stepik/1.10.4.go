package main

import "fmt"

func main() {
	var (
		max   = 0
		score = -1
		n     = 1
	)
	for fmt.Scan(&n); n > 0; fmt.Scan(&n) {
		if n > max {
			max = n
			score = 0
		}
		if n == max {
			score++
		}
	}
	fmt.Println(score)
}
