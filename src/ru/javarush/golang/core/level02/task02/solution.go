package main

import "fmt"

func main() {
	// TODO: Задайте значения candies и kids (тип int) напрямую в коде так,
	// чтобы итоговый вывод программы был строго: 3 2.
	// TODO: Убедитесь, что kids != 0.
	var candies int = 8
	var kids int = 3

	// TODO: Вычислите eachCandies с помощью целочисленного деления candies / kids.
	eachCandies := candies / kids
	// TODO: Вычислите leftCandies с помощью остатка от деления candies % kids.
	leftCandies := candies % kids

	fmt.Println(eachCandies, leftCandies)
}