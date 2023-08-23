package main

import (
	"fmt"
	ctt "AC3/contato"
	"AC3/array"
)

func main() {
	contato1 := ctt.Contato{
		Nome:  "Juliana",
		Email: "juju.azeredohall@gmail.com",
	}

	novoEmail := "emailnovo@example.com"
	ctt.alteraEmail(&contato1, novoEmail)

	fmt.Printf("Novo email:", contato1.Email)


}
