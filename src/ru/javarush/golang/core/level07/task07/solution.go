package main

/*
= Задача №7. Среднее по датчикам =
Вы снимаете показания с трёх датчиков (например, температуры или скорости ветра). Для пользователя в интерфейсе нужно показать красивое среднее значение — всего с двумя знаками после точки. Но для отладки вам важно видеть “сырое” значение с хвостом, чтобы понимать, нет ли накопленной погрешности.
Программа читает из stdin три числа firstMeasure, secondMeasure, thirdMeasure (все float64), считает average = (firstMeasure + secondMeasure + thirdMeasure) / 3 и выводит две строки. Обе начинаются с avg=: сначала “красиво” (2 знака), затем диагностически (17 знаков).
Требования:
• Программа должна прочитать три значения через fmt.Scan(&firstMeasure, &secondMeasure, &thirdMeasure).
• Переменные firstMeasure, secondMeasure, thirdMeasure (и average) должны иметь тип float64.
• Значение average должно вычисляться строго как (firstMeasure + secondMeasure + thirdMeasure) / 3.
• Первая строка вывода должна быть напечатана форматом "avg=%.2f\n" с подстановкой average.
• Вторая строка вывода должна быть напечатана форматом "avg=%.17f\n" с подстановкой average.
*/

import "fmt"

func main() {
	var firstMeasure float64
	var secondMeasure float64
	var thirdMeasure float64

	fmt.Scan(&firstMeasure, &secondMeasure, &thirdMeasure)

	// TODO: Вычислите среднее значение трёх измерений по формуле из условия.
	var average float64 = (firstMeasure + secondMeasure + thirdMeasure) / 3

	fmt.Printf("avg=%.2f\n", average)
	fmt.Printf("avg=%.17f\n", average)
}