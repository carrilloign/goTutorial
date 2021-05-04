package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

type Rectangulo struct {
	base float64
	altura float64
}

func (c Circulo) area() float64 {
	return math.Pow(c.radio,2)*math.Pi
}

func (r Rectangulo) area() float64 {
	return r.base*r.altura
}

func (r *Rectangulo) incrementarArea(){
	r.base++
	r.altura++
}

type Punto struct {
	x,y int
}

func main() {
	video38()
	video39()
}

func video39() {
	punto:= Punto{}
	fmt.Println("Punto:",punto)
}

func video38() {
	rectangulo:= Rectangulo{2,3.5}
	circulo:= Circulo{2.5}
	fmt.Println("Area Rectangulo:",rectangulo.area())
	fmt.Println("Area Circulo:",circulo.area())
	fmt.Println("Rectangulo:",rectangulo)
	rectangulo.incrementarArea()
	fmt.Println("Rectangulo:",rectangulo)
}
