package main

/*
= Задача №19. Скидка с округлением =
Вы делаете расчёт скидки для карточки лояльности. Сначала программа считает цену после скидки “как получится” (сырое значение), а затем обязана округлить её до 2 знаков именно как новое значение, чтобы дальше в системе работало уже округлённое число, а не просто красивый вывод.
Программа читает originalPrice (тип float64) и discountPercent (тип int). Считает discount = originalPrice * float64(discountPercent) / 100 и afterDiscount = originalPrice - discount. Затем создаёт отдельную переменную roundedPrice, где выполняет математическое округление до 2 знаков через math.Round(afterDiscount*100) / 100. В конце печатает три строки: сырое значение, сырое округлённое значение, и то же округлённое значение “для пользователя” с "%.2f".
Требования:
• Программа должна считать из stdin через fmt.Scan два значения: originalPrice типа float64 и discountPercent типа int.
• Программа должна вычислить discount по формуле discount = originalPrice * float64(discountPercent) / 100 и afterDiscount = originalPrice - discount.
• Программа должна создать отдельную переменную roundedPrice и выполнить округление afterDiscount до 2 знаков как новое числовое значение через math.Round(afterDiscount*100) / 100 с использованием пакета math и функции math.Round.
• Программа должна вывести ровно три строки через fmt.Printf: (1) afterDiscount в формате %.17f, (2) roundedPrice в формате %.17f, (3) roundedPrice в формате %.2f (пользовательский вид).
• Форматирование вывода через %.2f должно применяться только для отображения и не должно заменять вычисление roundedPrice (нельзя получать “округлённое значение” повторным форматированием afterDiscount вместо math.Round).
*/

import (
	"fmt"
	"math"
)

func main() {
	var originalPrice float64
	var discountPercent int
	fmt.Scan(&originalPrice, &discountPercent)

	// TODO: Посчитайте discount по формуле из условия задачи (процент от originalPrice).
	//          float64 * float64 / 100
	discount := originalPrice * float64(discountPercent) / 100

	// TODO: Посчитайте afterDiscount как цену после применения скидки.
	afterDiscount := originalPrice - discount

	// TODO: Округлите afterDiscount до 2 знаков как новое числовое значение (не через форматирование вывода).
	roundedPrice := math.Round(afterDiscount * 100) / 100

	fmt.Printf("%.17f\n", afterDiscount)
	fmt.Printf("%.17f\n", roundedPrice)
	fmt.Printf("%.2f\n", roundedPrice)
}