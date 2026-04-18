package main

/*
= Задача №8. Долгая копилка =
Вы моделируете “копилку”, куда много раз подряд добавляют по 0.1. Кажется, что через n шагов должно получиться ровно n/10. Но в реальной арифметике float32 и float64 могут накопить разные погрешности, и вы хотите это измерить.
Программа читает целое число steps (тип int). Затем считает две суммы: sum64 в float64 и sum32 в float32, каждый раз добавляя 0.1 ровно steps раз. После этого вычисляет ожидаемое expected = float64(steps) / 10.0, а также отклонения delta64 и delta32 (абсолютная разница), причём абсолютное значение делается через if, без math.Abs. Итог — четыре строки с точностью 17 знаков.
Требования:
• Программа должна считать целое число steps типа int с помощью fmt.Scan(&steps).
• Программа должна завести переменные sum64 (тип float64, начальное значение 0.0) и sum32 (тип float32, начальное значение 0) и прибавить к каждой по 0.1 ровно steps раз.
• Добавление 0.1 в обе суммы должно выполняться циклом вида for i := 0; i < steps; i++.
• Программа должна вычислить expected как float64(steps)/10.0 и посчитать delta64 и delta32 как абсолютную разницу между суммой и expected, получая модуль через if без использования math.Abs.
• При вычислении delta32 программа должна явно преобразовать sum32 в float64 перед сравнением/вычитанием с expected.
*/

import "fmt"

func main() {
	var steps int
	fmt.Scan(&steps)

	var sum64 float64 = 0.0
	var sum32 float32 = 0.0

	for i := 0; i < steps; i++ {
		// TODO: прибавьте 0.1 к sum64 (float64) и sum32 (float32) на каждой итерации
		sum64 += 0.1
		sum32 += 0.1
	}

    // TODO: вычислите ожидаемое значение суммы (expected) для заданного числа шагов
    expected := float64(steps) / 10.0

	// TODO: вычислите delta64 как модуль разницы между sum64 и expected через if (без math.Abs)
	delta64 := sum64 - expected
	if sum64 - expected < 0 {
	    delta64 = 0 - delta64
	}

	// TODO: вычислите delta32 как модуль разницы между float64(sum32) и expected через if (без math.Abs)
	delta32 := float64(sum32) - expected
	if float64(sum32) - expected < 0 {
	    delta32 = 0 - delta32
	}

	fmt.Printf("sum64=%.17f\n", sum64)
	fmt.Printf("sum32=%.17f\n", sum32)
	fmt.Printf("delta64=%.17f\n", delta64)
	fmt.Printf("delta32=%.17f\n", delta32)
}