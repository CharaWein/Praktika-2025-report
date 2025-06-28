package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Приветствую, я Вейнмайер"

	fmt.Println("Длина строки:", utf8.RuneCountInString(s))
	fmt.Println("Содержит 'Чара'?", strings.Contains(s, "Чара"))
	fmt.Println("Верхний регистр:", strings.ToUpper(s))
	fmt.Println("Нижний регистр:", strings.ToLower(s))
}
