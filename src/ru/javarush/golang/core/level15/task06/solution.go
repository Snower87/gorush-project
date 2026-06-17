package main

/*
= Задача №6. Статус доставки =
Вы пишете маленький модуль для службы доставки: внутри программы есть таблица map[string]bool, где ключ — orderID, а значение — доставлен ли заказ.
Важно помнить: значение false — это не “заказа нет”, это “заказ есть, но ещё не доставлен”. Поэтому проверять нужно именно наличие ключа через ok‑идиому.
Программа читает orderID. Если заказ известен и значение true, выведите delivered. Если заказ известен и значение false, выведите not delivered. Если ключа нет — unknown order.
Требования:
• В программе должна быть создана переменная типа map[string]bool с заранее заданными значениями, причём в этой map должны присутствовать как минимум один заказ со значением true и как минимум один заказ со значением false.
• Программа должна читать одну строку (orderID) через fmt.Scan.
• Статус заказа должен извлекаться только через конструкцию v, ok := orders[orderID], чтобы отличать “ключа нет” от значения false.
• Логика ветвления должна быть оформлена именно в виде if v, ok := orders[orderID]; ok { ... } else { ... }.
• Если ok == false, программа должна вывести unknown order; если ok == true и v == true — delivered; если ok == true и v == false — not delivered (ветка unknown order не должна зависеть от v).
*/

import "fmt"

func main() {
	// false тут значит "заказ есть, но ещё не доставлен", поэтому при чтении нужен ok (v, ok := m[k])
	orders := map[string]bool{
		"ORD-100": true,
		"ORD-200": false,
	}

	var orderID string
	fmt.Scan(&orderID)

	// Требование: именно if с коротким объявлением + ok-идиома.
	if v, ok := orders[orderID]; ok {
		// TODO: Реализуйте вывод статуса для известного заказа.
		// TODO: Важно различать "ключа нет" и "значение false": unknown order зависит только от ok.
		// TODO: Должна быть выведена ровно одна строка без лишних переводов строк.
		if ok == false {
		    fmt.Print("unknown order")
		}
        if ok == true && v == true {
            fmt.Print("delivered")
        }
        if ok == true && v == false {
            fmt.Print("not delivered")
        }
	} else {
		fmt.Print("unknown order")
	}
}