package main

/*
= Задача №13. Вставить код =
Вы помогаете в маленьком складском приложении: есть список кодов товаров productCodes, и менеджер хочет вставить новый код в конкретное место списка. Всё должно работать и для вставки в начало, и для вставки в конец, а при ошибочном индексе — ничего не ломать и просто вернуть список как есть.
Программа читает N, затем N целых чисел в productCodes, затем два целых числа insertIndex и newCode. Вставку реализуйте в функции insertAt(s []int, insertIndex int, newCode int) []int. Если insertIndex в диапазоне от 0 до len(s) включительно — вставьте newCode в позицию insertIndex. Если индекс некорректный — верните исходный слайс без изменений.
Итоговый слайс выведите в одной строке: числа через пробел.
Требования:
• Программа должна прочитать целое число N, затем N целых чисел в слайс productCodes, затем два целых числа insertIndex и newCode.
• В программе должна быть реализована функция insertAt(s []int, insertIndex int, newCode int) []int, которая выполняет вставку элемента и возвращает итоговый слайс.
• Функция insertAt должна вставлять newCode только если insertIndex находится в диапазоне от 0 до len(s) включительно; при любом другом значении insertIndex функция должна вернуть исходный слайс без изменений.
• Внутри insertAt увеличение длины слайса должно выполняться через append с обязательным присваиванием результата в переменную, а сдвиг элементов вправо должен выполняться через copy по перекрывающимся диапазонам (для освобождения позиции insertIndex).
• Реализация должна корректно обрабатывать случаи insertIndex == 0 (вставка в начало) и insertIndex == len(s) (вставка в конец).
*/

import (
	"bufio"
	"fmt"
	"os"
)

func insertAt(s []int, insertIndex int, newCode int) []int {
	if insertIndex < 0 || insertIndex > len(s) {
		return s
	}

	// TODO: Реализуйте вставку newCode в позицию insertIndex.
	// TODO: Увеличение длины сделайте через append (с обязательным присваиванием результата),
	// TODO: а сдвиг элементов вправо — через copy по перекрывающимся диапазонам.
	// TODO: Должны корректно работать случаи insertIndex == 0 и insertIndex == len(s).
	s = append(s, 0)    //1. увеличили len на 1
	// copy(--> куда, ?? откуда)
	copy(s[insertIndex+1:], s[insertIndex:])  //2. двигаем хвост вправо
	s[insertIndex] = newCode  //3. пишем новое значение
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	productCodes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &productCodes[i])
	}

	var insertIndex, newCode int
	fmt.Fscan(in, &insertIndex, &newCode)

	productCodes = insertAt(productCodes, insertIndex, newCode)

	for i, v := range productCodes {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
}