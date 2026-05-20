package main

/*
= Задача №16. Последние элементы =
Вы делаете “хвостовой отчёт”: иногда нужно взять последние k значений (например, последние измерения датчика) и отправить их дальше, но строго отдельным буфером, чтобы дальнейшие правки не затронули исходные данные.
Программа должна прочитать N, затем N целых чисел (слайс nums), затем прочитать k.
В программе должна быть функция SnapshotLastK(nums []int, k int) []int, которая возвращает последние k элементов nums как независимую копию (с отдельным backing array). Если k > len(nums), функция возвращает копию всего nums. Если k <= 0, функция возвращает пустой слайс.
После вызова функции программа должна, если результат не пустой, изменить result[0] на 777. Затем вывести две строки: nums и result (через fmt.Println).
func SnapshotLastK(nums []int, k int) []int {
    // ваша реализация
}
Требования:
• Программа должна прочитать из stdin число N, затем N целых чисел в слайс nums, затем число k.
• В программе должна быть реализована отдельная функция SnapshotLastK(nums []int, k int) []int с указанной сигнатурой.
• Функция SnapshotLastK должна возвращать последние k элементов nums; если k > len(nums) — возвращать копию всего nums; если k <= 0 — возвращать пустой слайс с len=0.
• SnapshotLastK должна создавать результат как новый буфер (с отдельным backing array) с использованием make и copy, а не возвращать под-слайс, разделяющий память с nums.
• После вызова SnapshotLastK программа должна, если result не пустой, изменить result[0] на 777; это изменение не должно приводить к изменению nums.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func SnapshotLastK(nums []int, k int) []int {
	if k <= 0 {
		return make([]int, 0)
	}

	if k >= len(nums) {
		// TODO: Верните независимую копию всего nums (новый backing array), а не исходный слайс.
		temp := make([]int, len(nums))
		copy(temp, nums)
		return temp
	}

	start := len(nums) - k
	// TODO: Верните независимую копию последних k элементов через make+copy, а не под-слайс, который делит память с nums.
	res := make([]int, k)
	copy(res, nums[start:])
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var k int
	fmt.Fscan(in, &k)

	result := SnapshotLastK(nums, k)

	// Проверяем независимость: меняем result, nums не должен измениться.
	if len(result) > 0 {
		result[0] = 777
	}

	fmt.Println(nums)
	fmt.Println(result)
}