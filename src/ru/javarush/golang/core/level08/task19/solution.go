package main

/*
= Задача №19. Права файла =
В вашем мини‑файловом менеджере права доступа упакованы в одно число permCode (битовая маска). Один бит означает “можно читать”, другой — “можно писать”, третий — “можно исполнять”, четвёртый — “файл скрытый”. При этом в маске могут случайно встретиться и другие биты — программа не должна ломаться из‑за этого.
Напишите программу, которая читает permCode, проверяет каждый из четырёх известных флагов через helper‑функцию и печатает отчёт в четырёх строках строго заданного вида.
Требования:
• В программе должен быть объявлен тип `FilePerm` на базе `uint`, и все операции с правами должны выполняться через этот тип.
• Должны быть объявлены 4 константы-флага через `1 << iota` в строгом порядке: `PermRead`, `PermWrite`, `PermExec`, `PermHidden`.
• Должна быть реализована функция `hasPerm(all, p FilePerm) bool`, которая возвращает `true`, если флаг установлен, используя проверку `all&p != 0`.
• Программа должна читать число в переменную `permCode` через `fmt.Scan(&permCode)` и затем преобразовать его в `FilePerm(permCode)` перед проверками флагов.
• Программа должна вывести ровно 4 строки строго в порядке `read`, `write`, `exec`, `hidden`, каждая в формате `name=<true/false>` без дополнительных строк или текста.
*/

import "fmt"

// FilePerm — пользовательский тип для работы с битовой маской прав.
type FilePerm uint

const (
	// TODO: Объявите флаги прав через 1 << iota в строгом порядке: read, write, exec, hidden.
	PermRead FilePerm = 1 << iota
	PermWrite
	PermExec
	PermHidden
)

// hasPerm проверяет, установлен ли флаг p в маске all.
func hasPerm(all, p FilePerm) bool {
	// TODO: Реализуйте проверку установленного флага p в маске all через побитовые операции.
	// Важно: неизвестные биты в all не должны ломать результат.
	return all & p != 0
}

func main() {
	var permCode uint
	fmt.Scan(&permCode)

	// Все проверки делаем через FilePerm.
	all := FilePerm(permCode)

	fmt.Printf("read=%t\n", hasPerm(all, PermRead))
	fmt.Printf("write=%t\n", hasPerm(all, PermWrite))
	fmt.Printf("exec=%t\n", hasPerm(all, PermExec))
	fmt.Printf("hidden=%t\n", hasPerm(all, PermHidden))
}