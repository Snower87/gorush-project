package main

/*
= Задача №4. Срочная вставка =
Вы пишете консольный “плейлист” из чисел, и вам нужно уметь вставлять новый трек trackID в середину списка, не теряя хвост. Если пользователь указал плохой индекс — добавляйте трек в конец, без паники.
Программа читает N, затем N целых чисел в слайс playlist, затем индекс insertIndex и значение trackID. Если 0 ≤ insertIndex ≤ N, вставка делается по алгоритму: увеличить длину на 1 через playlist = append(playlist, 0), затем сдвинуть хвост вправо copy(playlist[insertIndex+1:], playlist[insertIndex:]), и только после этого записать playlist[insertIndex] = trackID. Если индекс вне диапазона — просто добавьте trackID в конец.
Выведите итоговый слайс одной строкой: числа через пробел (после успешной вставки длина станет N+1).
Требования:
• Программа должна считать из stdin число N, затем N целых чисел в слайс playlist, затем целые insertIndex и trackID (в указанном порядке).
• Если 0 ≤ insertIndex ≤ N, программа должна выполнить вставку строго по шагам: сначала увеличить длину `playlist = append(playlist, 0)`, затем сдвинуть хвост вправо `copy(playlist[insertIndex+1:], playlist[insertIndex:])`, и только после этого записать `playlist[insertIndex] = trackID`.
• Если insertIndex вне диапазона [0, N], программа должна добавить trackID в конец (эквивалент вставки в позицию N+1) без выхода за границы и без паники.
• Программа не должна создавать отдельный слайс размера O(N) для результата; все изменения выполняются в playlist, используя слайс, который вернул append.
• Программа должна корректно обрабатывать случай N = 0 (пустой исходный playlist), включая вставку по индексу 0 и добавление в конец при некорректном индексе.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	playlist := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &playlist[i])
	}

	var insertIndex, trackID int
	fmt.Fscan(in, &insertIndex, &trackID)

	if 0 <= insertIndex && insertIndex <= n {
		// TODO: Реализуйте вставку trackID в playlist по индексу insertIndex без создания второго слайса размера O(N).
		// TODO: Вставка должна быть выполнена через увеличение длины append, затем сдвиг хвоста вправо через copy, затем запись trackID в позицию insertIndex.
		playlist = append(playlist, 0)
		copy(playlist[insertIndex+1:], playlist[insertIndex:])
		playlist[insertIndex] = trackID
	} else {
		// TODO: Обработайте некорректный индекс: добавьте trackID в конец без паники и без выхода за границы.
		playlist = append(playlist, trackID)
	}

	for i, v := range playlist {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
}