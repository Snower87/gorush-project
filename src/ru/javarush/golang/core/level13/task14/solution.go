package main

/*
= Задача №14. Вставить партию =
Вы пишете учёт поставок: есть текущий список warehouseStock (числа), и приходит партия incomingBatch из K чисел, которую нужно вставить внутрь списка на позицию insertPos. Важно: вставляется сразу весь блок, порядок внутри блока сохраняется.
Программа читает N, затем N целых чисел в warehouseStock, затем K, затем K целых чисел в incomingBatch, затем индекс insertPos. Реализуйте вставку в функции insertSlice(s []int, insertPos int, add []int) []int.
Допустимые значения insertPos: от 0 до len(s) включительно. Если insertPos вне диапазона — выведите исходный список без изменений. Если K == 0, результат — тоже исходный список.
Итоговый слайс выведите в stdout в одну строку: числа через пробел.
Требования:
• В программе должна быть реализована функция insertSlice(s []int, insertPos int, add []int) []int, которая возвращает итоговый слайс после вставки.
• Функция должна считать допустимыми значения insertPos только в диапазоне от 0 до len(s) включительно; если insertPos вне диапазона — вернуть исходный s без изменений (и программа должна вывести исходный список).
• Если K == 0 (то есть add пустой), функция должна возвращать s без изменений.
• Функция должна вставлять весь блок add целиком, сохраняя порядок элементов add; случай insertPos == len(s) должен приводить к корректному добавлению блока в конец.
• Увеличение длины результата должно выполняться строго через append(s, make([]int, K)...), затем хвост исходного слайса должен быть сдвинут вправо на K позиций с помощью copy, и после этого элементы add должны быть скопированы в освобождённый диапазон с помощью copy.
• Нельзя использовать пакет slices и нельзя собирать результат “склейкой” трёх частей через append без выполнения сдвига хвоста через copy.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func insertSlice(s []int, insertPos int, add []int) []int {
	k := len(add)
	if k == 0 {
		return s
	}
	if insertPos < 0 || insertPos > len(s) {
		return s
	}

	// Увеличиваем длину строго через append + make (по требованиям).
	s = append(s, make([]int, k)...)

	// TODO: Реализуйте вставку блока add в позицию insertPos.
	// TODO: Нужно сдвинуть хвост исходного слайса вправо на k позиций с помощью copy,
	// TODO: затем скопировать элементы add в освободившийся диапазон с помощью copy.
	copy(s[insertPos+k:], s[insertPos:])
	copy(s[insertPos:], add)

	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	warehouseStock := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &warehouseStock[i])
	}

	var k int
	fmt.Fscan(in, &k)

	incomingBatch := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &incomingBatch[i])
	}

	var insertPos int
	fmt.Fscan(in, &insertPos)

	res := insertSlice(warehouseStock, insertPos, incomingBatch)

	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, res[i])
	}
	fmt.Fprintln(out)
}