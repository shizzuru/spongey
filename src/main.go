package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Super Item Randomizer 3000 ---\n")

	oui := make(map[string]int)

	for i := 0; i < 100; i++ {
		item := newItem("Butcher's choice")
		oui[item.getStatedName()]++
	}
	for k, v := range oui {
		fmt.Printf("%v :\t\t %d\n", k, v)
	}
}
