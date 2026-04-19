package main

/*
= Задача №16. Сумма байтов =
Вы тестируете сетевой протокол, где два значения считаются байтами (0..255). Если сложить их как uint8, получится “переполнение по кругу” — ровно так, как работает байт в реальном протоколе. Но иногда нужен и “точный” результат без переполнения, чтобы сравнить и поймать ошибки.
Программа читает firstByteValue и secondByteValue как int. Если хотя бы одно число не в диапазоне 0..255, печатает "OUT". Иначе переводит оба значения в uint8, считает wrappedSum как сумму в uint8, а exactSum как сумму в int (через явное расширение). Затем печатает в одной строке wrappedSum и exactSum, разделяя их пробелом.
Требования:
• Программа должна считать `firstByteValue` и `secondByteValue` через `fmt.Scan` в переменные типа `int` (с использованием пакета `fmt`).
• Программа должна проверить через `if`, что оба введённых числа находятся в диапазоне от 0 до 255 включительно, с явными границами (`< 0` или `> 255`).
• Если хотя бы одно из чисел выходит за диапазон 0..255, программа должна вывести ровно `OUT` и не выполнять дальнейшие вычисления.
• При корректном вводе программа должна явно преобразовать оба значения из `int` в `uint8` (через `uint8(...)`) перед вычислением суммы байтов.
• Программа должна вычислить `wrappedSum` как сумму двух значений типа `uint8` так, чтобы результат оставался `uint8` и переполнение происходило “по кругу” (без дополнительных расширений типов в процессе сложения).
• Программа должна вычислить `exactSum` как сумму в `int`, выполняя явное преобразование каждого слагаемого из `uint8` в `int` перед сложением.
*/

import "fmt"

func main() {
	var firstByteValue, secondByteValue int
	fmt.Scan(&firstByteValue, &secondByteValue)

	// TODO: Реализуйте проверку, что оба числа в диапазоне 0..255 (включительно).
	// TODO: Если хотя бы одно число вне диапазона — установите outOfRange = true.
	outOfRange := (firstByteValue < 0 || firstByteValue > 255) ||
                  (secondByteValue < 0 || secondByteValue > 255)

	if outOfRange {
		fmt.Print("OUT")
		return
	}

	var firstByte, secondByte uint8
	// TODO: Явно преобразуйте входные значения из int в uint8 и присвойте в firstByte и secondByte.
	firstByte = uint8(firstByteValue)
	secondByte = uint8(secondByteValue)

	var wrappedSum uint8
	// TODO: Посчитайте wrappedSum как сумму двух uint8 так, чтобы результат оставался uint8 (с переполнением по кругу).
	wrappedSum = firstByte + secondByte

	var exactSum int
	// TODO: Посчитайте exactSum как сумму в int, явно расширив каждое слагаемое uint8 -> int перед сложением.
	exactSum = int(firstByte) + int(secondByte)

	fmt.Printf("%d %d", wrappedSum, exactSum)
}