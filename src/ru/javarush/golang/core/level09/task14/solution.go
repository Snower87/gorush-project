package main

/*
= Задача №14. Быстрый отчёт =
Вы пишете консольный отчёт для маленькой команды: в начале строки нужен «префикс» (например, название отдела), а дальше — сумма нескольких чисел. Самое важное требование от тимлида: в printSum нельзя пересчитывать сумму вручную — нужно использовать уже готовую sum(nums ...int) и прокинуть все числа дальше через nums..., чтобы не плодить дублирование логики.
Реализуйте sum(nums ...int) int, которая возвращает сумму всех чисел (а пустой вызов sum() возвращает 0). Реализуйте printSum(prefix string, nums ...int), которая печатает строку вида "<prefix> <sum>\n", получая сумму вызовом sum(nums...). Программа читает prefix и count (0..3), затем читает нужное количество чисел и вызывает printSum с 0..3 аргументами.
Требования:
• Должна быть реализована функция sum(nums ...int) int, возвращающая сумму всех переданных чисел; при пустом вызове sum() должна возвращать 0.
• Должна быть реализована функция printSum(prefix string, nums ...int), принимающая префикс и 0..N целых чисел.
• Функция printSum должна получать сумму исключительно вызовом sum(nums...) с прокидыванием аргументов через nums...; ручной подсчёт в printSum и вызов вида sum(nums) не допускаются.
• Программа должна прочитать prefix и count (0..3), затем прочитать ровно count целых чисел и вызвать printSum: при count==0 как printSum(prefix), при count==1/2/3 как printSum(prefix, a) / printSum(prefix, a, b) / printSum(prefix, a, b, c).
*/

import (
	"fmt"
	"os"
)

func sum(nums ...int) int {
	// TODO: Реализуйте подсчёт суммы всех чисел из nums.
	// TODO: Пустой вызов sum() должен возвращать 0.
	sum := 0
	for _, num := range nums {
	    sum += num
	}
	return sum
}

func printSum(prefix string, nums ...int) {
	// TODO: Напечатайте ровно одну строку в формате: "<prefix> <sum>\n".
	// TODO: Сумму нужно получать только вызовом sum(nums...) (с прокидыванием nums...).
	fmt.Printf("%s %d\n", prefix, sum(nums...))
}

func main() {
	var prefix string
	var count int
	fmt.Fscan(os.Stdin, &prefix, &count)

	switch count {
	case 0:
		printSum(prefix)
	case 1:
		var a int
		fmt.Fscan(os.Stdin, &a)
		printSum(prefix, a)
	case 2:
		var a, b int
		fmt.Fscan(os.Stdin, &a, &b)
		printSum(prefix, a, b)
	case 3:
		var a, b, c int
		fmt.Fscan(os.Stdin, &a, &b, &c)
		printSum(prefix, a, b, c)
	}
}