package main

import (
	"fmt"
	"time"
)

//mismo ejemplo de channels2 pero con foreach
func main() {
	numero:= make(chan float64)
	cuadrado:= make(chan float64)

	go func(){
		for x:=1;x<=5 ;x++{
			numero<-float64(x)
		}
		close(numero) //cierro el canal
	}()

	go func(){
		for num:= range numero{
			cuadrado<-num*num
		}
		close(cuadrado) //cierro el otro canal
	}()

	for it:= range cuadrado{
		fmt.Println("Cuadrado:",it)
		time.Sleep(1*time.Second)
	}

}
