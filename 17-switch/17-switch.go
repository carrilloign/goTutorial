package main

import "fmt"

func main() {
	fmt.Print("Ingrese el numero de dia de la semana: ")
	var dia int
	fmt.Scanln(&dia)
	switch dia {
	case 1:
		fmt.Println("El dia es lunes")
	case 2:
		fmt.Println("El dia es martes")
	case 3:
		fmt.Println("El dia es miercoles")
	case 4:
		fmt.Println("El dia es jueves")
	case 5:
		fmt.Println("El dia es viernes")
	case 6:
		fmt.Println("El dia es sabado")
	case 7:
		fmt.Println("El dia es domingo")
	default:
		fmt.Println("El dia es INVALIDO")
	}

	//Otro switch
	numero:=4
	switch {
	case numero<=3:
		fmt.Println("Esto es otra condicion de numero")
	case true!=false:
		fmt.Println("otra condicion")
	}

}
