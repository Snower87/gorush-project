package main

/*
= Задача №20. Инспектор задач =
Вы делаете маленький "TaskInspector" для команды: одна цифра отвечает за статус задачи (new, in_progress, done), а второе число — за набор флагов (срочно, есть дедлайн, заблокировано).
Менеджеру нужно видеть это в понятном текстовом виде: статус — одним словом, а флаги — одной строкой через | и в фиксированном порядке.
Программа читает statusCode и flagsCode, затем печатает две строки: status=... и flags=....
Если статус неизвестен — unknown. Если ни один известный флаг не включён — none.
Требования:
• В программе должны быть объявлены два разных пользовательских типа: Status (на базе int) и TaskFlags (на базе uint), без использования “голых” int/uint для логики статуса/флагов.
• Должны быть объявлены константы StatusNew, StatusInProgress, StatusDone через iota для типа Status.
• Должны быть объявлены константы FlagUrgent, FlagHasDeadline, FlagBlocked для типа TaskFlags через выражения вида 1 << iota.
• Программа должна читать два числа statusCode и flagsCode через fmt.Scan(&statusCode, &flagsCode).
• Должна быть реализована функция statusText(s Status) string, которая использует switch и возвращает "unknown" в ветке default для неизвестного значения.
• Должна быть реализована функция flagsText(f TaskFlags) string, которая проверяет включённость флагов побитовыми операциями (f & FlagX) и сравнением результата с 0, формирует строку флагов в порядке urgent, затем deadline, затем blocked, соединяя включённые флаги символом '|' без пробелов; если не включён ни один известный флаг, функция должна вернуть "none".
*/

import "fmt"

// Status — пользовательский тип статуса задачи (не "голый" int).
type Status int

const (
	// Enum-значения через iota.
	StatusNew Status = iota
	StatusInProgress
	StatusDone
)

// TaskFlags — пользовательский тип набора флагов (не "голый" uint).
type TaskFlags uint

const (
	// Флаги через битовые сдвиги: 1, 2, 4, ...
	FlagUrgent TaskFlags = 1 << iota
	FlagHasDeadline
	FlagBlocked
)

func statusText(s Status) string {
	// TODO: Реализуйте преобразование статуса в текст через switch:
	// new / in_progress / done, а для неизвестных значений — "unknown".
	switch s {
	case StatusNew:
        return "new"
	case StatusInProgress:
        return "in_progress"
	case StatusDone:
        return "done"
    default:
        return "unknown"
	}

}

func flagsText(f TaskFlags) string {
	// TODO: Реализуйте преобразование флагов в текст:
	// проверьте включённость флагов побитовыми операциями (&) в порядке:
	// urgent, затем deadline, затем blocked.
	// Соединяйте включённые флаги символом '|' без пробелов.
	// Если не включён ни один известный флаг — верните "none".
	out := ""
	if f & FlagUrgent != 0 {
	    out += "urgent|"
	}
	if f & FlagHasDeadline != 0 {
	    out += "deadline|"
	}
	if f & FlagBlocked != 0 {
	    out += "blocked|"
	}
    if out == "" {
        return "none"
    }
    return out[:len(out) - 1]
}

func main() {
	// По требованию: читаем ровно два числа через fmt.Scan(&statusCode, &flagsCode).
	var statusCode int
	var flagsCode uint
	fmt.Scan(&statusCode, &flagsCode)

	s := Status(statusCode)
	f := TaskFlags(flagsCode)

	// По требованию: ровно две строки.
	fmt.Printf("status=%s\nflags=%s\n", statusText(s), flagsText(f))
}