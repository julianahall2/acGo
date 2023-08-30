package main

import (
	"fmt"
)

func main() {
	e_primo()
	fmt.Println("--------------------------")
	calculaFibo()
	fmt.Println("--------------------------")
	semana()
	fmt.Println("--------------------------")
	e_bissexto(1900)
}

func e_primo() {
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
	n := 5
	result := fibo(n)
	fmt.Printf("O %d-ésimo elemento da série de Fibonacci é: %d\n", n, result)
}

func fibo(n int) int {
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
	var num_dia = 2

	switch num_dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("Erro,dia inválido")
	}
}

func e_bissexto(ano int) bool {
	if ano%4 == 0 {
		if ano%100 != 0 || ano%400 == 0 {
			fmt.Println(ano, "é um ano bissexto.")
			return true
		}
	}
	fmt.Println(ano, "não é um ano bissexto.")
	return false
}
