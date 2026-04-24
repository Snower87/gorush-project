package main

/*
= Задача №15. Честное среднее =
Вы анализируете результаты теста: иногда данных вообще нет (никто не прошёл тест), и тогда среднее считать нельзя — нужно вернуть ошибку. В остальных случаях среднее должно быть вещественным (с дробной частью), и в выводе вы хотите видеть ровно два знака после запятой, как в нормальном отчёте.
Напишите функцию average(nums ...int) (float64, error), которая считает среднее арифметическое. Если len(nums) == 0, функция должна вернуть ошибку. Программа читает count (0..4), затем читает ровно count целых чисел и вызывает average с соответствующим количеством аргументов (0..4). При ошибке печатайте "error: <текст ошибки>", при успехе — среднее с двумя знаками после запятой.
Требования:
• В программе должна быть реализована функция average(nums ...int) (float64, error), принимающая переменное число целых аргументов и возвращающая среднее (float64) и ошибку (error).
• Если len(nums) == 0, функция average должна выполнить проверку и сразу (ранним return) вернуть ненулевую ошибку (err != nil).
• Среднее арифметическое должно вычисляться как float64(sum)/float64(count), где sum — сумма всех чисел, count — количество чисел (len(nums)); нельзя допускать целочисленного деления.
• Программа должна прочитать из stdin число count (в диапазоне 0..4), затем прочитать ровно count целых чисел.
• Программа должна корректно вызвать average с соответствующим количеством аргументов (0..4), используя ветвление по значению count.
*/

import (
	"errors"
	"fmt"
)

func average(nums ...int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("no numbers")
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	// TODO: Посчитайте среднее арифметическое как float64 (без целочисленного деления).
	return float64(sum) / float64(len(nums)), nil
}

func main() {
	var count int
	fmt.Scan(&count)

	// читаем ровно count чисел (0..4)
	var a, b, c, d int
	switch count {
	case 0:
		// ничего не читаем
	case 1:
		fmt.Scan(&a)
	case 2:
		fmt.Scan(&a, &b)
	case 3:
		fmt.Scan(&a, &b, &c)
	case 4:
		fmt.Scan(&a, &b, &c, &d)
	}

	// вызываем average через ветвление по count (0..4)
	var (
		avg float64
		err error
	)
	switch count {
	case 0:
		avg, err = average()
	case 1:
		avg, err = average(a)
	case 2:
        avg, err = average(a, b)
	case 3:
        avg, err = average(a, b, c)
	case 4:
        avg, err = average(a, b, c, d)
	}

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Printf("%.2f", avg)
}