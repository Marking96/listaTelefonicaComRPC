package main

import (
	"errors"
	"fmt"
	"net/rpc"
	//"net/http"
	"net"
	//"net/http"
	"os"
)

type Contato struct {
	nome string
	numero int
}


type Agenda Contato


func (t *Agenda) adiciona(contato *Contato, agenda[] *Contato) error {
	if contato.nome == "" {
		return errors.New("Nome do novo contato não informado")
	}
	agenda = append(agenda,contato)
	return nil
}

func main() {

	ag := new(Agenda)
	rpc.Register(ag)
	//rpc.HandleHTTP()

	//err:= http.ListenAndServe(":1234",nil)
	tcpAddr, err := net.ResolveTCPAddr("tcp",":1234")
	checkError(err)

	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)

	for  {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkError(err error)  {
	if err != nil {
		fmt.Println("Fatal Erro: ",err.Error())
		os.Exit(1)
	}
}