package main

import "fmt"

func main() {
	// Исходные данные заданы прямо в коде и имеют тип int (без stdin).
	var price int = 1999
	var discountPercent int = 15
	var quantity int = 3

	// TODO: Вычислите discountAmount целочисленно по условию (без float), со скобками для приоритета операций.
	discountAmount := (price * discountPercent) / 100 //размер скидки

	// TODO: Вычислите finalPrice как цену после скидки, используя discountAmount.
	finalPrice := price - discountAmount //цену после скидки

	// TODO: Вычислите total как итоговую стоимость всех товаров, используя finalPrice и quantity.
	total := finalPrice * quantity //общий итог

	// Вывод строго: discountAmount finalPrice total.
	fmt.Print(discountAmount, " ", finalPrice, " ", total)
}