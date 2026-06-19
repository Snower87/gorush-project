package main

/*
= Задача №14. Прайс по алфавиту =
Представьте, что вы собираете прайс‑лист: ключ — код товара (одно слово без пробелов), значение — цена (целое число). Один и тот же код может встретиться несколько раз — тогда важна последняя версия цены, как в реальном обновлении прайса.
Программа читает n, затем n пар: itemCode и price. Все пары сохраняются в map[string]int, и при повторяющемся ключе в карте должно остаться последнее значение.
Потом прайс нужно напечатать “красиво и воспроизводимо”: по возрастанию ключа (алфавитно). Каждая строка вывода: key value.
Требования:
• Программа должна прочитать из stdin целое число n, затем n пар значений: itemCode (строка без пробелов) и price (целое число).
• Все пары itemCode → price должны сохраняться в map[string]int.
• Если один и тот же itemCode встречается несколько раз, в map должно оставаться последнее прочитанное значение price для этого ключа.
• Для печати прайса программа должна сформировать отдельный слайс всех ключей из map (сбор ключей в []string).
• Слайс ключей должен быть отсортирован по алфавиту (по возрастанию строк) с помощью sort.Strings или slices.Sort.
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	priceByCode := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var code string
		var price int
		fmt.Fscan(in, &code, &price)

		// Повторяющийся код перезаписываем: важна последняя цена.
		priceByCode[code] = price
	}

	// TODO: Сделайте стабильный вывод: соберите все ключи map в отдельный слайс,
	keys := make([]string, 0, len(priceByCode))
	for key := range priceByCode {
	    keys = append(keys, key)
	}
	// TODO: отсортируйте ключи по возрастанию и напечатайте пары (key value)
	sort.Strings(keys)
	// TODO: строго в порядке отсортированных ключей (не через range по map).

    for key := range keys {
    	fmt.Fprintln(out, key, priceByCode[key])
    }
    /*
	for code, price := range priceByCode {
		fmt.Fprintln(out, code, price)
	}
    */
}