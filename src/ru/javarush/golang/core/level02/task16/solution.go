package main

/*
= Задача №16. Зарплата и налог =
Вы делаете консольный расчёт для бухгалтерского черновика. Три числа приходят как строки: baseStr (оклад), bonusStr (премия)
и taxRateStr (ставка налога в процентах). Их нужно распарсить, посчитать общий доход, налог и “чистыми”, а параллельно
— вывести, были ли ошибки парсинга по каждому полю.
Программа должна прочитать три строки, распарсить их через strconv.Atoi в base, bonus, rate и ошибки errBase, errBonus,
errRate. Затем вычислить:
- income = base + bonus
- tax = income * rate / 100
- net = income - tax
и вывести четыре строки:
- income=<income>
- tax=<tax>
- net=<net>
- errs=<errBase>|<errBonus>|<errRate>
Требования:
• Программа должна прочитать ровно три значения типа string через fmt.Scan в переменные baseStr, bonusStr и taxRateStr (или эквивалентные по смыслу).
• Программа должна выполнить три отдельных вызова strconv.Atoi — по одному для baseStr, bonusStr и taxRateStr — сохранив результаты в base, bonus, rate и соответствующие ошибки errBase, errBonus, errRate.
• Программа должна вычислить income = base + bonus, tax = income * rate / 100 и net = income - tax, используя только int-арифметику; деление при вычислении tax должно быть целочисленным (оператор / без использования float-типов).
• Программа должна формировать строки для income, tax и net через strconv.Itoa и конкатенацию строк, без пробелов вокруг символа '=' (например, "income="+strconv.Itoa(income)).
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var baseStr, bonusStr, taxRateStr string

	// По условию читаем ровно 3 строки
	fmt.Scan(&baseStr, &bonusStr, &taxRateStr)

	// TODO: Распарсите baseStr в base через strconv.Atoi и сохраните ошибку в errBase (отдельный вызов Atoi).
	base := 0
	var errBase error
	base, errBase = strconv.Atoi(baseStr)

	// TODO: Распарсите bonusStr в bonus через strconv.Atoi и сохраните ошибку в errBonus (отдельный вызов Atoi).
	bonus := 0
	var errBonus error
	bonus, errBonus = strconv.Atoi(bonusStr)

	// TODO: Распарсите taxRateStr в rate через strconv.Atoi и сохраните ошибку в errRate (отдельный вызов Atoi).
	rate := 0
	var errRate error
	rate, errRate = strconv.Atoi(taxRateStr)

	// TODO: Вычислите income, tax и net только целочисленной арифметикой (без float).
	income := base + bonus
	tax := income * rate / 100
	net := income - tax

	// TODO: Сформируйте строки для income, tax и net через strconv.Itoa и конкатенацию без пробелов вокруг '='.
	fmt.Println("income=" + strconv.Itoa(income))
	fmt.Println("tax=" + strconv.Itoa(tax))
	fmt.Println("net=" + strconv.Itoa(net))

	// TODO: Выведите errs в формате errs=<errBase>|<errBonus>|<errRate> в указанном порядке.
	fmt.Println("errs=" + fmt.Sprint(errBase) + "|" + fmt.Sprint(errBonus) + "|" + fmt.Sprint(errRate))
}