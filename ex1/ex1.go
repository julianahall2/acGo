package main

import (
	"fmt"
)

func main() {
	e_primo()
	fmt.Println("--------------------------")
	calculaFibo()
	fmt.Println("--------------------------")

}


func e_primo () {
	var n int
	fmt.Println("Digite um número : ")
	fmt.Scan(&n)

	mult := 0

	for count := 2; count < n; count++ {
		if n%count == 0 {
			fmt.Println("Múltiplo de", count)
			mult++
		}
	}

	if mult == 0 {
		fmt.Println("É primo")
	} else {
		fmt.Println("Não é primo")
	}
}

func calculaFibo() {
	n := 10
	result := fibo(n)
	fmt.Printf("O %d-ésimo elemento da série de Fibonacci é: %d\n", n, result)
}

func fibo(n int) int{
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	prev1, prev2 := 1, 1
	current := 0

	for i := 3; i <= n; i++ {
		current = prev1 + prev2
		prev1, prev2 = prev2, current
	}

	return current

}

func semana() {
	var dia = "segunda"
	var 

	switch dia {
	case "segunda", "terça","quarta","quinta", "sexta":
		fmt.Println("Dia de semana")
	case "sabado", "domingo":
		fmt.Println("Final de semana")
	default:
		fmt.Println("Erro,dia inválido")
	}
}