package array

type Contato struct {
	Nome  string
	Email string
}

func adicionaContato(c Contato,a*[5]Contato){
	for ind,contato := range *a {
		if (contato == nil) {
			(*a)[ind] = &c
			break
		}
	}
}

func excluiContato(a *[5]*Contato) {
	if (*a)[0] == nil {
		fmt.Println("Lista de contatos vazia!")
		return
	}

	for ind, contato := range *a {
		if contato == nil {
			(*a)[ind-1] = nil
		}
	}
}
