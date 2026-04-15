package main

/*
= Задача №6. Пост охраны =
Вы пишете мини-пропускной пункт: охранник вводит login (одно слово) и code (целое число).
Внутри программы вы делаете две проверки: правильный ли логин и правильный ли код.
Считайте, что админ — это только тот, у кого login == "admin", а код верный только если code == 4321.
После этого вычислите canEnter = isAdmin && isCodeOK и выведите:
- строку canEnter=<значение>
- строку с итогом: ACCESS GRANTED или ACCESS DENIED
Требования:
• Программа должна считать со stdin два значения: login (одно слово) и code (целое число), используя fmt.Scan или fmt.Fscan.
• Переменные login и code должны иметь типы string и int соответственно; переменные isAdmin и isCodeOK должны иметь тип bool.
• Переменная isAdmin должна вычисляться только сравнением login == "admin".
• Переменная isCodeOK должна вычисляться только сравнением code == 4321.
• Переменная canEnter должна вычисляться только как isAdmin && isCodeOK (без других логических операторов).
*/

import "fmt"

func main() {
	var login string
	var code int

	fmt.Scan(&login, &code)

	var isAdmin bool = login == "admin"
	// TODO: Сравните login с "admin" и запишите результат в isAdmin (только через ==).

	var isCodeOK bool = code == 4321
	// TODO: Сравните code с 4321 и запишите результат в isCodeOK (только через ==).

	var canEnter bool = isAdmin && isCodeOK
	// TODO: Вычислите canEnter как isAdmin && isCodeOK (только через &&).

	fmt.Printf("canEnter=%v\n", canEnter)

	if canEnter {
		fmt.Println("ACCESS GRANTED")
	} else {
		fmt.Println("ACCESS DENIED")
	}
}