package main

import "fmt"

func main() {
	i, p := make(chan int),	 make(chan int)
	pronto := make(chan bool)


	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go separar(nums, i, p, pronto)

	var impares, pares []int

	fim := false

	for !fim {
		select {
			case n := <- i:
				impares = append(impares, n)
				fmt.Println("Impar adicionado")
			case n := <- p:
				pares = append(pares, n)
				fmt.Println("Par adicionado")
			case fim = <- pronto:
				fim = true
				fmt.Println("Fim do processamento")
		}
	}

	fmt.Printf("Impares: [%v] | Pares: [%v]", impares, pares)
}


func separar(nums []int, i, p chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n % 2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}

	pronto <- true
}
