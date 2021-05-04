package main

import "fmt"

func main() {

	//defer
	defer primera()
	segunda()

	//capturo el panic para que no se frene la ejecucion del programa
	defer func(){
		cadena:=recover()
		fmt.Println("El error fue:",cadena)
	}()
	panic("KERNEL PANIC !")
}

func primera(){
	fmt.Println("Primera funcion")
}
func segunda(){
	fmt.Println("Segunda funcion")
}
