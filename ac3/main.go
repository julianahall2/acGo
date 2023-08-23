package main

import (
	"fmt"
	"bufio"
	"os"
	ctt "AC3/contato"
	"AC3/array"
	"strings"
)

func main() {
	contato1 := ctt.Contato{
		Nome:  "Juliana",
		Email: "juju.azeredohall@gmail.com",
	}

	novoEmail := "emailnovo@example.com"
	ctt.alteraEmail(&contato1, novoEmail)

	fmt.Printf("Novo email:", contato1.Email)

	var contatos [5]*ctt.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = array.adicionaContato(ctt.Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = array.excluiContato(contatos)
		default:
			return
		}

		fmt.Println(contatos)
	}
}
