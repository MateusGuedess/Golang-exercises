package main

import "os"

func main() {
	palavras := os.Args[1:]

	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)
}

func colherEstatisticas(palavras[]string) (map[string]int) {

	estatisticas := make(map[string]int)

	for _, palavra := range palavras {
		
	}
}