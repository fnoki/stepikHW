package main

import "fmt"

func main() {

	var workArray [10]uint8
	var j uint8
	var (
		o uint8
		l uint8
	)
	var k=0
	for i := range workArray {
		fmt.Scan(&j)
		workArray[i] = j
	}
	for i := 0; i < 5; i++ {
		fmt.Scan(&k, &l)
		o = workArray[k]
		workArray[k] = workArray[l]
		workArray[l] = o
	}

	for _, x := range workArray {
		fmt.Print(x, " ")
	}
	fmt.Printf("type ok")
	return
}
