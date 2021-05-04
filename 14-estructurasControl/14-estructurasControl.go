package main

import (
	"fmt"
	"strconv"
)

func main() {
	bucleFor()
	bucleIf()
}

func bucleFor() {
	//El unico bloque de control que existe es for.
	//con este se hacen while, foreach, dowhile etc

	//while
	i := 0
	for i < 5 {
		i++
		fmt.Print(strconv.Itoa(i) + ", ")
	}
	println()

	//for var, cond, it
	for i := 1; i <= 5; i++ {
		fmt.Print(strconv.Itoa(i) + ", ")
	}
	println()

}

func bucleIf() {
	// if else if

	//if simple
	b:= true
	if b{
		fmt.Println(!b)
	}

	//if all in line
	if a := 5; a < 6 {
		fmt.Println("a menor q 6")
	} else if a > 6 {
		fmt.Println("a mayor q 6")
	} else {
		fmt.Println("son iguales")
	}
}
