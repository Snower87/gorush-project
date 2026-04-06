package main

import (
	"fmt"
	"os"
)

func main() {
	// Входные данные (цена и количество) — объявляем до чтения.
	var price, qty int

	// Переменная err нужна для вывода.
	//_, err := 0, error(nil)
	// TODO: Считайте price и qty из os.Stdin через fmt.Fscan, игнорируя количество считанных значений через "_", и сохраните ошибку в err.
    _, err := fmt.Fscan(os.Stdin, &price, &qty)
	total := 0
	// TODO: Вычислите total как произведение price * qty.
    total = price * qty
	// Сейчас формат вывода намеренно НЕ соответствует требованиям.
	fmt.Printf("err:%s\n", err)
	// TODO: Выведите err ровно в формате `err:<значение>` (без пробелов) и с переводом строки.

	fmt.Printf("total:%d\n", total)
	// TODO: Выведите total ровно в формате `total:<число>` (без пробелов) и с переводом строки.
}