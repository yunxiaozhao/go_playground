package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(os.Args[0])
	fmt.Println(-28 >> 12)
	fmt.Printf("%d %o %[1]o\n", 1, 2) // "438 666 0666"
	s := "hello, world"
	fmt.Println(len(s)) // "12"
	fmt.Println(comma("123123123"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
