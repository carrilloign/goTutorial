package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	for i, numero := range slice {
		println("Componente:", i, "Valor:", numero)
	}

	//si no necesito el indice
	for _, numero := range slice {
		fmt.Println("Numero:", numero)
	}

	//si no necesito el numero
	for index := range slice{
		fmt.Println(index)
	}

	cadena:="Hola a todos"
	for index,letra:= range cadena{
		fmt.Printf("Letra:%q - Indice:%d\n",letra,index)
	}
}
