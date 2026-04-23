package main

/*
= Задача №2. План комнаты =
Вы помогаете друзьям быстро прикинуть ремонт: они измерили комнату рулеткой и хотят сразу увидеть площадь и периметр, чтобы понять, сколько нужно напольного покрытия и плинтусов.
Данные приходят как целые числа (в условных единицах), и расчёты тоже должны быть целочисленными.
Программа читает из stdin два целых числа: roomWidth и roomHeight. Выносите математику в две отдельные функции: одна считает площадь, другая — периметр. Затем печатаете отчёт в двух строках с фиксированными подписями.
Требования:
• В коде должны быть реализованы две функции с точными сигнатурами `rectArea(width, height int) int` и `rectPerimeter(width, height int) int`.
• Каждая из функций `rectArea` и `rectPerimeter` должна возвращать ровно одно значение типа `int` через `return`.
• Внутри `rectArea` и `rectPerimeter` не должно быть чтения из stdin и печати в stdout (никаких `fmt.Scan/Print*` в этих функциях).
• Функция `main` должна считать из stdin два целых числа (ширину и высоту) с помощью `fmt.Scan`.
• `main` должна напечатать ровно две строки без лишнего текста: первая строка в формате `area=<значение>`, вторая строка в формате `perimeter=<значение>`.
*/

import "fmt"

// rectArea считает площадь прямоугольника; никакого I/O внутри быть не должно.
func rectArea(width, height int) int {
	// TODO: Реализуйте вычисление площади прямоугольника по целочисленным ширине и высоте.
	return width * height
}

// rectPerimeter считает периметр прямоугольника; никакого I/O внутри быть не должно.
func rectPerimeter(width, height int) int {
	// TODO: Реализуйте вычисление периметра прямоугольника по целочисленным ширине и высоте.
	return 2 * (width + height)
}

func main() {
	var roomWidth, roomHeight int
	fmt.Scan(&roomWidth, &roomHeight)

	area := rectArea(roomWidth, roomHeight)
	perimeter := rectPerimeter(roomWidth, roomHeight)

	fmt.Printf("area=%d\n", area)
	fmt.Printf("perimeter=%d\n", perimeter)
}