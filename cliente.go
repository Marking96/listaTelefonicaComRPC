package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Contato struct {
	Nome   string
	Numero int
}

type Search struct {
	Pnome string
}

var (
	reply   bool
	nm      string
	op, num int
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Falha ao conecta, verifique o endereço do servidor")
		os.Exit(1)
	}

	servidor := os.Args[1]

	cliente, err := rpc.Dial("tcp", servidor+":1234")
	if err != nil {
		log.Fatal("dialong", err)
	}

	for op != 5 {
		opcoes()
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Println("informe o nome: ")
			fmt.Scan(&nm)
			fmt.Println("informe o numero(somente numeros):")
			fmt.Scan(&num)

			Contatos := Contato{nm, num}

			err = cliente.Call("Agenda.Adiciona", Contatos, &reply)

			if reply {
				fmt.Println("Contato salvo com sucesso.")
			} else {
				fmt.Println("Falha ao salva!")
			}
		case 2:
			fmt.Println("informe o nome: ")
			fmt.Scan(&nm)
			var respo Contato
			Pequisa := Search{nm}
			err = cliente.Call("Agenda.Pesquisar", Pequisa, &respo)
			if respo.Nome == "" {
				fmt.Println("Não encontrado")
			} else {
				fmt.Printf("Nome: %s Numero: %d \n", respo.Nome, respo.Numero)
			}
			if err != nil {
				log.Fatal("arith error:", err)
			}
		case 3:
			fmt.Println("Informe o nome")
			fmt.Scan(&nm)
			Remove := Search{nm}
			reply = false
			err = cliente.Call("Agenda.Remover", Remove, &reply)
			if reply {
				fmt.Println("Contato removido com sucesso.")
			} else {
				fmt.Println("Falha ao remover")
			}
			if err != nil {
				log.Fatal("arith error:", err)
			}
		case 4:
			fmt.Println("-------CONTATOS----------\n")
			var result [] Contato
			err = cliente.Call("Agenda.Lista", _, &result)
		case 5:
			fmt.Println("Encerrando Cliente Lista telefonica")
		default:
			fmt.Println("Opção invalida")
		}
	}

}

func opcoes() {
	fmt.Println("Escolha uma opção: ")
	fmt.Println("1 - adiciona novo contato")
	fmt.Println("2 - Pesquisa contato")
	fmt.Println("3 - Remover contato")
	fmt.Println("4 - Listar todos os contatos")
	fmt.Println("5 - Sair")
}
