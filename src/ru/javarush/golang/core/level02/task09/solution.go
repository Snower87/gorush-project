package main

import "fmt"

func main() {
	var n int // по требованию: объявляем переменную до fmt.Scan

	fmt.Scan(&n)

	// TODO: Выведите введённое число n одной строкой (с переводом строки) без лишнего текста.
	fmt.Println(n)
}