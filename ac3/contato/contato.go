package contato

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func alteraEmail(msg *string) {
	*msg = strings.ReplaceAll(*msg, "email", "novoEmail")
}