package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(os.Args[0])
	loopTest()
}

func loopTest() {
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