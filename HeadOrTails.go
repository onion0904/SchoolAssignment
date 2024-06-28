package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Tossing a coin...")
	rand.Seed(time.Now().UnixNano())
	var count1, count2 int
	for i := 0; i < 3; i++ {
		R := rand.Intn(100)
		if R%2 == 0 {
			fmt.Printf("Round %d: Heads\n", i+1)
			count1++
		} else {
			fmt.Printf("Round %d: Tails\n", i+1)
			count2++
		}
	}
	fmt.Printf("Heads: %d, Tails: %d\n", count1, count2)
}
