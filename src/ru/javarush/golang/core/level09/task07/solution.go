package main

/*
= Задача №7. Позитивный ввод =
Вы пишете маленькую утилиту для терминала: оператор вводит число — количество коробок в заказе. Ноль или отрицательное значение — это ошибка, и программа должна сразу остановиться, не продолжая работу, чтобы не наделать глупостей дальше по цепочке.
Программа читает один токен-строку rawCount. Реализуйте функцию parsePositiveInt(s string) (int, error): она пытается распарсить int через strconv.Atoi. Если парсинг не удался — верните 0 и исходную ошибку. Если число получилось, но оно <= 0, верните 0 и новую ошибку "number must be positive" (через errors.New). В main после вызова сразу проверяйте err != nil, печатайте "error: <текст>" и делайте ранний return. Если ошибки нет — печатайте "n = <число>".
Требования:
• Код должен импортировать ровно необходимые пакеты: fmt для вывода, strconv для Atoi, errors для создания собственной ошибки.
• В программе должна быть реализована функция parsePositiveInt(s string) (int, error), возвращающая число и ошибку по стандартному контракту Go.
• В parsePositiveInt нужно вызывать strconv.Atoi(s); если парсинг не удался, функция должна вернуть (0, исходную ошибку) без модификаций и без специальных “сентинельных” чисел (например, -1).
• Если число распарсилось, но значение <= 0, parsePositiveInt должна вернуть (0, errors.New("number must be positive")).
• В main после вызова parsePositiveInt должна сразу выполняться проверка if err != nil { ...; return }, без продолжения выполнения программы при ошибке.
*/

import (
	"errors"
	"fmt"
	"strconv"
)

func parsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if n <= 0 {
		// TODO: Верните ошибку с корректным текстом (через errors.New), как требуется в условии.
		return 0, errors.New("number must be positive")
	}

	return n, nil
}

func main() {
	var rawCount string
	fmt.Scan(&rawCount)

	n, err := parsePositiveInt(rawCount)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		// TODO: При ошибке нужно сразу остановить выполнение (ранний выход), не печатая ничего больше.
		return
	}

	fmt.Printf("n = %d\n", n)
}