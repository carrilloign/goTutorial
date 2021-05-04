package main

import (
	"errors"
	"fmt"
)

type User struct {
	name string
	ban bool
}

var users []User

var(
	NotValidUserError = errors.New("El usuario es invalido")
	DefaultError = errors.New("Default error")
)

func main() {
	users = append(users, User{"Nacho",false})
	users = append(users, User{"Pedro",false})
	users = append(users, User{"Juan",true})
	users = append(users, User{})

	//si se genera un error "lo catcheo"
	err:=isBanned(users)
	checkError(err)

}

func checkError(err error) {
	if err!=nil{
		fmt.Println("Error:",err)
	}
}

func isBanned(u []User) error{
	for _, user := range u {
		if user.name == "" {
			return NotValidUserError
		}
		switch user.ban {
		case true:
			fmt.Println("El usuario", user.name, "esta banneado")
		case false:
			fmt.Println("El usuario", user.name, "no esta banneado")
		}
	}
	return nil
}
