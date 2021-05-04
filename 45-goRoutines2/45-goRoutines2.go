package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp","localhost:8080") //abro una conexion
	if err!=nil{
		log.Fatal(err) //Detiene el programa si hubo un error
	}
	for { //bucle infinito
		conn,err := listener.Accept()  //se para el programa aqui esperando nuevas conexiones
		if err!=nil{
			log.Println(err)
			continue
		}
		go manejarCon(conn) //realizo la respuesta del listener
	}
}

func manejarCon(conn net.Conn) {
	defer conn.Close()
	for {
			for _,dig := range `\|/-` {
			_, err := io.WriteString(conn,string(dig)+"\r")
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
