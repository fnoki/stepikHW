package main

import "fmt"

func main() {
	var (
		a int
		b int

		a1 int
		a2 int
		a3 int

		b1 int
		b2 int
		b3 int

		g int
		h int
	)
	fmt.Scan(&a)

	b = a / 1000
	a %= 1000

	a1 = a % 10
	a2 = a / 100
	a3 = (a / 10) % 10

	b1 = b % 10
	b2 = b / 100
	b3 = (b / 10) % 10

	g = a1 + a2 + a3
	h = b1 + b2 + b3

	if g == h {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
