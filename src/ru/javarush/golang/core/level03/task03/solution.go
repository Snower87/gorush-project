package main

/*
= Задача №3. Датчик температуры =
Вы настраиваете "умный термометр" для лаборатории. Он получает температуру в целых градусах, а дальше система хочет сразу
четыре простых булевых сигнала: холодно ли "ниже нуля", кипит ли "сотка и выше", и попадает ли значение в комфортные границы.
Программа читает одно целое число currentTemperature. Далее она вычисляет четыре булевых значения:
isFreezing — currentTemperature < 0
isBoiling — currentTemperature >= 100
isComfortLow — currentTemperature >= 18
isComfortHigh — currentTemperature <= 26
После этого программа выводит эти четыре значения на одной строке через пробел в указанном порядке.
Ввод: одно целое число. Вывод: 4 булевых значения.
Пример ввода:
20
Пример вывода:
false false true true
Требования:
• Программа должна прочитать одно целое число `currentTemperature` с помощью `fmt.Scan`, при этом переменная `currentTemperature`
должна иметь тип `int`.
• Программа должна вычислить четыре значения типа `bool` через операции сравнения и сохранить каждое в отдельную переменную:
`isFreezing` (`currentTemperature < 0`), `isBoiling` (`currentTemperature >= 100`), `isComfortLow` (`currentTemperature >= 18`),
`isComfortHigh` (`currentTemperature <= 26`).
• В решении не должны использоваться `if/else`, а также логические операторы `&&`, `||`, `!` — только сравнения (`<`, `<=`, `>=`, при необходимости `>`).
• Программа должна вывести ровно 4 значения `true/false` в одной строке, разделяя их одним пробелом, строго в порядке: `isFreezing`, `isBoiling`, `isComfortLow`, `isComfortHigh`.
*/

import "fmt"

func main() {
	var currentTemperature int
	fmt.Scan(&currentTemperature)

	// TODO: Вычислите isFreezing через сравнение currentTemperature с 0 (без if/else и без && || !).
	isFreezing := false
	isFreezing = currentTemperature < 0

	// TODO: Вычислите isBoiling через сравнение currentTemperature со 100 (без if/else и без && || !).
	isBoiling := false
	isBoiling = currentTemperature >= 100

	// TODO: Вычислите isComfortLow через сравнение currentTemperature с 18 (без if/else и без && || !).
	isComfortLow := false
	isComfortLow = currentTemperature >= 18

	// TODO: Вычислите isComfortHigh через сравнение currentTemperature с 26 (без if/else и без && || !).
	isComfortHigh := false
	isComfortHigh = currentTemperature <= 26

	// TODO: Проверьте, что выводите ровно 4 значения, через один пробел, в нужном порядке.
	fmt.Printf("%t %t %t %t", isFreezing, isBoiling, isComfortLow, isComfortHigh)
}