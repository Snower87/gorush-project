package main

/*
= Задача №11. Скидка без дроби =
Вы пишете кассу, которая работает строго в центах — никаких float, чтобы не поймать неожиданные хвосты. У вас есть цена за штуку unitPriceCents, количество quantity и скидка в процентах discountPercent.
Программа должна посчитать промежуточную сумму subtotalCents, скидку discountCents, итог totalCents, а затем вывести результат как “евро центы”, причём центы всегда двумя цифрами (например, 05). Вывод — ровно два числа через пробел: евро и центы.
Требования:
• Программа должна считать из stdin три целых числа: unitPriceCents, quantity и discountPercent.
• Программа должна вычислить subtotalCents = unitPriceCents * quantity и сохранить результат в отдельной промежуточной переменной.
• Программа должна вычислить discountCents = subtotalCents * discountPercent / 100, выполняя умножение до деления и используя только целочисленные операции (без float).
• Программа должна вычислить totalCents = subtotalCents - discountCents, не объединяя все вычисления в одну длинную строку без subtotalCents и discountCents.
• Программа должна вычислить eur = totalCents / 100 и cen = totalCents % 100 с помощью целочисленного деления и остатка.
*/

import "fmt"

func main() {
	var unitPriceCents, quantity, discountPercent int
	fmt.Scan(&unitPriceCents, &quantity, &discountPercent)

	// Все считаем в центах, без float.
	subtotalCents := unitPriceCents * quantity

	// TODO: Посчитайте discountCents (скидку в центах) строго целочисленными операциями.
	// TODO: Важно сохранить результат в отдельной переменной discountCents.
	discountCents := subtotalCents * discountPercent / 100

	// TODO: Посчитайте totalCents через subtotalCents и discountCents (отдельной переменной).
	totalCents := subtotalCents - discountCents

	eur := totalCents / 100
	cen := totalCents % 100

	fmt.Printf("%d %02d", eur, cen)
}