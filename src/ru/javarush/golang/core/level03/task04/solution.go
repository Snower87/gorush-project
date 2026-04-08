package main

/*
= Задача №4. Длина логина =
Вы делаете экран регистрации в консольном прототипе. Пользователь вводит логин, а правила такие: логин не должен быть
слишком коротким и не должен быть слишком длинным. Ещё вам важно отдельно знать, совпал ли он ровно с минимально
допустимой длиной (например, для подсказок UI).
Программа читает строку login, затем два целых числа minLen и maxLen.
Она считает длину логина в байтах: loginLen := len(login). После этого выводит на одной строке три булевых значения через пробел:
isTooShort — loginLen меньше minLen
isTooLong — loginLen больше maxLen
isExactMin — loginLen равно minLen
Ввод: login, затем minLen и maxLen (в любом разбиении по пробелам/переводам строк).
Вывод: isTooShort isTooLong isExactMin.
Пример ввода:
gopher 3 10
Пример вывода:
false false false
Требования:
• Программа должна считать из stdin строку login и два целых числа minLen и maxLen с помощью fmt.Scan (разделители могут быть пробелы и/или переводы строк).
• Программа должна вычислить длину login как loginLen := len(login) и сохранить результат в отдельную переменную типа int.
• Программа должна вычислить и сохранить в три bool-переменные с понятными именами значения: - isTooShort: loginLen < minLen - isTooLong: loginLen > maxLen - isExactMin: loginLen == minLen При этом для сравнений должны использоваться только операторы ==, <, > (и при необходимости <=, >=), без логических операторов &&, ||, !.
*/

import "fmt"

func main() {
	var login string
	var minLen, maxLen int

	// Ввод: строка логина и две границы длины.
	fmt.Scan(&login, &minLen, &maxLen)

	// Длина логина считаем в байтах (как требует условие).
	loginLen := len(login)

	// TODO: Посчитайте isTooShort, isTooLong и isExactMin через прямые сравнения loginLen с minLen/maxLen
	isTooShort := loginLen < minLen
	isTooLong := loginLen > maxLen
	isExactMin := loginLen == minLen

	// Вывод: три bool в одной строке через пробел.
	fmt.Printf("%v %v %v", isTooShort, isTooLong, isExactMin)
}