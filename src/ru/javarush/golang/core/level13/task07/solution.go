package main

/*
= Задача №7. Аккуратно удалить =
Вы собираете список гостей на мероприятие. Если кто-то отменился, нужно удалить его по индексу, но порядок остальных гостей должен сохраниться. И как хорошая привычка — вы чистите освободившийся слот в хвосте backing array, чтобы там не “висела” старая строка.
Реализуйте функцию removeAtStable(guestNames []string, removeIndex int) []string. Если removeIndex вне диапазона [0; len(guestNames)), функция возвращает исходный слайс без изменений. Иначе: сдвиньте хвост влево через copy, очистите последний элемент старой длины через clear и верните слайс длиной на 1 меньше.
Основная программа читает n, затем n строк, затем removeIndex. После вызова removeAtStable выведите итоговый слайс в одну строку (через пробел).
Требования:
• В решении должна быть реализована функция removeAtStable(guestNames []string, removeIndex int) []string, которая принимает слайс строк и индекс удаления и возвращает слайс (возможно, обновлённый по длине).
• Если removeIndex не входит в диапазон [0; len(guestNames)), функция должна вернуть исходный слайс без изменения содержимого и без изменения длины.
• При корректном removeIndex функция должна удалить элемент с сохранением порядка остальных элементов, сдвинув хвост влево ровно через copy(guestNames[removeIndex:], guestNames[removeIndex+1:]).
• После сдвига функция должна очистить последний элемент старой длины (тот слот в backing array, который стал “лишним” после удаления) с помощью clear на диапазоне, соответствующем этому последнему элементу.
• После успешного удаления функция должна вернуть слайс, у которого len уменьшен ровно на 1 (то есть возвращается новый “вид” слайса с обновлённым len, без создания нового массива).
*/

import (
	"bufio"
	"fmt"
	"os"
)

func removeAtStable(guestNames []string, removeIndex int) []string {
	if removeIndex < 0 || removeIndex >= len(guestNames) {
		return guestNames
	}

	oldLen := len(guestNames)

	// TODO: Реализуйте стабильное удаление элемента по индексу removeIndex (порядок остальных должен сохраниться).
	// TODO: Сдвиньте хвост влево с помощью copy (без append и без создания нового массива).
	// TODO: Очистите "освободившийся" последний слот в backing array через clear.
	// TODO: Верните слайс с len на 1 меньше (через реслайс), не создавая новый backing array.
	copy(guestNames[removeIndex:], guestNames[removeIndex+1:])

	// Временная упрощённая реализация: уменьшаем длину, удаляя последний элемент (это НЕ удаление по removeIndex).
	clear(guestNames[oldLen-1 : oldLen])
	return guestNames[:oldLen-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	guestNames := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &guestNames[i])
	}

	var removeIndex int
	fmt.Fscan(in, &removeIndex)

	guestNames = removeAtStable(guestNames, removeIndex)

	for i, name := range guestNames {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, name)
	}
	fmt.Fprintln(out)
}