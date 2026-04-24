package main

/*
= Задача №16. Статистика партии =
Вы ведёте журнал партии в настольной игре: по имени игрока и списку его бросков нужно сделать короткую статистику — минимум, максимум и сколько значений пришло. Если значений нет, это ошибка: нельзя говорить про min/max «из воздуха».
При этом вы хотите, чтобы форматирование отчёта было отдельной функцией и использовало minMax через прокидывание nums....
Реализуйте minMax(nums ...int) (min int, max int, err error): для пустого набора возвращайте ошибку, для непустого — минимальное и максимальное. Реализуйте formatStats(name string, nums ...int) (string, error), которая возвращает строку формата "<name>: min=<min> max=<max> count=<n>" и внутри вызывает minMax(nums...).
Программа читает name и count (0..5), затем читает count целых чисел и вызывает formatStats с 0..5 аргументами. Если ошибка — печатайте "error: <текст>", иначе печатайте строку статистики.
Требования:
• Должна быть реализована функция minMax(nums ...int) (min int, max int, err error), принимающая произвольное число целых чисел.
• Если len(nums) == 0, функция minMax должна возвращать ненулевую ошибку (min/max при этом не должны вычисляться «из воздуха»).
• Для непустого набора minMax должна находить минимум и максимум одним проходом по nums (один цикл без дополнительных проходов/сортировки).
• Должна быть реализована функция formatStats(name string, nums ...int) (string, error), которая формирует строку отчёта и внутри вызывает minMax с прокидыванием аргументов именно как minMax(nums...) (с тремя точками).
• Программа должна прочитать name и count (0..5), затем прочитать ровно count целых чисел, и вызвать formatStats так, чтобы count определял количество переданных variadic-аргументов через ветвление по count (т.е. вызов с 0, 1, 2, 3, 4 или 5 числами).
• При успешном вычислении программа должна вывести строку строго формата "<name>: min=<min> max=<max> count=<n>", где n — фактическое количество значений (count).
*/

import (
	"errors"
	"fmt"
	"os"
)

func minMax(nums ...int) (min int, max int, err error) {
	if len(nums) == 0 {
		// Нельзя вычислить min/max "из воздуха".
		return 0, 0, errors.New("no numbers")
	}

	min, max = nums[0], nums[0]
	// TODO: Реализуйте поиск минимума и максимума за один проход по nums (без сортировки и без дополнительных проходов)
	for _, n := range nums {
	    if n > max {
	        max = n
	    }
	    if n < min {
	        min = n
	    }
	}

	return min, max, nil
}

func formatStats(name string, nums ...int) (string, error) {
	// Важно: прокидываем variadic как minMax(nums...) (с тремя точками).
	min, max, err := minMax(nums...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s: min=%d max=%d count=%d", name, min, max, len(nums)), nil
}

func main() {
	in := os.Stdin

	var name string
	var count int
	fmt.Fscan(in, &name, &count)

	var a1, a2, a3, a4, a5 int
	if count >= 1 {
		fmt.Fscan(in, &a1)
	}
	if count >= 2 {
		fmt.Fscan(in, &a2)
	}
	if count >= 3 {
		fmt.Fscan(in, &a3)
	}
	if count >= 4 {
		fmt.Fscan(in, &a4)
	}
	if count >= 5 {
		fmt.Fscan(in, &a5)
	}

	// По требованию: вызываем formatStats с 0..5 аргументами через ветвление.
	var s string
	var err error
	switch count {
	case 0:
		s, err = formatStats(name)
	case 1:
		s, err = formatStats(name, a1)
	case 2:
		s, err = formatStats(name, a1, a2)
	case 3:
		s, err = formatStats(name, a1, a2, a3)
	case 4:
		s, err = formatStats(name, a1, a2, a3, a4)
	default: // 5
		s, err = formatStats(name, a1, a2, a3, a4, a5)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	fmt.Println(s)
}