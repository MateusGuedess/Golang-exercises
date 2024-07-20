package main

import "fmt"

func main() {


	gerarFib := GerarFibonacci(20)

	gerarFib()

}

func GerarFibonacci(qtGerar int) (func()) {
	return func() {
		a, b := 0, 1;

		fib := func() int {
			a, b = b, a + b
	
			return a
		}


		for i := 0; i < qtGerar; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}