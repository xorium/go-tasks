package main

import "fmt"

// fanIn возвращает канал, который будет содержать все элементы из каналов channels (порядок между элементами
// не важен)
func fanIn(channels ...chan int) chan int {
	return nil
}

func main() {
	a := make(chan int, 3)
	b := make(chan int, 3)
	c := make(chan int, 3)

	a <- 1
	a <- 2
	a <- 3
	b <- 4
	b <- 5
	b <- 6
	c <- 7
	c <- 8
	c <- 9

	for num := range fanIn(a, b, c) {
		fmt.Println(num)
	}
}
