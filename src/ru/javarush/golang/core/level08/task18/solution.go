package main

/*
= Задача №18. Заказ финальность =
Вы делаете консольную диагностику для магазина: по числовому statusCode нужно вывести название статуса заказа, а затем — понять, финальный он или нет. Финальные статусы — это те, после которых заказ уже не «двигается»: он либо доставлен, либо отменён. Все остальные (и любые неизвестные коды) считаются не финальными.
Программа читает один код, печатает имя статуса в первой строке, а во второй — final или not_final.
Требования:
• В программе должен быть объявлен пользовательский тип type OrderStatus int для представления статуса заказа.
• Должны быть объявлены константы Created, Paid, Shipped, Delivered, Cancelled через iota строго в этом порядке, чтобы значения были created=0, paid=1, shipped=2, delivered=3, cancelled=4.
• Должна быть реализована функция statusName(s OrderStatus) string, которая возвращает человекочитаемое имя статуса через switch по enum-константам, а для всех неизвестных значений возвращает "unknown" через ветку default.
• Должна быть реализована функция isFinal(s OrderStatus) bool, которая возвращает true только для статусов Delivered и Cancelled, а для всех остальных (включая unknown) возвращает false.
• Программа должна прочитать один числовой statusCode с использованием fmt.Scan(&statusCode) (один вводимый код).
*/

import "fmt"

// OrderStatus — enum-тип статуса заказа (по требованию задачи).
type OrderStatus int

const (
	Created OrderStatus = iota // 0
	Paid                      // 1
	Shipped                   // 2
	Delivered                 // 3
	Cancelled                 // 4
)

// statusName возвращает человекочитаемое имя статуса, неизвестное — "unknown".
func statusName(s OrderStatus) string {
	// TODO: Реализуйте отображение всех известных статусов в их строковые имена через switch.
	// TODO: Для неизвестных значений обязательно возвращайте "unknown" через ветку default.
	switch s {
	case Created:
		return "created"
	case Paid:
		return "paid"
	case Shipped:
		return "shipped"
	case Delivered:
		return "delivered"
	case Cancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

// isFinal — финальные статусы: Delivered и Cancelled, остальные (включая unknown) — не финальные.
func isFinal(s OrderStatus) bool {
	// TODO: Реализуйте проверку финальности статуса без "магических чисел":
	// TODO: сравнивайте только с enum-константами (Delivered/Cancelled), неизвестные считаются не финальными.
	switch s {
    case Delivered:
        return true
    case Cancelled:
        return true
	}
	return false
}

func main() {
	var statusCode int
	fmt.Scan(&statusCode)

	status := OrderStatus(statusCode)

	fmt.Println(statusName(status))
	if isFinal(status) {
		fmt.Println("final")
	} else {
		fmt.Println("not_final")
	}
}