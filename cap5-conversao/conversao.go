package main

import "fmt"

type ListaDeCompras []string


func imprimirSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func imprimirLista(lista []string) {
	fmt.Println("Lista: ", lista)
}

func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"} 

	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}