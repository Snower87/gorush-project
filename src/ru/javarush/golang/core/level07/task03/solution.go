package main

/*
= Задача №3. Баланс кошелька =
Вы делаете мини‑учёт для кошелька в приложении: всё хранится в центах, чтобы не связываться с float.
Пользователь приносит два числа: сколько денег пришло и сколько ушло — оба в центах.
Программа читает incomeCents и expenseCents (оба int64), считает balanceCents = incomeCents - expenseCents, а затем печатает “квитанцию” из двух строк: сначала баланс с типом, затем статус — прибыль, убыток или ноль.
Требования:
• Программа должна считать два целых числа через fmt.Scan в переменные incomeCents и expenseCents, обе переменные должны иметь тип int64.
• Программа должна вычислить balanceCents = incomeCents - expenseCents, при этом в вычислениях должны использоваться только значения и переменные типа int64.
• Переменная balanceCents должна быть объявлена/создана так, чтобы её тип был именно int64 (а не int или другой целочисленный тип).
• Первая строка вывода должна печататься через fmt.Printf и должна содержать информацию о балансе и его типе, используя плейсхолдер %T.
*/

import "fmt"

func main() {
	// Все суммы храним в центах, поэтому используем int64.
	var incomeCents int64
	var expenseCents int64
	fmt.Scan(&incomeCents, &expenseCents)

	// Баланс тоже строго int64 (никаких int/float).
	var balanceCents int64
	// TODO: вычислите balanceCents как incomeCents - expenseCents (все переменные int64)
	balanceCents = incomeCents - expenseCents

	// Первая строка "квитанции": значение баланса + его тип через %T.
	// TODO: выведите тип именно balanceCents через %T (не используйте литералы)
	fmt.Printf("balanceCents=%d type=%T\n", balanceCents, balanceCents)

	// Вторая строка: статус по знаку баланса (строго одно слово).
	//fmt.Println("status")
	// TODO: определите статус по balanceCents через if/else if/else и выведите profit/loss/zero
	if balanceCents > 0 {
	    fmt.Println("profit")
	} else if balanceCents < 0 {
	    fmt.Println("loss")
	} else if balanceCents == 0 {
	    fmt.Println("zero")
	}
}