package main

import "fmt"

func main() {

	//declaracion
	var vector [4]int
	var matriz [3][2]int
	vector2 := [3]int{1, 2, 3}
	vector4 := [3]int{3, 2, 1}

	//declaracion sin tamaño. Debo pasarle los elementos para q determine la leng
	vector3 := [...]string{"Palabra1", "Palabra2"}
	fmt.Println(vector)
	fmt.Println(matriz)
	fmt.Println(vector2)
	fmt.Println(vector2[0])
	fmt.Println(vector3)
	fmt.Println("len:", len(vector3))
	fmt.Println("ultimo valor:", vector3[len(vector3)-1])
	fmt.Println("Son iguales? ", vector4 == vector2)
}
