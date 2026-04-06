package main

import (
    "fmt"
        )

func main() {
	// Переменные объявлены заранее, чтобы можно было передавать их адреса в Scan.
	var x int
	var s string

	// Эти переменные должны хранить результат чтения.
	//count := 0
	var err error = nil

	// TODO: Считайте x и s из stdin одним вызовом fmt.Scan, передав &x и &s, и сохраните результаты в count и err.
	count, err := fmt.Scan(&x, &s)

	// TODO: Выведите ровно 4 строки в требуемом формате и порядке: count, err, x, s (каждая строка с подписью).
	fmt.Println("count:", count)
	fmt.Println("err:", err)
	fmt.Println("x:", x)
	fmt.Println("s:",s)
}