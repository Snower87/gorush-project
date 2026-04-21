package main

/*
= Задача №11. Статус посылки =
Складская система получает код статуса посылки. По договорённости 0 — это «ничего не знаем» (unknown), а реальные статусы начинаются с 1.
Вы хотите, чтобы нумерация статусов была аккуратной и расширяемой, поэтому используете iota, но так, чтобы «живые» статусы стартовали с единицы.
Программа читает код статуса, преобразует его в OrderStatus и печатает человекочитаемый текст. Любой непонятный код — это unknown.
Требования:
• Должен быть объявлен тип OrderStatus на базе int и набор констант этого типа для статусов created, paid, shipped и unknown.
• Константы OrderStatus должны быть объявлены через iota без ручной нумерации, при этом код 0 должен соответствовать неизвестному статусу (допускается пропуск первого значения через _).
• Программа должна прочитать одно целое число из stdin с помощью fmt.Scan.
• Прочитанное число должно быть преобразовано в OrderStatus, а человекочитаемый текст выбран через switch по значению OrderStatus с обязательной веткой default для unknown/непонятных кодов.
• Программа должна напечатать ровно одну строку (created, paid, shipped или unknown) и ничего лишнего.
*/

import "fmt"

type OrderStatus int

const (
	// TODO: Объявите константы OrderStatus через iota так, чтобы unknown соответствовал коду 0,
	// TODO: а "живые" статусы (created, paid, shipped) начинались с 1 (без ручной нумерации).
	unknown OrderStatus = iota
	created
	paid
	shipped
)

func main() {
	var code int
	fmt.Scan(&code)

	status := OrderStatus(code)

	// TODO: Выберите человекочитаемый текст через switch по status.
	// TODO: Нельзя сравнивать code/status с числовыми литералами 1/2/3 — используйте только enum-константы.
	// TODO: Для всех непонятных кодов (включая 0 и любые другие) выводите "unknown" через ветку default.
	switch status {
	case created:
        fmt.Println("created")
	case paid:
        fmt.Println("paid")
	case shipped:
        fmt.Println("shipped")
	default:
		fmt.Println("unknown")
	}
}