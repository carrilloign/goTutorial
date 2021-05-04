package main

import "fmt"

func main() {
	//Declaracion variables
	var numero int = 3
	var numero1 = 4
	nombre := "Nacho"

	fmt.Println(numero, numero1)
	fmt.Println("El nombre es:", nombre)

	//Reasignacion multiple
	nombre, numero = "Ignacio", 26
	fmt.Println("Nombre:", nombre, ". Numero:", numero)

	//Declaracion multiple
	var numero2, numero3, numero4 int = 3, 4, 5
	numero5, palabra := 5, "hola"

	var (
		numero6 = 6
		otrapalabra= "Otra palabra"
	)

	fmt.Println("Numeros:", numero2, numero3, numero4)
	fmt.Println("Numero: ", numero5, "Palabra:", palabra)
	fmt.Println("Numero: ", numero6, "Palabra:", otrapalabra)

	//Intercambiar valores
	fmt.Println("Palabra: ", palabra, "Nombre:", nombre)
	palabra, nombre = nombre, palabra
	fmt.Println("Palabra: ", palabra, "Nombre:", nombre)
}
