package main

/*
= Задача №10. Вход на концерт =
На входе в концертный зал два независимых признака: у человека может быть билет, а может быть VIP-пропуск. Охране не нужны ваши рассуждения — ей нужен чёткий флаг “пускаем/не пускаем”.
Программа читает две строки: ticketStr и vipStr. Каждая строка равна либо "yes", либо "no".
Далее программа превращает их в булевы значения: hasTicket := ticketStr == "yes" isVIP := vipStr == "yes"
После этого вычисляет canEnter по правилу: вход разрешён, если есть билет ИЛИ есть VIP (||). Если canEnter истинно — вывести ENTER, иначе — DENY.
Требования:
• Программа должна прочитать ровно две строки (ticketStr и vipStr) с использованием fmt.Scan.
• Программа должна вычислить hasTicket как (ticketStr == "yes") и isVIP как (vipStr == "yes").
• Программа должна вычислить canEnter с использованием оператора || по правилу: canEnter истинно, если hasTicket || isVIP.
• Программа должна вывести ровно одно слово: ENTER если canEnter == true, иначе DENY.
*/

import "fmt"

func main() {
	// Минимальная структура Go-программы с main и fmt

	var ticketStr, vipStr string
	fmt.Scan(&ticketStr, &vipStr) // читаем ровно две строки

	// Превращаем строки в булевы флаги строго по условию
	hasTicket := ticketStr == "yes"
	isVIP := vipStr == "yes"

	// TODO: Вычислите canEnter по правилу задачи (используйте логическое ИЛИ).
	// Сейчас логика намеренно неверная.
	canEnter := hasTicket || isVIP

	if canEnter {
		// TODO: Выведите корректный результат для случая, когда canEnter == true.
		fmt.Print("ENTER")
	} else {
		// TODO: Выведите корректный результат для случая, когда canEnter == false.
		fmt.Print("DENY")
	}
}