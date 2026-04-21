package main

/*
= Задача №10. Два счётчика =
Вы пишете маленький «пульт управления» для сервисов. В одном месте у вас команды (старт/стоп/рестарт), а в другом — приоритеты (низкий/высокий). И важно, чтобы оба набора нумерации начинались с нуля независимо друг от друга: это как два отдельных табло, которые не должны влиять друг на друга.
Поэтому вы делаете два разных const-блока, каждый со своим iota, и печатаете значения в двух строках.
Требования:
• В программе должны быть объявлены два отдельных const-блока, и в каждом из них должен использоваться iota.
• В каждом из двух const-блоков значения констант должны начинаться с 0 (то есть iota в каждом блоке должен давать 0 для первой константы этого блока).
• Должны быть объявлены типы Command и Priority как пользовательские типы на базе int (например, type Command int и type Priority int).
• В первом const-блоке должны быть объявлены константы CmdStart, CmdStop, CmdRestart с типом Command и с последовательными значениями от 0.
• Во втором const-блоке должны быть объявлены константы PrLow, PrHigh с типом Priority и с последовательными значениями от 0.
*/

import "fmt"

type Command int
type Priority int

const (
	// TODO: Описать перечисление команд через iota в отдельном const-блоке:
	// TODO: CmdStart=0, CmdStop=1, CmdRestart=2 (тип Command).
	CmdStart   Command = iota
	CmdStop    Command = iota
	CmdRestart Command = iota
)

const (
	// TODO: Описать перечисление приоритетов через iota во втором (независимом) const-блоке:
	// TODO: PrLow=0, PrHigh=1 (тип Priority).
	PrLow  Priority = iota
	PrHigh Priority
)

func main() {
	fmt.Println(CmdStart, CmdStop, CmdRestart)
	fmt.Println(PrLow, PrHigh)
}