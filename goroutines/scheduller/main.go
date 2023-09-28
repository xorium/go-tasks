package main

import (
	"fmt"
)

/*
Сделать так, чтобы гарантированно всегда выводились числа от 0 до 9 (в любом порядке).
*/

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
