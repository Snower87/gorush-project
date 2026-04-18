package main

/*
= Задача №4. Свободные слоты =
Вы настраиваете лимиты в сервисе: есть общий объём хранилища и сколько уже занято. По смыслу эти числа не должны быть отрицательными, поэтому вы храните их в uint. Но есть подвох: если занято внезапно больше вместимости, вычитать нельзя — в uint это превратится в огромное число и сломает отчёт.
Программа читает capacitySlots и usedSlots как uint. Если usedSlots > capacitySlots, она выводит две строки: сначала error, а затем 0 (uint) (как безопасный “остаток”). Если всё корректно, программа выводит одну строку: сколько свободно и какой у этого значения тип.
Требования:
• Программа должна прочитать `capacitySlots` и `usedSlots` через `fmt.Scan` в переменные типа `uint`.
• Перед вычислением свободных слотов программа должна сравнить значения и определить, что `usedSlots > capacitySlots` — это ошибка.
• Если `usedSlots <= capacitySlots`, программа должна вычислить `freeSlots = capacitySlots - usedSlots`, выполняя вычитание строго в типе `uint`.
• Если `usedSlots > capacitySlots`, программа должна вывести ровно две строки: сначала `error`, затем строку `0 (uint)`, где вторая строка печатается через формат `"%v (%T)\n"`. Если `usedSlots <= capacitySlots`, программа должна вывести ровно одну строку: значение `freeSlots` и его тип (`uint`) в формате `"%v (%T)\n"`
*/

import "fmt"

func main() {
	var capacitySlots uint
	var usedSlots uint

	fmt.Scan(&capacitySlots, &usedSlots)

	// TODO: Добавьте проверку, что usedSlots > capacitySlots — это ошибка.
	// TODO: В ветке ошибки нельзя делать вычитание в uint (иначе будет underflow).
	// TODO: В ветке ошибки нужно вывести две строки: сначала "error", затем "0 (uint)" через fmt.Printf("%v (%T)\n", ...).
	if usedSlots > capacitySlots {
	    fmt.Println("error")
	    fmt.Printf("%v (%T)\n", uint(0), uint(0))
	    return
	}
	freeSlots := capacitySlots - usedSlots

	// TODO: Если ошибки нет, нужно вывести ровно одну строку: freeSlots и его тип (uint) через fmt.Printf("%v (%T)\n", ...).
	fmt.Printf("%v (%T)\n", freeSlots, freeSlots)
}