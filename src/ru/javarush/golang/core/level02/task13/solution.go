package main

import (
	"fmt"
	"strconv"
)

func main() {
	var idStr string
	fmt.Scan(&idStr)

	_ = strconv.IntSize // оставляем импорт strconv используемым

	var id int
	var err error

	// TODO: Преобразуйте строку idStr в число через strconv.Atoi и сохраните результаты в переменные id и err.

	nextID := id
	// TODO: Вычислите следующий идентификатор: nextID должен быть равен id + 1 (без дополнительных проверок).

	fmt.Println("id =", id)
	fmt.Println("nextID =", nextID)
	fmt.Println("err =", err)
}