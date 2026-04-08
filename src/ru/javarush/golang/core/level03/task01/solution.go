package main

/*
= Задача №1. Дуэль ставок =
Вы пишете крошечный модуль для мини-аукциона: два участника одновременно называют свои ставки, и системе нужно мгновенно
понять две вещи — одинаковые ли ставки и кто перебил кого.
Программа читает из stdin два целых числа: firstBid и secondBid. Затем она вычисляет два булевых значения: isEqualBid
(равны ли ставки) и isFirstHigher (первая ставка больше второй). После этого программа печатает оба результата на одной
строке через пробел: сначала isEqualBid, затем isFirstHigher.
Ввод: два целых числа, разделённые пробелами или переводом строки. Вывод: два булевых значения.
Пример ввода:
 10 7
Пример вывода:
 false true
Требования:
• Программа должна прочитать два целых числа из stdin в переменные firstBid и secondBid типа int, используя fmt.Scan.
• Программа должна вычислить булево значение isEqualBid как результат сравнения firstBid == secondBid и сохранить его в отдельную переменную типа bool.
• Программа должна вычислить булево значение isFirstHigher как результат сравнения firstBid > secondBid и сохранить его в отдельную переменную типа bool.
• В решении не должны использоваться if/else и любые циклы; логика должна быть выражена только через операции сравнения и присваивания.
*/

import "fmt"

func main() {
	var firstBid, secondBid int
	fmt.Scan(&firstBid, &secondBid)

	// TODO: Вычислите, равны ли ставки (firstBid и secondBid), и сохраните результат в isEqualBid без if/else и циклов.
	var isEqualBid bool
	isEqualBid = firstBid == secondBid

	// TODO: Вычислите, больше ли первая ставка второй (firstBid и secondBid), и сохраните результат в isFirstHigher без if/else и циклов.
	var isFirstHigher bool
	isFirstHigher = firstBid > secondBid

	fmt.Println(isEqualBid, isFirstHigher)
}