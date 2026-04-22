package main

/*
= Задача №17. Светофор переводчик =
Вы подключились к контроллеру перекрёстка: он присылает короткий числовой код цвета, а оператору нужно видеть понятный текст — red, yellow или green. Иногда контроллер может прислать мусорный код, и тогда важно честно показать unknown, а не угадывать.
Напишите программу, которая читает lightCode, превращает его в ваш enum-тип Light, прогоняет через функцию расшифровки и печатает ровно одну строку с результатом.
Требования:
• Программа должна объявить пользовательский тип `type Light int`.
• Программа должна объявить константы `LightRed`, `LightYellow`, `LightGreen` через `iota` строго в порядке red, yellow, green, чтобы значения соответствовали `0, 1, 2`.
• Программа должна реализовать функцию `lightText(l Light) string`, которая возвращает ровно одну из строк: `red`, `yellow`, `green` или `unknown` в зависимости от значения `l`.
• В логике определения текста (и в других проверках по коду цвета) нельзя использовать числовые литералы `0/1/2`; сравнения и ветвления должны опираться только на `LightRed/LightYellow/LightGreen`.
• Программа должна прочитать `lightCode` через `fmt.Scan`, преобразовать его в `Light(lightCode)` в `main`, и вывести через `fmt.Println` ровно одну строку — результат вызова `lightText(...)`.
*/

import "fmt"

// Light — наш enum-тип для кода светофора.
type Light int

const (
	LightRed Light = iota
	LightYellow
	LightGreen
)

// lightText превращает код в человекочитаемый текст.
func lightText(l Light) string {
	// TODO: Реализуйте расшифровку кода светофора в текст.
	// TODO: Функция должна возвращать ровно одну из строк: "red", "yellow", "green" или "unknown".
	// TODO: В логике нельзя использовать числовые литералы 0/1/2 — сравнивайте только с LightRed/LightYellow/LightGreen.
	switch l {
    case LightRed:
        return "red"
    case LightYellow:
        return "yellow"
    case LightGreen:
        return "green"
    default:
        return "unknown"
	}
}

func main() {
	var lightCode int
	fmt.Scan(&lightCode)

	// По условию: преобразование делаем в main.
	fmt.Println(lightText(Light(lightCode)))
}