package main

/*
= Задача №7. Датчики стабильности =
Представьте, что у вас три датчика (например, температура в трёх точках комнаты). Они присылают три вещественных значения x, y, z (тип float64). Вам нужно быстро понять диапазон, среднее и “стабильность” показаний.
Посчитайте:
- minVal — минимальное из трёх
- maxVal — максимальное из трёх
- avg — среднее арифметическое
- stable — булево значение: разница между максимумом и минимумом не больше 1.0, то есть (maxVal - minVal) <= 1.0
Затем выведите 4 строки в формате min=..., max=..., avg=..., stable=....
Требования:
• Программа должна считать из stdin три числа x, y, z как значения типа float64.
• Программа должна вычислить minVal (минимум из x, y, z) и maxVal (максимум из x, y, z) используя только ветвления if/else, без массивов/слайсов и без циклов.
• Программа должна вычислить avg как (x + y + z) / 3 в арифметике float64.
• Программа должна вычислить stable типа bool по правилу (maxVal - minVal) <= 1.0; значение stable должно выводиться как true или false.
• Программа должна вывести ровно 4 строки (каждая с новой строки) в формате: min=..., max=..., avg=..., stable=... (с точными префиксами min=, max=, avg=, stable=).
*/

import "fmt"

func main() {
	var x, y, z float64
	fmt.Scan(&x, &y, &z)

	// TODO: Найдите минимальное из трёх значений x, y, z (тип float64) только через if/else
	minVal := x
	if y < minVal {
	    minVal = y
	}
	if z < minVal {
	    minVal = z
	}

	// TODO: Найдите максимальное из трёх значений x, y, z (тип float64) только через if/else
	maxVal := x
	if y > maxVal {
	    maxVal = y
	}
	if z > maxVal {
	    maxVal = z
	}

	avg := (x + y + z) / 3

	// TODO: Вычислите stable по правилу: (maxVal - minVal) <= 1.0.
	stable := (maxVal - minVal) <= 1.0

	fmt.Printf("min=%v\n", minVal)
	fmt.Printf("max=%v\n", maxVal)
	fmt.Printf("avg=%v\n", avg)
	fmt.Printf("stable=%v\n", stable)
}