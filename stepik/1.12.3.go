package main

import "fmt"

func main() {
	var (
		x string
	)
	fmt.Print("Ты хочешь кушать?")
	fmt.Scan(&x)
	if x == "да" {
		for i := 0; i < 1000; i++ {

			fmt.Println("Я хочу кушать")
		}
	} else if x == "нет" {
		for i := 0; i < 1000; i++ {
			fmt.Println("Я не хочу кушать")
		}
	}

}
