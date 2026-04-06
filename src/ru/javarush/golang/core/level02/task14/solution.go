package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orderNumber int
	fmt.Scan(&orderNumber)

	// TODO: Преобразуйте orderNumber в строку строго через strconv.Itoa и сохраните в переменную.
	orderStr := strconv.Itoa(0)
	_ = orderStr

	// TODO: Соберите итоговую метку заказа конкатенацией через + в формате "ORD-" + <строковый номер>.
	label := "ORD-"

	// TODO: Выведите ровно одну строку с готовой меткой без лишних слов и форматирования.
	fmt.Println(label)
}