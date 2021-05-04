package main

import "fmt"

func main() {
	video25()
	video26()
	video28()
}

func video28() {
	fmt.Printf("La firma de la variadic suma es: %T\n",sumar)
	fmt.Println("El valor de la variadic suma es:",sumar(3,4,5))
	//pasandole slice a una func que recibe numero de par variables
	slice := [] int{1,2,3,4}
	fmt.Println("El valor de la variadic suma es:",sumar(slice...))

	//string variable
	slice2:= [] string{"el","mejor","lenguaje","->go"}
	res:=encadenar("Aprendiendo",slice2...)
	fmt.Println(res)

}

func video26() {
	resta:=restar(10,5)
	fmt.Printf("Firma de la funcion resta: %T\n",restar)
	fmt.Println("El valor de la resta es:",resta)
}

func video25() {
	imprimirNombre("Jose")
	valor:=suma(3,5)
	fmt.Printf("Firma de la funcion suma: %T\n",suma)
	println("El valor de la suma es:",valor)
}

//tipo void
func imprimirNombre(nombre string) {
	fmt.Println("El nombre es:",nombre)
}

//return int
func suma(n1 int, n2 int) int  {
	return n1+n2
}

//otro tipo return
func restar(n1 int, n2 int ) (r int){
	r= n1-n2
	return
}


//variadic func
func sumar(numeros ... int) (r int){
	for _,numero:= range numeros{
		r+=numero
	}
	return
}

//multiple variadic func
func encadenar(cadenaInicial string,adicional ... string) (r string){
	r=cadenaInicial
	for _,word:= range adicional{
		r+=" "+word
	}
	return
}

