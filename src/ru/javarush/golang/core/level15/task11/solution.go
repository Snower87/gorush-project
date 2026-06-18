package main

/*
= Задача №11. Печать по очереди =
Вы собираете список карточек: у каждой есть id и короткий title. Хранить удобно в map (быстро находить по id ), но печатать нужно строго в том порядке, в котором карточки вводились — как в журнале.
Программа читает число n , затем n пар: cardID (int) и cardTitle (одно слово). Значения хранятся в map[int]string. Одновременно сохраняйте порядок ввода cardID в слайс inputOrder []int.
После чтения всех данных программа выводит n строк строго в порядке ввода: cardID и соответствующий cardTitle.
Требования:
• Программа должна хранить соответствие cardID → cardTitle в структуре данных map[int]string.
• Программа должна сохранять cardID в слайс inputOrder []int в том порядке, в котором пары вводятся.
• Программа должна читать n, затем n пар (cardID int и cardTitle одно слово) и сразу при чтении записывать title в map и добавлять id в inputOrder.
• Программа должна выводить n строк, обходя inputOrder по индексам/значениям, и для каждого id брать title из map; итерация по map через range для печати не используется.
• Каждая строка вывода должна иметь формат: id, пробел, title, перевод строки.
• Если n равно 0, программа не должна выводить ничего.
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

	// Храним данные по id в map (быстрый доступ),
	// а порядок ввода отдельно — в слайсе (map порядок не гарантирует).
	cards := make(map[int]string, n)
	inputOrder := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var id int
		var title string
		fmt.Fscan(in, &id, &title)

		cards[id] = title

		// TODO: Сохраняйте в inputOrder именно введённый cardID (id), чтобы затем вывести карточки строго в порядке ввода.
		inputOrder = append(inputOrder, id)
	}

	for _, id := range inputOrder {
		// TODO: Убедитесь, что печатается title, соответствующий cardID из inputOrder (а не по индексу цикла).
		fmt.Fprintln(out, id, cards[id])
	}
}