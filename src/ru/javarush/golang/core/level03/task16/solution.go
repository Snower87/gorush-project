package main

/*
= Задача №16. Маршрут посылки =
Вы делаете консольный трекер заказов. У заказа есть статус ("new", "canceled", "done") и два флага: оплачено ли (paid) и отправлено ли (shipped). По этим данным нужно вывести одно понятное слово для экрана оператора — и при этом не утонуть во вложенных проверках.
Программа читает: status (строка), paid (0 или 1), shipped (0 или 1). status может быть "new", "canceled" или "done".
- Если status равен "canceled", программа должна вывести "canceled" и завершить работу (значения paid/shipped в этом случае игнорируются).
- Если status не равен "new", "canceled" или "done", программа должна вывести "invalid status" и завершить работу.
- Если paid или shipped не равны 0/1, программа должна вывести "invalid flags" и завершить работу.
Дальше правила такие:
* Для status == "new":
   - paid == 0 → "waiting_payment"
   - paid == 1 и shipped == 0 → "waiting_shipment"
   - paid == 1 и shipped == 1 → "in_transit"
* Для status == "done":
   - paid == 0 или shipped == 0 → "invalid state"
   - paid == 1 и shipped == 1 → "completed"
Требования:
• Программа должна считать три значения через fmt.Scan(&status, &paid, &shipped), где status — строка, paid и shipped — целые флаги (0 или 1).
• Если status == "canceled", программа должна вывести "canceled" и немедленно завершить работу (paid/shipped игнорируются). Если status не равен "new", "canceled" или "done", программа должна вывести "invalid status" и немедленно завершить работу. Если paid или shipped не равны 0 или 1, программа должна вывести "invalid flags" и немедленно завершить работу.
• После успешной валидации программа должна определить и вывести ровно одно итоговое слово по правилам: для status == "new": waiting_payment / waiting_shipment / in_transit в зависимости от paid/shipped; для status == "done": вывести "invalid state", если paid == 0 или shipped == 0, иначе "completed" при paid == 1 и shipped == 1.
• В решении нельзя использовать switch и нельзя использовать циклы; ранние выходы должны быть реализованы через guard clauses (печать + return), а основная логика после валидации должна оставаться плоской и читаемой.
*/

import "fmt"

func main() {
	var status string
	var paid, shipped int

	fmt.Scan(&status, &paid, &shipped)

	// Отмена имеет приоритет: paid/shipped игнорируются полностью.
	if status == "canceled" {
		fmt.Print("canceled")
		return
	}
	// Валидируем статус сразу, чтобы дальше логика оставалась плоской.
	if status != "new" && status != "done" {
		fmt.Print("invalid status")
		return
	}

	// Флаги должны быть строго 0/1.
	if (paid != 0 && paid != 1) || (shipped != 0 && shipped != 1) {
		fmt.Print("invalid flags")
		return
	}

	// TODO: Реализуйте вывод итогового слова по правилам задачи для валидных status/paid/shipped.
	// TODO: Решение должно быть без switch и без циклов, с плоской логикой и ранними выходами (guard clauses).
	if status == "new" {
		// TODO: Реализуйте правила для статуса new.
		//fmt.Print("TODO")
		if paid == 0 {
		    fmt.Print("waiting_payment")
		    return
		}
        if paid == 1 && shipped == 0 {
            fmt.Print("waiting_shipment")
            return
        }
        if paid == 1 && shipped == 1 {
            fmt.Print("in_transit")
            return
        }
	}

	if status == "done" {
	    // TODO: Реализуйте правила для статуса done.
	    if paid == 0 || shipped == 0 {
	        fmt.Print("invalid state")
	        return
	    }
	    if paid == 1 && shipped == 1 {
	        fmt.Print("completed")
	        return
	    }
	}
}