package main

/*
= Задача №10. Две сметы =
Вы считаете стоимость работ для небольшого проекта и видите, как легко ошибиться из‑за приоритета операций: без скобок калькулятор умножит раньше, а со скобками — сначала сложит. Чтобы никто в команде не спорил “как правильно”, вы печатаете оба варианта.
Программа читает три целых числа baseCost, extraUnits и multiplier. Затем вычисляет estimateA = baseCost + extraUnits*multiplier и estimateB = (baseCost + extraUnits) * multiplier. Оба результата нужно вывести в одной строке через пробел — без подписей и лишнего текста.
Требования:
• Программа должна импортировать пакет fmt и использовать его для ввода/вывода.
• Программа должна считать три целых числа baseCost, extraUnits и multiplier через fmt.Scan(&baseCost, &extraUnits, &multiplier).
• Переменные baseCost, extraUnits, multiplier, estimateA и estimateB должны иметь целочисленный тип (int или int64).
• Значение estimateA должно вычисляться выражением baseCost + extraUnits*multiplier (без добавления скобок, полагаясь на приоритет операций).
• Значение estimateB должно вычисляться выражением (baseCost + extraUnits) * multiplier (со скобками вокруг суммы).
*/

import "fmt"

func main() {
	var baseCost, extraUnits, multiplier int64
	fmt.Scan(&baseCost, &extraUnits, &multiplier)

	var estimateA int64
	var estimateB int64

	// TODO: Вычислите две оценки стоимости по условию задачи и присвойте их переменным estimateA и estimateB.
	estimateA = baseCost + extraUnits*multiplier
	estimateB = (baseCost + extraUnits) * multiplier

	fmt.Print(estimateA, " ", estimateB)
}