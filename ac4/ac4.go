package main

import "fmt"

func main() {
	contador := 0
	var n int
	fmt.Scanln(&n)
	hanoi(n, "A", "B", "C", &contador)
	fmt.Println("Os discos foram movidos", contador, "vezes")
}

func hanoi(n int, origem string, destino string, trabalho string, contador *int) {

	if n > 0 {
		hanoi(n-1, origem, trabalho, destino, contador)
		fmt.Println("Mova o disco", n, "de", origem, "para", destino)
		*contador += 1
		hanoi(n-1, trabalho, destino, origem, contador)
	}
}
