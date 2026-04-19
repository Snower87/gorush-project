package main

/*
= Задача №20. Коробки доставки =
Вы помогаете службе доставки: есть общий вес посылки, ограничение на вес одной коробки и стоимость отправки одной коробки. Если коробок будет меньше, чем нужно, часть посылки “не поместится” — так нельзя. Поэтому по правилу количество коробок всегда округляется вверх math.Ceil. Но есть ловушка: если просто напечатать число без дробной части через "%.0f", это лишь визуальное округление, которое может ввести в заблуждение.
Программа читает totalWeight, maxBoxWeight и boxPrice (все float64). Считает rawBoxes = totalWeight / maxBoxWeight. Затем выводит “красивое” количество коробок как PrettyBoxes (печать rawBoxes через "%.0f", без math.Ceil), и отдельно считает boxesByRule = math.Ceil(rawBoxes) — столько коробок реально нужно по правилу. Итоговая стоимость total = boxesByRule * boxPrice считается именно по правилу, а не по красивому отображению.
Требования:
• Программа должна считать из stdin три значения типа float64: totalWeight, maxBoxWeight и boxPrice, используя fmt.Scan.
• Программа должна вычислить rawBoxes как totalWeight / maxBoxWeight, сохраняя результат в переменной типа float64.
• Программа должна вывести PrettyBoxes как rawBoxes, напечатанный через формат %.0f, без применения math.Ceil и без преобразования в целый тип.
• Программа должна импортировать пакет math и вычислить boxesByRule как math.Ceil(rawBoxes), после чего вывести BoxesByRule именно из результата math.Ceil(rawBoxes).
• Программа должна вычислить итоговую стоимость total как boxesByRule * boxPrice, используя boxesByRule (результат math.Ceil), а не PrettyBoxes или любое “красивое” округление.
• Программа должна вывести Total с двумя знаками после точки, используя формат %.2f.
*/

import (
	"fmt"
	"math"
)

func main() {
	var totalWeight, maxBoxWeight, boxPrice float64
	fmt.Scan(&totalWeight, &maxBoxWeight, &boxPrice)

	rawBoxes := totalWeight / maxBoxWeight

	// TODO: Выведите PrettyBoxes: распечатайте rawBoxes через формат "%.0f" (без math.Ceil и без перевода в целый тип)
	fmt.Printf("%.0f\n", rawBoxes)

	// TODO: Исправьте расчёт boxesByRule: по правилу нужно округлять вверх (math.Ceil), а затем печатать через "%.0f"
	boxesByRule := math.Ceil(rawBoxes)
	fmt.Printf("%.0f\n", boxesByRule)

	total := boxesByRule * boxPrice

	// TODO: Выведите Total с двумя знаками после точки через формат "%.2f"
	fmt.Printf("%.2f\n", total)
}