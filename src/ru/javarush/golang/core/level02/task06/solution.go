package main

import "fmt"

func main() {
	var wordFirst string = "Go"  // TODO: Присвойте значение по условию
	var wordSecond string = "is" // TODO: Присвойте значение по условию
	var wordThird string = "fun"  // TODO: Присвойте значение по условию

	_ = wordSecond
	_ = wordThird

	sentence := wordFirst + " " + wordSecond + " " + wordThird // TODO: Соберите итоговую строку конкатенацией через + с явными пробелами " "

	fmt.Println(sentence)
}