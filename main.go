package main

import "fmt"

func isOdd(number int) string {
	if number%2 == 0 {
		return "It is "
	}
	return "Its Not"
}

func main() {
	result := isOdd(11)
	fmt.Println(result)
}


