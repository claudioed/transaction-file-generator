package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print((rand.Float64()*5)+5, ",")
	}
}
