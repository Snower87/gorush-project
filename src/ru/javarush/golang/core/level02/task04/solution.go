package main

import "fmt"

func main() {
	// TODO: Задайте totalMinutes (тип int) прямо в коде так, чтобы итоговый вывод был ровно: "2 10 20".
	// TODO: Важно: не используйте ввод из stdin.
	var totalMinutes int = 3260

	// Длительность дня в минутах по требованию: 24*60
	const minutesInDay = 24 * 60

	// TODO: Вычислите количество полных дней в totalMinutes, используя только целочисленное деление (/).
	days := totalMinutes / minutesInDay

	// TODO: Вычислите остаток минут после вычитания полных дней, используя оператор %.
	restAfterDays := totalMinutes % minutesInDay

	// TODO: Вычислите количество полных часов из остатка минут (restAfterDays), используя /.
	hours := restAfterDays / 60

	// TODO: Вычислите остаток минут после вычитания часов, используя %.
	minutes := restAfterDays % 60

	fmt.Printf("%d %d %d", days, hours, minutes)
}