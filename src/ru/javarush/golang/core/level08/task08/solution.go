package main

/*
= Задача №8. Канбан доска =
У вас мини‑канбан в консоли: в систему прилетают три «сырых» числа — идентификатор задачи, код статуса и сколько дней она делалась. Но вы не хотите мешать эти смыслы в одном int, поэтому вводите отдельные типы: TaskID, Status, Days.
Программа конвертирует входные значения в смысловые типы и печатает понятный статус. Если статус неизвестен, важно вывести именно исходный сырой код, чтобы человек мог быстро найти проблему в интеграции.
Требования:
• Программа должна считать через fmt.Scan три значения типа int: rawID, rawStatus и rawDays.
• В программе должны быть объявлены пользовательские типы type TaskID int, type Status int, type Days int, а также две Status-константы с явными значениями (без iota): StatusNew = 0 и StatusDone = 1.
• Программа должна явно сконвертировать сырые int в смысловые значения: id := TaskID(rawID), status := Status(rawStatus), days := Days(rawDays) (без неявных преобразований).
• Логика выбора ветки должна быть реализована через if / else if / else, при этом status должен сравниваться только с StatusNew и StatusDone (не с числами и не с rawStatus).
• В ветке done для вывода длительности должно использоваться значение days типа Days без обратной конверсии в int; в ветке invalid для вывода кода статуса должна использоваться переменная rawStatus типа int.
*/

import "fmt"

// Смысловые типы, чтобы не мешать разные значения в одном int.
type TaskID int
type Status int
type Days int

const (
	StatusNew  Status = 0
	StatusDone Status = 1
)

func main() {
	var rawID, rawStatus, rawDays int
	fmt.Scan(&rawID, &rawStatus, &rawDays)

	// Явные конверсии из "сырого" int в доменные типы.
	id := TaskID(rawID)
	status := Status(rawStatus)
	days := Days(rawDays)

    if status == StatusNew {
        fmt.Printf("task #%d: new\n", id)
    } else if status == StatusDone {
        fmt.Printf("task #%d: done in %d days\n", id, days)
    } else {
        fmt.Printf("task #%d: invalid status (%d)\n", id, rawStatus)
    }
}