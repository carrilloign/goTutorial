package main

import "fmt"

func main() {

	var slice1 []int
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3}
	slice1 = slice2[:2]
	slice1[1] = 20
	fmt.Println(slice2)
	fmt.Println("ultimo valor:", slice2[len(slice2)-1])

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice[%d]: %d\n",i,slice2[i])
	}


	//ejemplo para analizar el concepto de capacidad
	var slice3 []int
	for i := 0; i < 10; i++ {
		slice3=append(slice3,i)
		fmt.Println("Slice:",slice3)
		fmt.Printf("Long: %d, Cap: %d\n",len(slice3),cap(slice3))
	}

	//ejemplo para analizar la funcion copy de slice igual tamaño
	origen:= []string{"Origen"}
	destino:= [] string{"Destino"}
	copy(destino,origen)
	fmt.Println("Destino:",destino)

	//diferentes tamaño
	origen2:= [] int {1,2,3}
	destino2:= [] int {2,4,6,8}

	copy(destino2,origen2)
	fmt.Println("Destino:",destino2)



}
