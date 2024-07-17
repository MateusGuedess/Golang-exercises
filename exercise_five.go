package main

import (
	"fmt"
	"time"

	"math/rand"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)

		
		if 1 % 42 == 0 {
			break;
		}
		fmt.Println("O valor é: %d", n)
		
	}


}