package main

/*
= Задача №4. Касса расходов =
Вы делаете консольную “кассу расходов” для похода с друзьями. Каждый вносит траты: категория (например, "food") и сумма (например, "120"). Но ввод бывает грязным: кто-то может вместо суммы написать "one hundred". В таком случае касса должна остановиться на первой же ошибке и честно сказать, какая ошибка произошла — иначе итог станет недостоверным.
Сначала программа читает количество записей n. Затем читает n пар: expenseCategory (строка) и amountToken (строка). n и каждая сумма должны парситься через strconv.Atoi, чтобы ошибки были обычными значениями error. Если что-то пошло не так (ошибка чтения или парсинга), программа печатает "FAILED" и значение ошибки, после чего немедленно завершает работу. Если всё корректно — печатает общую сумму.
Требования:
• Программа должна считать первый токен ввода как строку через fmt.Scan (или fmt.Fscan) и преобразовать его в число через strconv.Atoi, получая ошибку как значение error.
• После успешного парсинга n программа должна выполнить цикл ровно на n итераций и на каждой итерации считать пару строк: expenseCategory и amountToken (оба через fmt.Scan/fmt.Fscan).
• На каждой итерации программа должна преобразовать amountToken в int через strconv.Atoi и прибавить полученную сумму к общему total (категория не влияет на расчёт).
• При любой ошибке чтения (fmt.Scan/fmt.Fscan) или парсинга (strconv.Atoi для n или amountToken) программа должна вывести строку, начинающуюся с FAILED, затем вывести значение ошибки, и немедленно выйти из main через return, не продолжая чтение и расчёты.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// n читаем именно как строковый токен.
	var nToken string
	if _, err := fmt.Fscan(in, &nToken); err != nil {
		fmt.Printf("FAILED %v\n", err)
		return
	}

	// TODO: Преобразуйте nToken в число через strconv.Atoi и сохраните получившийся n.
	// TODO: При ошибке парсинга выведите "FAILED <err>" и немедленно завершите main через return.
	//n := strconv.Atoi(nToken)
	var n int
	var err3 error
	n, err3 = strconv.Atoi(nToken)
	if err3 != nil {
		fmt.Printf("FAILED %v\n", err3)
		return
	}

	total := 0

	for i := 0; i < n; i++ {
		// Категория не влияет на расчёт, но токен всё равно нужно прочитать.
		var expenseCategory string
		var amountToken string
		if _, err := fmt.Fscan(in, &expenseCategory, &amountToken); err != nil {
			fmt.Printf("FAILED %v\n", err)
			return
		}

		// TODO: Преобразуйте amountToken в int через strconv.Atoi.
		var intAmountToken int
		intAmountToken, err2 := strconv.Atoi(amountToken)
		// TODO: При ошибке парсинга выведите "FAILED <err>" и немедленно завершите main через return.
		if err2 != nil {
			fmt.Printf("FAILED %v\n", err2)
			return
		}
		// TODO: При успехе прибавьте сумму к total.
		total += intAmountToken
	}

	// Используем strconv, чтобы импорт не был лишним в шаблоне.

	fmt.Printf("TOTAL=%d\n", total)
}