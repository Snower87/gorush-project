package main

/*
= Задача №20. Сборный калькулятор =
Вы собираете «командный калькулятор» для консоли: на вход приходит четыре токена — два числа, оператор и код правила пост-обработки результата. Важно сделать код чистым и модульным: отдельные маленькие функции для парсинга, выбора операции, выбора правила, применения правила и единый сценарий run(), который либо печатает результат, либо возвращает ошибку.
Программа читает четыре токена: a, op, b, ruleCode, где a и b — целые числа, op — один из + - * /, ruleCode — один из d/s/n.
Разложите программу на функции:
- parseInt(s string) (int, error)
- chooseOp(op string) (func(int, int) (int, error), error)
- chooseRule(code string) (func(int) int, error)
- apply(x int, r func(int) int) int
- run() error
При успехе печатайте result: <число>. При делении на ноль печатайте error: division by zero. При неизвестном операторе печатайте error: unknown operator. При неизвестном правиле печатайте error: unknown rule.
Требования:
• Основная логика должна быть вынесена в функцию run() error; функция main должна только вызвать run() и, если вернулась ошибка, напечатать сообщение об ошибке.
• Программа должна прочитать из stdin ровно четыре токена, в порядке: a, op, b, ruleCode, где a и b читаются как строки для последующего парсинга, а op и ruleCode — как строки-коды.
• Должна быть реализована функция parseInt(s string) (int, error), которая преобразует строку в int через strconv.Atoi и возвращает ошибку при некорректном числе.
• Должна быть реализована функция chooseOp(op string) (func(int, int) (int, error), error), которая по строке-оператору возвращает функцию для вычисления +, -, * или /; при неизвестном операторе должна возвращаться ошибка, приводящая к выводу ровно "error: unknown operator".
• Операция деления должна быть реализована так, чтобы при b == 0 возвращалась ошибка, приводящая к выводу ровно "error: division by zero" (а не panic и не другое сообщение).
• Должна быть реализована функция chooseRule(code string) (func(int) int, error), которая возвращает короткую (анонимную) функцию для кодов d/s/n; при неизвестном коде правила должна возвращаться ошибка, приводящая к выводу ровно "error: unknown rule".
• Должна быть реализована функция apply(x int, r func(int) int) int, которая применяет переданную rule-функцию к значению x и возвращает результат.
*/

import (
	"errors"
	"fmt"
	"strconv"
)

// Sentinel-ошибки: текст должен совпадать с требованиями.
var (
	errDivisionByZero = errors.New("error: division by zero")
	errUnknownOp      = errors.New("error: unknown operator")
	errUnknownRule    = errors.New("error: unknown rule")
)

func main() {
	if err := run(); err != nil {
		fmt.Print(err.Error())
	}
}

func run() error {
	var aS, opS, bS, ruleCode string
	_, err := fmt.Scan(&aS, &opS, &bS, &ruleCode)
	if err != nil {
		return err
	}

	a, err := parseInt(aS)
	if err != nil {
		return err
	}

	b, err := parseInt(bS)
	if err != nil {
		return err
	}

	op, err := chooseOp(opS)
	if err != nil {
		return err
	}

	res, err := op(a, b)
	if err != nil {
		return err
	}

	rule, err := chooseRule(ruleCode)
	if err != nil {
		return err
	}

	res = apply(res, rule)
	fmt.Printf("result: %d", res)
	return nil
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func chooseOp(op string) (func(int, int) (int, error), error) {
	switch op {
	case "+":
		return func(a, b int) (int, error) { return a + b, nil }, nil
	case "-":
		// TODO: Реализуйте операцию вычитания и верните функцию, которая вычисляет a - b.
		return func(a, b int) (int, error) { return a - b, nil}, nil
	case "*":
		// TODO: Реализуйте операцию умножения и верните функцию, которая вычисляет a * b.
		return func(a, b int) (int, error) { return a * b, nil}, nil
	case "/":
		// TODO: Реализуйте операцию деления и обработку деления на ноль (вернуть errDivisionByZero).
		return func(a, b int) (int, error) {
            if b == 0 {
                return 0, errDivisionByZero
            }
            return a / b, nil
		}, nil
	default:
		return nil, errUnknownOp
	}
}

func chooseRule(code string) (func(int) int, error) {
	switch code {
	case "d":
		// TODO: Реализуйте правило "d" и верните функцию пост-обработки результата.
		return func(x int) int {return 2*x}, nil
	case "s":
		// TODO: Реализуйте правило "s" и верните функцию пост-обработки результата.
		return func(x int) int {return x*x}, nil
	case "n":
		return func(x int) int { return x }, nil
	default:
		return nil, errUnknownRule
	}
}

func apply(x int, r func(int) int) int {
	return r(x)
}