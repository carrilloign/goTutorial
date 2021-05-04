package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena = "Hello World"
	fmt.Println("Size:",len(cadena))
	fmt.Println("H:",cadena[0]) //Nos muestra en decimal el valor del byte ascii
	fmt.Println("H:",cadena[:5]) //hasta el 5 caracter sin incluirlo
	fmt.Println("H:",cadena[6:]) //Desde el 6 caracter incluido

	cadena+=" Nuevamente"
	fmt.Println("Concatenada:",cadena) //Desde el 6 caracter incluido

	//comilla invertida permite literalidad, ignora caracteres especiales. etc
	cadena2:=`
				hola
				 como te va
				?? `
	fmt.Println(cadena2)

	//simbolos especiales
	//slash invertido \n \t
	cadena3:="Literal \"\\25\""
	fmt.Println(cadena3)

	cadena4 := "Edad: "
	numero := 4
	fmt.Println(cadena4+strconv.Itoa(numero))
}
