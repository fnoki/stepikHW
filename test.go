package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	switch{
	case a < 10: fmt.Println(a)
	case a < 100: fmt.Println(a/10)
	case a < 1000: fmt.Println(a/100)
	case a < 10000: fmt.Println(a/1000)
	case a < 100000: fmt.Println(a/10000)
	}
}