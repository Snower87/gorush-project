package main

/*
= Задача №15. Скидка магазина =
Вы собираете финальную цену товара в онлайн-магазине. Пользователь вводит базовую цену, промокод и флаг "студент/не студент". Важно: любые странные данные (не тот промокод, цена 0, флаг 2) должны отсекаться сразу, чтобы расчёт скидок не пытался работать с мусором.
Программа читает price (целое число), promo (строка) и isStudent (целое число 0 или 1). promo может быть SAVE10, SAVE20 или - (минус означает “нет промокода”).
Если price меньше или равен 0, программа должна вывести invalid price и завершить работу. Если isStudent не равен 0 и не равен 1, программа должна вывести invalid student flag и завершить работу. Если promo не равен SAVE10, SAVE20 или -, программа должна вывести invalid promo и завершить работу.
Скидка по промокоду: SAVE10 даёт 10%, SAVE20 даёт 20%, - даёт 0%. Если isStudent равен 1 и promo не SAVE20, добавляется ещё 5% скидки. Итоговая цена вычисляется целыми числами:
final = price - price*discount/100
Программа должна вывести final.
Требования:
• Программа должна считать из stdin через fmt.Scan три значения: price (int), promo (string) и isStudent (int).
• Если price <= 0 — вывести ровно строку "invalid price" и завершить программу через return; если isStudent не 0 и не 1 — вывести "invalid student flag" и завершить; если promo не "SAVE10", не "SAVE20" и не "-" — вывести "invalid promo" и завершить.
• Программа должна вычислить discount (в процентах) без вложенной “пирамиды” if: базовая скидка зависит только от promo (SAVE10=10, SAVE20=20, "-"=0), затем при isStudent==1 и promo!="SAVE20" добавляется ещё 5%.
• Итоговая цена должна вычисляться как final = price - price*discount/100 с использованием int-арифметики (естественное округление вниз за счёт целочисленного деления).
*/

import "fmt"

func main() {
	var price int
	var promo string
	var isStudent int
	fmt.Scan(&price, &promo, &isStudent)

	// Guard clauses: отсекаем мусор до расчётов.
	if price <= 0 {
		fmt.Print("invalid price")
		return
	}
	if isStudent != 0 && isStudent != 1 {
		fmt.Print("invalid student flag")
		return
	}
	if promo != "SAVE10" && promo != "SAVE20" && promo != "-" {
		fmt.Print("invalid promo")
		return
	}

	discount := 0

	// Базовая скидка зависит только от промокода.
	switch promo {
	case "SAVE10":
		// TODO: Установите базовую скидку для промокода SAVE10.
		discount = 10
	case "SAVE20":
		// TODO: Установите базовую скидку для промокода SAVE20.
		discount = 20
	case "-":
		// TODO: Установите базовую скидку при отсутствии промокода.
		discount = 0
	default:
		fmt.Print("invalid promo")
		return
	}

	// Студентская скидка добавляется, но не поверх SAVE20.
	if isStudent == 1 && promo != "SAVE20" {
		// TODO: Добавьте дополнительную скидку для студента по условию задачи.
		discount += 5
	}

	final := price - price*discount/100
	fmt.Print(final)
}