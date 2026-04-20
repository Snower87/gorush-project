package main

/*
= Задача №7. Статус задачи =
Вы пишете консольный трекер задач. На вход вам дают taskID и числовой statusCode. Внутри программы вы хотите работать не с «голыми числами», а с типом Status и именованными константами, чтобы код читался как человеческий.
После чтения вы преобразуете statusCode в Status и печатаете ровно одну строку в стиле: task #<id>: .... Если код неизвестен, нужно честно показать пользователю сырой код в скобках.
Требования:
• Программа должна прочитать через fmt.Scan ровно два целых числа типа int: taskID и statusCode.
• В программе должен быть объявлен пользовательский тип на базе числа: type Status int.
• Программа должна объявить константы типа Status: StatusNew = 0, StatusInProgress = 1, StatusDone = 2.
• После чтения statusCode (int) программа должна явно конвертировать его в Status через Status(statusCode), при этом исходный statusCode должен оставаться доступным отдельно для вывода в unknown-ветке.
• Логика определения текста статуса должна быть реализована через if / else if / else (без switch) и сравнивать значение Status только с именованными константами StatusNew/StatusInProgress/StatusDone (не с числами 0/1/2).
*/

import "fmt"

// Status — пользовательский тип вместо "голых чисел".
type Status int

const (
	StatusNew        Status = 0
	StatusInProgress Status = 1
	StatusDone       Status = 2
)

func main() {
	var taskID, statusCode int
	fmt.Scan(&taskID, &statusCode)

	// Явная конверсия в Status, но исходный statusCode (int) держим отдельно для unknown-ветки.
	status := Status(statusCode)

	var text string

	// TODO: Реализуйте выбор текста статуса через if / else if / else:
	// TODO: сравнивайте status только с StatusNew / StatusInProgress / StatusDone (не с числами),
	// TODO: а для неизвестного кода формируйте "unknown status (<statusCode>)".
	if status == StatusNew {
		// TODO: Установите корректный текст для статуса "new".
		text = "new"
	} else if status == StatusInProgress {
		// TODO: Установите корректный текст для статуса "in progress".
		text = "in progress"
	} else if status == StatusDone {
		// TODO: Установите корректный текст для статуса "done".
		text = "done"
	} else {
		text = fmt.Sprintf("unknown status (%d)", statusCode)
	}

	fmt.Printf("task #%d: %s\n", taskID, text)
}