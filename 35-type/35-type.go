package main

import "fmt"

//declaracion de tipo

type Dinero int

//Implementamos el metodo string de dinero
func (d Dinero) String() string {
	return fmt.Sprintf("Su sueldo es: $%d",d)
}

func main() {
	var sueldo Dinero
	sueldo=25000
	fmt.Println(sueldo)

	aumento:=10000

	//para sumarlo a dinero debo castearlo
	sueldo+=Dinero(aumento)
	fmt.Println(sueldo)
}
