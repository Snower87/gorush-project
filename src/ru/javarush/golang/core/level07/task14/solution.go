package main

/*
= Задача №14. Деление отчёта =
Вы делите "общий объём" на "число частей" и замечаете, что в отчётах могут быть два разных смысла: иногда нужна целая часть (для грубой статистики), а иногда нужна дробь (для точного анализа). И самое важное — преобразование в float64 нужно сделать до деления, иначе дробная часть потеряется.
Программа читает из stdin два целых числа totalUnits и partsCount (оба int, причём partsCount != 0). Затем вычисляет integerQuotient = totalUnits / partsCount и floatQuotient = float64(totalUnits) / float64(partsCount). Печатает два результата в двух строках: сначала целочисленный, затем дробный.
Требования:
• Программа должна считать из stdin два значения через fmt.Scan и сохранить их в переменные totalUnits и partsCount типа int (гарантируется, что partsCount != 0).
• Программа должна вычислить integerQuotient как результат целочисленного деления totalUnits / partsCount без каких-либо преобразований типов.
• Программа должна вычислить floatQuotient как float64(totalUnits) / float64(partsCount), при этом оба операнда должны быть приведены к float64 до выполнения операции деления.
• Программа должна вывести два результата в stdout в двух отдельных строках: сначала integerQuotient, затем floatQuotient.
• Программа должна использовать только пакет fmt и не импортировать другие пакеты.
*/

import "fmt"

func main() {
	var totalUnits, partsCount int
	fmt.Scan(&totalUnits, &partsCount)

	integerQuotient := totalUnits / partsCount

	// TODO: Исправьте вычисление дробного частного: сейчас дробная часть теряется из-за целочисленного деления.
	floatQuotient := float64(totalUnits) / float64(partsCount)

	fmt.Println(integerQuotient)
	fmt.Println(floatQuotient)
}