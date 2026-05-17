package main

/*
= Задача №11. Безопасный префикс =
Вы обрабатываете поток чисел: это может быть телеметрия, цены, очки — неважно. Вам нужно взять первые k чисел как префикс, а потом дописать в этот префикс ещё одно значение x… но так, чтобы исходный массив данных nums не “поплыл” и не испортился из-за append.
Программа должна прочитать из stdin: n, затем n целых чисел (слайс nums), затем числа k и x. После этого сформировать prefix из первых k элементов так, чтобы у prefix ёмкость была ограничена длиной ( cap(prefix) == len(prefix) ), и выполнить prefix = append(prefix, x).
Требования:
• Программа должна прочитать из stdin число n, затем n целых чисел в слайс nums типа []int, затем числа k и x, используя fmt.Scan или fmt.Fscan.
• Программа должна сформировать prefix как полный срез первых k элементов: prefix := nums[:k:k], чтобы выполнялось cap(prefix) == len(prefix).
• Программа должна выполнить prefix = append(prefix, x), обязательно сохранив результат append обратно в prefix.
• В решении нельзя использовать copy, массивы вида [N]T и пакет slices; исходные данные nums должны оставаться слайсом []int.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var k, x int
	fmt.Fscan(in, &k, &x)

	// TODO: Сформируйте prefix из первых k элементов так, чтобы cap(prefix) == len(prefix),
	// TODO: и append в prefix не мог изменить исходный слайс nums.
	prefix := nums[:k:k]

	// TODO: Выполните append и обязательно сохраните результат обратно в prefix.
	prefix = append(prefix, x)

	fmt.Println(nums)
	fmt.Println(prefix)
}