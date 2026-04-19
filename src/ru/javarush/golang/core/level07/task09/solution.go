package main

/*
= Задача №9. Упаковка склада =
Вы отвечаете за упаковку заказов на складе. В систему пришло число itemsToPack — сколько предметов нужно разложить по коробкам, и boxCapacity — сколько предметов помещается в одну коробку. Менеджеру нужны два ответа: сколько получится полностью заполненных коробок и сколько предметов останется “на следующий заход”.
Программа читает два целых числа, считает fullBoxes как целую часть деления, а leftoverItems как остаток, и выводит оба числа в одной строке через пробел.
Требования:
• Программа должна импортировать пакет fmt и использовать его для ввода и вывода.
• Программа должна считать из stdin два целых числа `itemsToPack` и `boxCapacity` с помощью `fmt.Scan`.
• Программа должна вычислить `fullBoxes` как целую часть деления `itemsToPack / boxCapacity`, используя оператор `/` и целочисленную арифметику.
• Программа должна вычислить `leftoverItems` как остаток от деления `itemsToPack % boxCapacity`, используя оператор `%` и целочисленную арифметику.
• Программа должна вывести `fullBoxes` и `leftoverItems` в одной строке, ровно два числа, разделённые одним пробелом.
*/

import "fmt"

func main() {
	var itemsToPack, boxCapacity int
	fmt.Scan(&itemsToPack, &boxCapacity)

	var fullBoxes int
	var leftoverItems int

	// TODO: Посчитайте количество полностью заполненных коробок (целочисленное деление itemsToPack на boxCapacity).
	// TODO: Посчитайте количество оставшихся предметов (остаток от деления itemsToPack на boxCapacity).
	fullBoxes = itemsToPack / boxCapacity
	leftoverItems = itemsToPack % boxCapacity

	fmt.Print(fullBoxes, " ", leftoverItems)
}