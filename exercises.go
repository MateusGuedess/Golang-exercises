package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valorOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string


	switch unidadeOrigem {
	case "celsius":
		unidadeDestino = "fahrenheit"
		break
	case "quilometros":
		unidadeDestino = "milhas"
		break
	default:
		fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}

	for i, v := range valorOrigem{
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não e um numero valido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = (valorOrigem * 1.8) + 32
		} else {
			valorDestino = valorOrigem * 1.621371
		}

		fmt.Printf ("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}



}