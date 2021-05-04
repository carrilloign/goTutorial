package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "Hello to closure's world"
	imprimir := print //una forma de meter una funcion dentro de otra
	imprimir(cadena)

	//otra forma - Debido al scope puedo acceder directamente a la variable
	imprimir2 := func() {
		fmt.Println(cadena + ". Ejemplo 2")
	}
	imprimir2()

	// como print y print2 tienen la misma firma, puedo hacer que
	imprimir = print2
	imprimir(cadena)
	fmt.Println("*********************")
	fmt.Println("Diferencia de firmas:")
	fmt.Printf("Firma 1: %T\n",imprimir)
	fmt.Printf("Firma 2: %T\n",imprimir2)
	fmt.Println("*********************")
	print3(print)

	//closure
	otra:= imprimirDentro()
	for i:=0;i<4;i++{
		otra()
	}

	//funciones anonimas
	editarCadena:="ABCD"
	sumador:=1

	editarCadena = strings.Map(func(r rune)rune{
		return r+rune(sumador)
	},editarCadena)

	fmt.Println(editarCadena)
}

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena string) {
	fmt.Println("Escribo otra cosa concatenada a " + cadena)
}

//funcion que recibe como parametro una funcion q imprime un valor
func print3(printExt func(string)){
	printExt("Escribiendo externamente desde print3")
}

//closure
func imprimirDentro() func() int{
	i:=0
	return func() (r int) {
		r=i
		i++
		fmt.Println("Valor de r:",r)
		return
	}
}
