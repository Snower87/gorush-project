package main

/*
= Задача №15. Бюджет поездки =
Вы делаете простую печать сметы для командировки: сотрудник вводит город, сколько дней он будет в пути, и сколько денег в день ему выделили.
Программа читает city (одно слово), days (int) и perDay (int), затем считает total = days * perDay и печатает одну строку: Trip to <city>: <days> days, $<perDay>/day, total=$<total>
В конце строки должен быть перевод строки.
Для вывода используйте fmt.Printf, "%s" для города и "%d" для всех целых чисел.
Требования:
• Программа должна прочитать три значения из stdin через fmt.Scan (или fmt.Fscan): city (одно слово), days (int) и perDay (int).
• В решении должны быть переменные city типа string, days типа int и perDay типа int для хранения введённых значений.
• Программа должна вычислить total как произведение days * perDay.
• Программа должна вывести ровно одну строку с помощью fmt.Printf в формате: Trip to <city>: <days> days, $<perDay>/day, total=$<total> При этом символы двоеточия, запятых, $, и /day должны быть частью форматной строки.
*/

import "fmt"

func main() {
	var city string
	var days int
	var perDay int

	// Ввод: одно слово + два целых числа.
	fmt.Scan(&city, &days, &perDay)

	// TODO: Посчитайте общую стоимость поездки в переменной total.
	total := days * perDay

	// TODO: Выведите ровно одну строку в требуемом формате (через fmt.Printf) и завершите её переводом строки \n.
	fmt.Printf("Trip to %s: %d days, $%d/day, total=$%d\n", city, days, perDay, total)
}