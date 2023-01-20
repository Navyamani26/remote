package main

import "fmt"

func positive(a int) {
	if a < 0 {
		fmt.Println("it is a negative number")
	} else {
		fmt.Println("it is a positive number")
	}
}
