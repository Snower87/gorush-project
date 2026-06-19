package main

/*
= Задача №18. Проверка доступа =
Вы делаете простую систему “списка разрешённых”: есть словарь известных кодовых фраз, и дальше идут запросы — нужно отвечать, есть ли фраза в списке. Это похоже на проверку доступа: "знаешь пароль?" → "YES/NO".
Ввод: целое число n, затем n строк (словарь). Далее целое число q, затем q строк (запросы).
Вывод: для каждого запроса вывести на новой строке YES, если строка есть в словаре, иначе NO.
Храните словарь в map[string]struct{} и проверяйте наличие через ok‑идиому: _, ok := set[x].
Требования:
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

	// Словарь храним как set: ключ есть -> фраза разрешена.
	dictionary := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var phrase string
		fmt.Fscan(in, &phrase)

		// TODO: Добавьте phrase в dictionary как в множество (set) через map[string]struct{}.
		dictionary[phrase] = struct{}{}
	}

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var query string
		fmt.Fscan(in, &query)

		// TODO: Реализуйте проверку наличия query в dictionary строго через ok-идиому.
		// TODO: Выведите "YES", если query есть в словаре, иначе "NO". Обрабатывайте запросы потоково.
		_, ok := dictionary[query]
		if ok {
		    fmt.Fprintln(out, "YES")
		} else {
		    fmt.Fprintln(out, "NO")
		}
	}
}