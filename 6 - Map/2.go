package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Умолчание это литературный приём недосказанности который даёт возможность слушателю или читателю додумать свою версию происходящего причём обычно умалчиваются явные понятия, которые можно легко определить по контексту"
	words := strings.Fields(text)
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	fmt.Println(freq)
}
