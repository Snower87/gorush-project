package main

/*
= Задача №10. Режим удаления =
Вы пишете небольшую консольную адресную книгу. Пользователь вводит список имён, потом выбирает режим удаления: аккуратный stable или быстрый unstable. И вводит номер записи “по-человечески”, начиная с 1.
Программа читает n, затем n строк в contacts, затем строку mode (либо "stable", либо "unstable"), затем число k — номер удаляемого элемента (1-based). Переведите номер в индекс removeIndex := k - 1 в одном месте и дальше работайте только с индексом. Если k некорректен — выводите список без изменений. Если mode неизвестен — тоже без изменений. В обоих режимах очищайте освобождённый слот через clear, чтобы не удерживать ссылку в backing array.
Вывод: каждая строка списка — на новой строке, без лишнего текста.
Требования:
• Программа должна содержать функции removeAtStable([]string, int) []string и removeAtUnstable([]string, int) []string, которые удаляют элемент по индексу removeIndex из слайса строк и возвращают новый (resliced) результат.
• Программа должна прочитать k (1-based) и перевести его в removeIndex := k - 1 ровно в одном месте; далее во всей логике удаления использовать только removeIndex.
• В режиме "stable" удаление должно выполняться через сдвиг элементов copy (например, copy(contacts[i:], contacts[i+1:])) и последующий reslice (уменьшение длины на 1), при этом относительный порядок оставшихся элементов должен сохраняться.
• В режиме "unstable" удаление должно выполняться через замену удаляемого элемента последним (swap-with-last: contacts[i] = contacts[len-1]) и последующий reslice (уменьшение длины на 1); порядок элементов при этом может измениться.
• В обоих режимах перед уменьшением длины слайса программа должна очистить освобождённый слот в backing array с помощью clear (например, clear(contacts[len(contacts)-1:])) чтобы не удерживать ссылку на удалённую строку.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func removeAtStable(contacts []string, removeIndex int) []string {
	// TODO: Реализуйте стабильное удаление элемента по индексу removeIndex (с сохранением порядка).
	// TODO: Сдвиг элементов сделайте через copy.
	// TODO: Перед уменьшением длины очистите освободившийся слот в backing array через clear.
	// TODO: Верните результат через reslice (уменьшение длины на 1).
	if removeIndex < 0 || removeIndex >= len(contacts) {
	    return contacts
	}
	copy(contacts[removeIndex:], contacts[removeIndex+1:])
	clear(contacts[len(contacts)-1:])
	return contacts[:len(contacts)-1]
}

func removeAtUnstable(contacts []string, removeIndex int) []string {
	// TODO: Реализуйте быстрое (unstable) удаление элемента по индексу removeIndex (порядок может измениться).
	// TODO: Используйте подход "swap-with-last": положите последний элемент на место удаляемого.
	// TODO: Перед уменьшением длины очистите освобождённый слот в backing array через clear.
	// TODO: Верните результат через reslice (уменьшение длины на 1).
	n := len(contacts)
	contacts[removeIndex] = contacts[n-1]
	clear(contacts[n-1:])
	return contacts[:n-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	contacts := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &contacts[i])
	}

	var mode string
	fmt.Fscan(in, &mode)

	var k int
	fmt.Fscan(in, &k)

	// Перевод "человеческого" номера в индекс — ровно в одном месте
	removeIndex := k - 1

	res := contacts
	if removeIndex >= 0 && removeIndex < len(contacts) {
		if mode == "stable" {
			res = removeAtStable(res, removeIndex)
		} else if mode == "unstable" {
			res = removeAtUnstable(res, removeIndex)
		}
	}

	for _, c := range res {
		fmt.Fprintln(out, c)
	}
}