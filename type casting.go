package main

import "fmt"

func typecasting(a int, b float64, s string) {
	a1 := int(b)
	b1 := float64(a)
	fmt.Println("Before typecasting\n")
	fmt.Printf("The type: %T  value :%v\n", a, a)
	fmt.Printf("The type: %T  value :%v\n", b, b)
	fmt.Printf("The type: %T  value :%q\n", s, s)

	fmt.Println("after typecasting\n")
	fmt.Printf("The type: %T  value :%v\n", b1, b1)
	fmt.Printf("The type: %T  value :%v\n", a1, a1)
	fmt.Printf("The type: %T  value :%v\n", s, s)
}
