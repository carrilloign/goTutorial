package main

import "fmt"

func main() {
	a := 3
	b := 2
	fmt.Println("A:", a, "B:", b)
	fmt.Println("Suma:", a+b)
	fmt.Println("Resta:", a-b)
	fmt.Println("Mult:", a*b)
	fmt.Println("Div:", a/b)
	fmt.Println("Resto:", a%b)

	fmt.Println("***************")

	fmt.Println("A:", a, "B:", b)
	fmt.Println("Igual", a == b)
	fmt.Println("Distinto", a != b)
	fmt.Println("Mayor o igual", a >= b)
	fmt.Println("Mayor", a > b)
	fmt.Println("Menor o igual", a <= b)
	fmt.Println("Menor", a < b)

}
