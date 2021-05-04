package main

import "fmt"

func main() {
	//Programa para averiguar numeros impares

	var inf, sup int
	fmt.Print("Ingrese el limite inferior de analisis: ")
	fmt.Scanln(&inf)
	fmt.Print("Ingrese el limite superior de analisis: ")
	fmt.Scanln(&sup)

	count := 0
	for i := inf; i <= sup; i++ {
		if i%2 != 0 {
			fmt.Printf("El numero %d es impar\n", i)
			count++
		}
	}
	fmt.Printf("La cantidad de numeros impares entre %d y %d es: %d\n", inf, sup, count)
}
