package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Умолчание это литературный приём недосказанности который даёт возможность слушателю или читателю додумать свою версию происходящего"
	words := strings.Fields(sentence)

	for _, word := range words {
		fmt.Println(word)
	}
}
