package main

import (
	"fmt"
)

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n >= M {
		fmt.Println("Overflow")
	} else {
		if busca(*v, *n, novoValor) != -1 {
			i := *n - 1
			for i >= 0 && v[i] > novoValor {
				v[i+1] = v[i]
				i--
			}
			v[i+1] = novoValor
			*n++
		} else {
			fmt.Println("Elemento já existe na tabela")
		}
	}
}

func busca(v [M]int, n int, x int) int {
	i := 0
	for i := 0; i < n; i++ {
		if v[i] == x {
			return -1
		}
		if v[i] > x {
			return i
		}
	}
	return i
}
