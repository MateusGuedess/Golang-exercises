package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	palavras := os.Args[1:]

	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)
}

func colherEstatisticas(palavras[]string) (map[string]int) {

	estatisticas := make(map[string]int)

	for _, palavra := range palavras {
		inicial := strings.ToUpper(palavra[:1])
		contador, encontrado := estatisticas[inicial]

		if encontrado {
			estatisticas[inicial] = 1 + contador
		} else {
			estatisticas[inicial] = 1
		}
	}

	return estatisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra: ")

	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}