package main

/*
= Задача №16. Копим до цели =
Вы делаете мини-симуляцию накоплений: есть стартовый баланс, есть цель, и есть фиксированное пополнение каждый "месяц". Менеджер хочет понять, за сколько месяцев накопления достигнут цели, но при этом важно не допустить бесконечный цикл, если пополнение неправильное.
Вводятся три целых числа: startBalance, targetBalance и monthlyAdd. Каждый месяц баланс увеличивается на monthlyAdd. Нужно вывести, через сколько месяцев баланс станет больше либо равен цели. Если старт уже подходит, выводится 0. Если startBalance < targetBalance и при этом monthlyAdd <= 0, цель недостижима — нужно вывести -1, обработав это до запуска цикла.
Требования:
• Программа должна считать из stdin три целых числа: startBalance, targetBalance и monthlyAdd.
• Если startBalance < targetBalance и monthlyAdd <= 0, программа должна вывести -1 и завершиться, не входя в цикл накопления.
• Если startBalance >= targetBalance (включая startBalance == targetBalance и targetBalance меньше startBalance), программа должна вывести 0 без выполнения цикла.
• Если требуется накопление (startBalance < targetBalance и monthlyAdd > 0), программа должна использовать цикл вида for condition { ... }, где условие контролирует завершение (баланс < цели).
• На каждой итерации цикла программа должна увеличивать баланс на monthlyAdd и увеличивать счётчик месяцев на 1, чтобы гарантировать приближение к выполнению условия завершения.
*/

import "fmt"

func main() {
	var startBalance, targetBalance, monthlyAdd int
	fmt.Scan(&startBalance, &targetBalance, &monthlyAdd)

	// Уже достигли цели — цикл не нужен.
	if startBalance >= targetBalance {
		fmt.Println(0)
		return
	}

	// Цель недостижима — важно обработать ДО входа в цикл.
	if monthlyAdd <= 0 {
		fmt.Println(-1)
		return
	}

	months := 0
	balance := startBalance

	// Цикл в стиле while: завершается только по условию (без break).
	for balance < targetBalance {
		// TODO: На каждой итерации увеличивайте balance на monthlyAdd (а не на 1).
		// TODO: На каждой итерации увеличивайте months на 1.
		balance += monthlyAdd
		months++
	}

	fmt.Println(months)
}