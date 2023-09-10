package main

import "fmt"

func main() {
	listaNumeros := []int{1, 2, 3, 4, 5}
	fmt.Println(acharPair(listaNumeros, 3, 9))
	fmt.Println("------------------")
	fmt.Println(acharPair(listaNumeros, 5, 10))
	fmt.Println(acharPair(listaNumeros, 2, 9))
	fmt.Println(acharPair(listaNumeros, 3, 4))
}

func acharPair(matriz []int, n int, alvo int) (int, int) {
	numAnterior, numSeguinte := 0, len(matriz)-1

	for numAnterior < numSeguinte {
		somaInicial := matriz[numAnterior] + matriz[numSeguinte]

		if somaInicial == alvo {
			return matriz[numAnterior], matriz[numSeguinte]
		} else if somaInicial < alvo {
			numAnterior++
		} else {
			numSeguinte--
		}
	}
	return -1, -1
}
