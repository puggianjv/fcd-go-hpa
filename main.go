package main

import (
	"fmt"
	"math"
)

func main() {
	doSqrt()
	fmt.Println("Code.education rocks!")
}

func doSqrt() int {
	x := 0.0001
	i := 0
	for i <= 1000000 {
		x = math.Sqrt(x)
		i++
	}
	return i
}
