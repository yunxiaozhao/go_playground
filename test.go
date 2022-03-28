package main

import "fmt"

func LoopTest() {
	i := 1
	for {
		if i == 10 {
			break
		} else {
			i++
		}
	}
	fmt.Println("Loop finished.")
}
