package main

import "fmt"

func main() {

	//1ra declaracion map

	x:= make(map[int]string)
	fmt.Println(x)

	x[0]="Zero"
	x[1]="One"
	x[12]="Twelve"
	fmt.Println(x)
	fmt.Println(x[12])

	//2do tipo declaracion
	y:= map[int]string{0:"Zero",1:"One",2:"Two",3:"Three"}
	delete(y,2)
	y[4]+="carlos"
	fmt.Println(y) //imprime una cadena vacia

	//si el elemento buscado no existe devuelve el 0 del tipo de dato value
	fmt.Println(y[12])
	z:= make(map[int]int)
	fmt.Println(z[3])

	//existencia de clave, con for e if
	nombres:= map[int]string{1:"Carlos",2:"Ruben",3:"Rosario"}
	for j:=-2;j<=len(nombres);j++ {
		if _, ok := nombres[j]; ok {
			fmt.Println("El nombre existe")
		} else{
			fmt.Println("No existe nombre en el indice",j)
		}
	}

	//for each
	for clave,valor := range nombres{
		fmt.Println("Clave:",clave,"Valor:",valor)
	}
}
