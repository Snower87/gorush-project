package main

/*
= Задача №8. Статус задачи =
Вы делаете простую консольную панель для трекера задач. Пользователь вводит: title (одно слово), planned (сколько дней планировали, int), spent (сколько реально потратили часов или дней, float64) и done (готово ли, bool).
Дальше нужно вычислить:
- isBig: planned >= 5
- isLong: spent >= 3.0
- needsAttention: (!done) && (isBig || isLong)
И вывести:
1. строку статуса: либо <title>: DONE, либо <title>: IN PROGRESS
2. строку isBig=<значение>
3. строку needsAttention=<значение>
Требования:
• Программа должна считать из stdin четыре значения в порядке: title (string, одно слово), planned (int), spent (float64), done (bool) с использованием чтения до пробела для title (fmt.Scan).
• Программа должна вычислить isBig типа bool как результат int-сравнения planned >= 5.
• Программа должна вычислить isLong как результат float64-сравнения spent >= 3.0 и использовать его только для расчёта needsAttention.
• Программа должна вычислить needsAttention типа bool по формуле (!done) && (isBig || isLong), используя операции !, &&, || и круглые скобки.
• Программа должна сформировать строку статуса конкатенацией через + в формате либо "<title>: DONE", либо "<title>: IN PROGRESS" в зависимости от значения done.
*/

import "fmt"

func main() {
	var title string
	var planned int
	var spent float64
	var done bool

	// Ввод строго в указанном порядке, title читаем до пробела.
	fmt.Scan(&title, &planned, &spent, &done)

	// TODO: Вычислите isBig по условию задачи на основе planned (сравнение int).
	isBig := planned >= 5

	// TODO: Вычислите isLong по условию задачи на основе spent (сравнение float64).
	// Важно: используйте isLong только при расчёте needsAttention.
	isLong := spent >= 3.0

	// TODO: Вычислите needsAttention по условию задачи, используя !, &&, || и скобки.
	// Подсказка по форме: needsAttention должен зависеть от done, isBig и isLong.
	needsAttention := (!done) && (isBig || isLong)

	// TODO: Сформируйте строку статуса через конкатенацию (+) в нужном формате
	// и в зависимости от done.
	status := title + ": IN PROGRESS"
	if done {
		// TODO: Исправьте статус для случая done == true.
		status = title + ": DONE"
	}

	// Ровно три строки и в точном порядке.
	fmt.Println(status)
	fmt.Println("isBig=" + fmt.Sprint(isBig))
	fmt.Println("needsAttention=" + fmt.Sprint(needsAttention))
}