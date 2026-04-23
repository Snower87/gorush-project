package main

/*
= Задача №8. Одна операция =
Вы делаете «карманный калькулятор» для админки: на вход приходит выражение из трёх токенов — два числа и оператор. Это должно работать безопасно: если пользователь ввёл не число, неизвестную операцию или попытался поделить на ноль — вы не падаете, а печатаете понятную ошибку и завершаетесь.
Программа читает три токена через fmt.Scan: aStr, operator, bStr. Парсинг и вычисление разнесены по функциям: строгий парсер числа, безопасное деление и вычисление одной операции. При успехе выводите result = <число>, при любой ошибке — error: <текст ошибки> и сразу выходите.
Требования:
• Программа должна быть в package main и импортировать ровно (как минимум) пакеты fmt, strconv и errors.
• В main программа должна прочитать ровно три значения через fmt.Scan в переменные aStr, op (operator) и bStr (в этом порядке), без использования чтения из stdin другими способами.
• Должна быть реализована функция parseIntStrict, которая принимает строку, парсит её в int через strconv и возвращает (value, nil) при успехе, либо (0, error) при любой ошибке парсинга.
• Должна быть реализована функция safeDivInt, которая выполняет целочисленное деление a/b и при b == 0 возвращает (0, errors.New("division by zero")), иначе возвращает (result, nil).
• Должна быть реализована функция calcOneOp с контрактом (value, error), которая выполняет одну операцию над числами a и b по оператору op; при неизвестном операторе она должна возвращать (0, error), а для деления должна использовать safeDivInt (а не делить напрямую).
• Каждая из функций parseIntStrict, safeDivInt и calcOneOp при любой ошибке должна возвращать zero value результата и не-nil error (ошибки не должны “прятаться” через bool/спец-значения).
*/

import (
	"errors"
	"fmt"
	"strconv"
)

func parseIntStrict(s string) (int, error) {
	// TODO: Реализуйте строгий парсинг строки в int через strconv.
	// TODO: При любой ошибке парсинга возвращайте (0, non-nil error) с согласованным текстом ошибки.

	v, err := strconv.Atoi(s)
	if err != nil {
		// В шаблоне намеренно используется другой текст ошибки, чтобы решение не было готовым.
		return 0, errors.New("invalid integer")
	}
	return v, nil
}

func safeDivInt(a, b int) (int, error) {
	// TODO: Реализуйте безопасное целочисленное деление a/b.
	// TODO: При b == 0 возвращайте (0, errors.New(...)) с согласованным текстом ошибки.

	if b == 0 {
		// В шаблоне намеренно используется другой текст ошибки, чтобы решение не было готовым.
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func calcOneOp(a int, op string, b int) (int, error) {
	// TODO: Реализуйте вычисление одной операции по оператору op.
	// TODO: Поддержите +, -, *, /.
	// TODO: Для деления используйте safeDivInt (не делите напрямую).
	// TODO: Для неизвестного оператора возвращайте (0, non-nil error) с согласованным текстом ошибки.

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return safeDivInt(a, b)
	default:
		// В шаблоне умышленно не реализованы * и /.
		return 0, errors.New("unknown operator")
	}
}

func main() {
	var aStr, op, bStr string
	fmt.Scan(&aStr, &op, &bStr)

	a, err := parseIntStrict(aStr)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	b, err := parseIntStrict(bStr)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	res, err := calcOneOp(a, op, b)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("result = %d\n", res)
}