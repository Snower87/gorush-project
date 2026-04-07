/*
= Задача №19. Сдача кассы =
Вы пишете логику для мини-кассы в автомате: цена и внесённая сумма приходят в центах, причём обе величины приходят текстом
(так бывает, когда данные прилетают из простого протокола). Покупатель уже оплатил достаточно, осталось только посчитать сдачу.
Программа должна прочитать две строки priceStr и paidStr, распарсить их через strconv.Atoi, вычислить сдачу в центах changeCents,
а затем разложить её на доллары и оставшиеся центы: changeDollars и changeRestCents. После этого вывести два числа через
пробел: dollars cents.
Ввод: две строки, каждая — целое число в центах. Вывод: два целых числа через пробел: dollars cents, затем перевод строки.
Требования:
• Программа должна прочитать из stdin две строки и сохранить их в две переменные типа string (например, priceStr и paidStr).
• Программа должна преобразовать обе строки в целые числа, используя strconv.Atoi (для priceStr и paidStr).
• Код должен быть явно разделён на четыре этапа с комментариями // read, // parse, // compute, // print, и между этапами
должны быть пустые строки.
• На этапе compute программа должна вычислить сдачу в центах в отдельной переменной changeCents, а также отдельно
вычислить changeDollars и changeRestCents (или эквивалентные по смыслу переменные).
• Для получения changeDollars и changeRestCents программа должна использовать целочисленные операции / и % (деление и остаток).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	priceStr := scanner.Text()
	scanner.Scan()
	paidStr := scanner.Text()

	// parse
	priceCents, _ := strconv.Atoi(priceStr)
	paidCents, _ := strconv.Atoi(paidStr)

	// compute
	// TODO: вычислите сдачу в центах (changeCents) и разложите её на доллары (changeDollars) и оставшиеся центы (changeRestCents)
	changeCents := paidCents - priceCents
	changeDollars := changeCents / 100
	changeRestCents := changeCents % 100

	// print
	fmt.Println(changeDollars, changeRestCents)
}