package main

import (
	"fmt"
	"math/rand"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	for {
// 		fmt.Println("Hello", os.Args[1])
// 		// os.Exit(42)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {
	a := rand.Intn(100)
	b := rand.Intn(100)

	m := IntMin(a, b)
	fmt.Println("Min of ", a, " and ", b, " is: ", m)
}
