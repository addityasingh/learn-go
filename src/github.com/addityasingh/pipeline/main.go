package main

import (
	"fmt"
)

// IntChannel type for integer channels
type IntChannel chan int

func genNumbers(num ...int) IntChannel {
	out := make(IntChannel)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(num IntChannel) IntChannel {
	out := make(IntChannel)
	go func() {
		for n := range num {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("Print square of numbers upto 4")
	num := genNumbers(1, 2, 3, 4)
	sqChan := square(num)

	for n := range sqChan {
		fmt.Println(n)
	}

	fmt.Println("Print square of squares of numbers upto 4")
	num = genNumbers(1, 2, 3, 4)
	for n := range square(square(num)) {
		fmt.Println(n)
	}
}
