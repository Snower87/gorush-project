package main

/*
= Задача №2. Проверка роли =
Представьте, что вы делаете простую консольную админку. Пользователь вводит свою роль одним словом (без пробелов),
а программа должна быстро собрать “флажки”, чтобы другие части системы могли решать, что разрешено.
Программа читает строку userRole. Затем вычисляет и выводит на одной строке три булевых значения через пробел:
isAdmin — userRole равно строке "admin"
isGuest — userRole равно строке "guest"
isNotUser — userRole не равно строке "user"
Ввод: одна строка userRole. Вывод: isAdmin isGuest isNotUser.
Пример ввода:
admin
Пример вывода:
true false true
Требования:
• Программа должна прочитать из stdin одну строку `userRole` типа `string` с помощью `fmt.Scan`.
• Программа должна вычислить три значения типа `bool`: `isAdmin` ( `userRole == "admin"` ), `isGuest` ( `userRole == "guest"` ), `isNotUser` ( `userRole != "user"` ), используя только операторы `==` и `!=` для сравнения строк.
• Результаты `isAdmin`, `isGuest`, `isNotUser` должны быть сохранены в отдельные переменные типа `bool`.
• В решении не должны использоваться конструкции `if/else`.
*/

import "fmt"

func main() {
	var userRole string
	fmt.Scan(&userRole)

	// Флажки разрешений: только сравнение строк через == и != (без if/else).
	var isAdmin bool
	var isGuest bool
	var isNotUser bool

	// TODO: Вычислите значения isAdmin, isGuest, isNotUser через сравнение userRole со строками из условия (используйте только == и !=, без if/else).
	isAdmin = userRole == "admin"
	isGuest = userRole == "guest"
	isNotUser = userRole != "user"

	// Ровно три bool через пробел, без лишнего текста.
	fmt.Printf("%t %t %t", isAdmin, isGuest, isNotUser)
}