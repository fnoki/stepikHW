package main

import "fmt"

func main() {

	var i uint8
	var workArray [10]uint8
	for m := 0; m < len(workArray); m++ {
		fmt.Scan(&i)
		workArray[m] = i
	}
	
	fmt.Println(workArray)
}
