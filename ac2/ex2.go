package main

import "fmt"

type Contato struct {
	Nome	string
	Email	string
}

func (c *Contato) Alterar(contatoEmail string) {
	c.Email = contatoEmail
}

func main() {
	var contato Contato
	contato := nome, email
	var nome, email string
		fmt.Println("Digite o nome do contato:")
		fmt.Scan(&nome)
		fmt.Println("Digite o e-mail do contato:")
		fmt.Scan(&email)

	contatos [5]Contato

	contatos[0] = contato

	for _, contato := range contatos {
		if (contato != Contato{}) {
			fmt.Println(contato)
		}
	}

}


func adicionar(contato,contatos [0]Contato) [5]Contato {

}
func excluirContato(contatos []Contato) []Contato {
	if len(contatos) > 0 {
		contatos = contatos[:len(contatos)-1]
	} else {
		fmt.Println("Não há contatos para excluir.")
	}
	return contatos
}

