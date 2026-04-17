package main

/*
= Задача №16. Чек и отладка =
Вы пишете консольный прототип кассы. Пользователь вводит товар item (одно слово), количество qty и цену price (оба целые). Программа должна посчитать сумму sum = qty * price, напечатать “чек”, а затем — отладочную строку, где видны типы всех переменных (чтобы убедиться, что нигде не “поплыл” тип).
Нужно вывести три строки:
- Item=<item> Qty=<qty> Price=<price>
- Sum=<sum>
- DEBUG: itemType=<тип> qtyType=<тип> priceType=<тип> sumType=<тип>
В строках 1–2 используйте %s и %d. В строке DEBUG используйте %T и выводите именно типы переменных.
Требования:
• Программа должна прочитать из stdin три значения через fmt.Scan (или fmt.Fscan): item (одно слово), qty (целое), price (целое).
• Программа должна хранить item в переменной типа string, а qty, price и sum — в переменных типа int.
• Программа должна вычислить sum как результат умножения qty * price.
• Программа должна вывести ровно три строки в указанном порядке, без пропусков и без дополнительных строк: строку с Item/Qty/Price, затем строку с Sum, затем строку DEBUG.
• Первая строка должна иметь вид `Item=<item> Qty=<qty> Price=<price>`, а вторая — `Sum=<sum>`; для вывода значений в этих строках нужно использовать fmt.Printf с плейсхолдерами %s (для item) и %d (для целых).
*/

import "fmt"

func main() {
	var item string
	var qty, price int

	fmt.Scan(&item, &qty, &price)

	var sum int
	// TODO: Вычислите сумму покупки: sum = qty * price.
	sum = qty * price

	// TODO: Выведите первую строку чека в нужном формате через fmt.Printf (используйте %s для item и %d для целых).
	// TODO: Выведите вторую строку с суммой в нужном формате через fmt.Printf (используйте %d).
	// TODO: Выведите третью строку DEBUG с типами переменных item, qty, price, sum через fmt.Printf (используйте %T).
	fmt.Printf("Item=%s Qty=%d Price=%d\n", item, qty, price)
    fmt.Printf("Sum=%d\n", sum)
    fmt.Printf("DEBUG: itemType=%T qtyType=%T priceType=%T sumType=%T\n", item, qty, price, sum)
}