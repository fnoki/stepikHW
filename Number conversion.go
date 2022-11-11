package main

import "fmt"

func main() {
	var (
		x string
		h bool
	)
	fmt.Scan(&x)
	for len(x) > 4 {
		if len(x) > 4 {
			fmt.Println("Введите правильное число")
			fmt.Scan(&x)
		}
	}
	j := len(x)
	for i := 0; i < j; i++ {
		q := string(x[i])
		if j == 4 && i == 0 {
			switch q {
			case "1":
				q = "одна тысяча"
			case "2":
				q = "две тысячи"
			case "3":
				q = "три тысячи"
			case "4":
				q = "четыре тысячи"
			case "5":
				q = "пять тысяч"
			case "6":
				q = "шесть тысяч"
			case "7":
				q = "семь тысяч"
			case "8":
				q = "восемь тысяч"
			case "9":
				q = "девять тысяч"
			case "0":
				q = "-"
			}
		}
		if j == 4 && i == 1 || j == 3 && i == 0 {
			switch q {
			case "1":
				q = "сто"
			case "2":
				q = "двести"
			case "3":
				q = "триста"
			case "4":
				q = "четырста"
			case "5":
				q = "пятьсот"
			case "6":
				q = "шестьсот"
			case "7":
				q = "семьсот"
			case "8":
				q = "восемьсот"
			case "9":
				q = "девятьсот"
			case "0":
				q = ""
			}
		}
		if j == 4 && i == 2 || j == 3 && i == 1 || j == 2 && i == 0 {
			switch q {
			case "1":
				q = ""
			case "2":
				q = "двадцать"
			case "3":
				q = "тридцать"
			case "4":
				q = "сорок"
			case "5":
				q = "пятьдесят"
			case "6":
				q = "шестьдесят"
			case "7":
				q = "семьдесят"
			case "8":
				q = "восемьдесят"
			case "9":
				q = "девяносто"
			case "0":
				q = ""
			}
		} else if q == "1" {
			h = false
		}
		if j == 4 && i == 3 || j == 3 && i == 2 || j == 2 && i == 1 || j == 1 && i == 0 {
			if h == true {
				switch q {
				case "1":
					q = "один"
				case "2":
					q = "два"
				case "3":
					q = "три"
				case "4":
					q = "четыре"
				case "5":
					q = "пять"
				case "6":
					q = "шесть"
				case "7":
					q = "семь"
				case "8":
					q = "восемь"
				case "9":
					q = "девять"
				case "0":
					q = ""
				}
			} else {
				switch q {
				case "1":
					q = "одинадцать"
				case "2":
					q = "двенадцать"
				case "3":
					q = "тринадцать"
				case "4":
					q = "четырнадцать"
				case "5":
					q = "пятьнадцать"
				case "6":
					q = "шестьнадцать"
				case "7":
					q = "семьнадцать"
				case "8":
					q = "восемьнадцать"
				case "9":
					q = "девятьнадцать"
				case "0":
					q = "десять"
				}
			}
		}
		fmt.Print(q, " ")
	}
	if j > 1 {
		fmt.Print("долларов")

	} else {
		fmt.Print("доллара")
	}
}
