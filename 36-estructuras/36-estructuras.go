package main

import "fmt"

type Persona struct {
	nombre string
	edad int
}

type Estudiante struct{
	Persona
	carrera string
}

type Profesor struct {
	Estudiante
	carrera string
}

func older(p1,p2 Persona) Persona{
	if(p1.edad>p2.edad){
		return p1
	}
	return p2
}
func main() {
	video36()
	video37()
}

func video37() {
	//Composicion, simil herencia
	estudiante1:= Estudiante{Persona{"Ignacio",26},"Informatica"}
	fmt.Println("Nombre:",estudiante1.nombre)
	fmt.Println("Edad:",estudiante1.edad)
	fmt.Println("Carrera:",estudiante1.carrera)
	fmt.Println("Estudiante:",estudiante1)

	//Acceder a un atributo con el mismo nombre en clase hija
	profesor1 := Profesor{Estudiante{Persona{"Ignacio",24},"informatica"},"contabilidad"}
	fmt.Println(profesor1)
	fmt.Println("carrera:",profesor1.carrera)
	fmt.Println("carrera alumno:",profesor1.Estudiante.carrera)
}


func video36() {
	//1ra forma de inicializacion
	var persona1 Persona
	persona1.nombre="Igna"
	persona1.edad=23

	fmt.Println(persona1)
	fmt.Println("Nombre:",persona1.nombre)
	fmt.Println("Edad:",persona1.edad)

	//2da forma de inicializacion
	juan:= Persona{nombre: "Juan",edad: 25}
	fmt.Println(juan)

	//3ra forma de inicializacion
	ignacio:= Persona{"Ignacio",26}
	pedro:= Persona{"Pedro",28}

	fmt.Println("La persona mayor es:",older(ignacio,pedro))
}
