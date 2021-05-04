package main

import "fmt"

func main() {
	a:=32

	fmt.Println("Valor de a:",a)
	fmt.Println("Direccion de memoria de a:",&a)

	b:=&a
	fmt.Println("Valor de b:",b)
	fmt.Println("Valor del puntero:",*b)
	fmt.Println("Direccion de memoria de b:",&b)

	*b=24
	fmt.Println("Valor de a:",a)

	//el valor cero es nil
	var puntero *int
	fmt.Println(puntero)

	//iniciar puntero
	puntero= new(int)
	fmt.Println(puntero)
	fmt.Println("Valor apuntado:",*puntero)

	//punteros como parametro
	numero:=20
	fmt.Println("Valor antes de incrementar:",numero)
	incrementar(&numero)
	fmt.Println("Valor despues de incrementar:",numero)
}

func incrementar(num *int) {
	*num++
}

