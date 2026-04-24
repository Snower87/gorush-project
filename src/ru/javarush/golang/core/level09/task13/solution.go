package main

/*
= Задача №13. Базовый счёт =
Вы делаете систему начисления очков в игре: есть базовые очки за уровень, и к ним могут добавиться бонусы — иногда ни одного, иногда один, иногда сразу два (например, "за скорость" и "за безошибочность"). Вам нужно, чтобы функция суммирования принимала базу и любое количество бонусов, но в main вы строго передавали 0/1/2 бонуса именно как отдельные аргументы.
Напишите функцию totalWithBonuses(base int, bonuses ...int) int, которая возвращает сумму base и всех бонусов. Программа читает base и bonusCount, после чего, в зависимости от bonusCount (0, 1 или 2), читает нужное количество бонусов и вызывает totalWithBonuses подходящим образом. Результат выводится одним целым числом.
Требования:
• В программе должна быть реализована функция totalWithBonuses(base int, bonuses ...int) int, возвращающая сумму base и всех переданных бонусов.
• Параметр bonuses ...int должен быть последним параметром функции totalWithBonuses.
• Внутри totalWithBonuses должна быть обработка всех переданных бонусов через цикл (например, for/range) с добавлением каждого бонуса к итоговой сумме.
• Программа должна считать из stdin два целых числа base и bonusCount, а затем (в зависимости от bonusCount = 0/1/2) считать 0, 1 или 2 дополнительных целых числа бонусов.
• При bonusCount == 0 нужно вызвать totalWithBonuses(base) без бонусов; при bonusCount == 1 — totalWithBonuses(base, b1); при bonusCount == 2 — totalWithBonuses(base, b1, b2).
*/

import (
	"fmt"
	"os"
)

// totalWithBonuses суммирует базу и любое количество бонусов.
// Важно: variadic-параметр (...int) должен быть последним.
func totalWithBonuses(base int, bonuses ...int) int {
	// TODO: Реализуйте суммирование base и всех бонусов.
	// TODO: Используйте цикл по bonuses (например, for/range) и прибавляйте каждый бонус к итоговой сумме.
	for _, bonus := range bonuses {
	    base += bonus
	}
	return base
}

func main() {
	var base, bonusCount int
	fmt.Fscan(os.Stdin, &base, &bonusCount)

	var result int
	switch bonusCount {
	case 0:
		// TODO: Вызовите totalWithBonuses(base) без бонусов и сохраните результат в result.
		result = totalWithBonuses(base)
	case 1:
		var b1 int
		fmt.Fscan(os.Stdin, &b1)
		// TODO: Вызовите totalWithBonuses(base, b1) (бонус отдельным аргументом) и сохраните результат в result.
		result = totalWithBonuses(base, b1)
	case 2:
		var b1, b2 int
		fmt.Fscan(os.Stdin, &b1, &b2)
		// TODO: Вызовите totalWithBonuses(base, b1, b2) (оба бонуса отдельными аргументами) и сохраните результат в result.
		result = totalWithBonuses(base, b1, b2)
	}

	fmt.Println(result)
}