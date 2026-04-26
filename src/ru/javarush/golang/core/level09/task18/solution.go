package main

/*
= Задача №18. Выбор эффекта =
Вы настраиваете «магический модификатор» в текстовой игре: игрок вводит число силы и короткий код эффекта. По коду нужно выбрать правило преобразования числа и применить его. Но коды могут быть неправильными, и тогда важно не пытаться вызвать nil-функцию — нужно аккуратно вернуть ошибку.
Программа читает целое число powerValue и строку ruleCode. Опишите именованный тип функции: type Rule func(int) int. Реализуйте chooseRule(code string) (Rule, error), которая возвращает правило: "d" — удвоить, "s" — возвести в квадрат, "n" — оставить как есть. Для неизвестного кода функция должна вернуть nil и ошибку.
После выбора правила программа применяет его к powerValue через отдельную функцию apply(x int, r Rule) int и печатает result: <число>. Если код неизвестен — печатает error: unknown rule.
Требования:
• Нужно объявить именованный тип функции: type Rule func(int) int.
• Нужно реализовать chooseRule(code string) (Rule, error), которая для кодов "d"/"s"/"n" возвращает соответствующее правило, а для неизвестного кода возвращает (nil, error).
• Для кодов "d" (удвоить), "s" (возвести в квадрат), "n" (оставить как есть) должны использоваться анонимные функции (func(int) int).
• Нужно реализовать apply(x int, r Rule) int, которая применяет правило r к числу x и возвращает результат.
• После вызова chooseRule ошибка должна проверяться сразу; при ошибке нельзя пытаться вызывать правило (nil), вместо этого нужно вывести сообщение об ошибке.
*/

import (
	"errors"
	"fmt"
)

// Rule — правило преобразования силы.
type Rule func(int) int

func chooseRule(code string) (Rule, error) {
	switch code {
	case "n":
		return func(x int) int { return x }, nil
	// TODO: добавьте обработку кода "d" (удвоить) через анонимную функцию func(int) int
	case "d":
        return func(x int) int {return 2*x}, nil
	// TODO: добавьте обработку кода "s" (возвести в квадрат) через анонимную функцию func(int) int
	case "s":
	    return func(x int) int {return x*x}, nil
	default:
		// TODO: для неизвестного кода нужно вернуть nil и ошибку
		return nil, errors.New("unknown rule")
	}
}

func apply(x int, r Rule) int {
	return r(x)
}

func main() {
	var powerValue int
	var ruleCode string
	fmt.Scan(&powerValue, &ruleCode)

	r, err := chooseRule(ruleCode)
	if err != nil {
		fmt.Print("error: unknown rule")
		return
	}

	result := apply(powerValue, r)
	fmt.Printf("result: %d", result)
}