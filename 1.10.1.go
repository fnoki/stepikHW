package main

import "fmt"

func main() {
	i := 1
	for n := 1; i < 11; i++ {
		n *= n
		fmt.Println(n)
		n = i + 1
	}
}
