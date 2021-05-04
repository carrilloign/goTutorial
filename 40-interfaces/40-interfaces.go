package main

import "fmt"

type Persona struct {
	name string
	age int
}

type Admin struct {
	Persona
	email string
}

func (a Admin) Nombre() string  {
	return fmt.Sprintf("Mi nombre como admin es: %s",a.name)
}

type User interface{
	Nombre() string
	Age() string
}

func (p Persona) Nombre() string  {
	return fmt.Sprintf("Mi nombre es: %s",p.name)
}

func (p Persona) Age() string{
	return fmt.Sprintf("Mi edad es: %d",p.age)
}

func presentarse(u User){
	fmt.Println(u.Nombre())
	fmt.Println(u.Age())
}

func main() {
	persona := Persona{"nacho",25}
	admin := Admin{persona,"nacho@gmail.com"}

	presentarse(persona)
	presentarse(admin)

	/**********************************/
	fmt.Println("*****************")
	var i []User
	i = append(i, persona)
	i = append(i, admin)

	for j:=0; j< len(i); j++{
		presentarse(i[j])
	}
}
