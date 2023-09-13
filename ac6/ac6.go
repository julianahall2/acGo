package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	fmt.Println(busca(lista, 5, 17))
	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n == 0 {

	}
}

func busca(v [M]int, n int, x int) int {
	i := 0
	v[n] = x
	for v[i] < x {
		i++
	}
	if i == n+1 || v[i] != x {
		return -1
	}
	return i
}
