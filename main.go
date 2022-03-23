package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func choiceMaker(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func main() {
	// example data
	choices := []string{"blue", "orange", "white"}

	// make a random choice
	res := choiceMaker(choices)

	// pick up the choice and don't worry, be happy.
	fmt.Println(res)
}
