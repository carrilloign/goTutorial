package main

import (
	"fmt"
	"time"
)

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
		for{
			n,state:=<-numero
			//pregunto si el canal esta abierto
			if !state{
				break //si esta cerrado salgo del for
			}
			cuadrado<- n*n
		}
		close(cuadrado) //cierro el otro canal
	}()

	for{
		sq,state:=<-cuadrado
		if!state{
			break //salgo del for
		}
		fmt.Println("Cuadrado:",sq)
		time.Sleep(1*time.Second)
	}

}
