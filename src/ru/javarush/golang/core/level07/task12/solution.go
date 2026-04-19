package main

/*
= Задача №12. Звонок по минутам =
Вы настраиваете тарификацию звонка для оператора связи. По правилам оплачиваются только полные минуты, а “лишние” секунды уходят в статистику как остаток. Ещё есть фиксированная базовая плата.
Программа читает usedSeconds, pricePerMinuteCents и baseFeeCents (все целые). Считает billableMinutes = usedSeconds / 60, leftoverSeconds = usedSeconds % 60, переменную часть variableCents = billableMinutes * pricePerMinuteCents, итог totalCents = baseFeeCents + variableCents. Затем выводит одной строкой: евро, центы (двумя цифрами) и остаток секунд.
Требования:
• Программа должна прочитать три целых числа: usedSeconds, pricePerMinuteCents и baseFeeCents.
• Программа должна вычислить billableMinutes как usedSeconds / 60 и leftoverSeconds как usedSeconds % 60, используя целочисленные операторы / и %.
• Программа должна вычислить variableCents как billableMinutes * pricePerMinuteCents (без деления) и totalCents как baseFeeCents + variableCents.
• Программа должна получить euros = totalCents / 100 и cents = totalCents % 100, используя целочисленную арифметику.
*/

import "fmt"

func main() {
	var usedSeconds, pricePerMinuteCents, baseFeeCents int
	fmt.Scan(&usedSeconds, &pricePerMinuteCents, &baseFeeCents)

	// TODO: Посчитайте количество полных оплачиваемых минут (billableMinutes) из usedSeconds.
	// TODO: Посчитайте остаток секунд (leftoverSeconds), который не попал в полные минуты.
	billableMinutes := usedSeconds / 60
	leftoverSeconds := usedSeconds % 60

	// TODO: Посчитайте переменную часть стоимости (variableCents) в центах.
	// TODO: Посчитайте итоговую стоимость (totalCents) в центах с учётом базовой платы.
	variableCents := billableMinutes * pricePerMinuteCents
	totalCents := baseFeeCents + variableCents

	// TODO: Разбейте totalCents на евро (euros) и центы (cents) с помощью целочисленной арифметики.
	euros := totalCents / 100
	cents := totalCents % 100

	fmt.Printf("%d %02d %d", euros, cents, leftoverSeconds)
}