package main

/*
= Задача №3. Очки с нуля =
Вы делаете прототип табло для маленькой игры: игроки получают (или теряют) очки, а ваша программа аккуратно всё суммирует. Нюанс в том, что карта очков может начать жизнь как nil, и ваша функция должна уметь “оживлять” её сама — иначе первое же обновление всё сломает.
Реализуйте функцию addScore(scores map[string]int, playerName string, delta int) map[string]int. Функция должна прибавлять delta к текущему значению scores[playerName]. Если scores равна nil, внутри функции создайте карту через make, затем выполните обновление. Функция обязательно возвращает карту (это важно, когда на вход пришла nil).
В main объявите var scores map[string]int (то есть nil-map). Программа читает число n, затем n пар (playerName и delta) и для каждой пары вызывает addScore. После этого читается строка queryName, и программа выводит на двух строках: итоговые очки queryName и общее количество игроков (len карты).
Требования:
• Функция должна быть реализована с сигнатурой addScore(scores map[string]int, playerName string, delta int) map[string]int и обязательно возвращать map[string]int (в т.ч. когда на вход пришла nil-карта).
• Внутри addScore должна быть проверка на nil и “оживление” карты строго в виде if scores == nil { scores = make(map[string]int) }.
• Функция addScore должна прибавлять delta к текущему значению по ключу playerName через выражение scores[playerName] = scores[playerName] + delta.
• В main карта должна быть объявлена как var scores map[string]int (без make), а при каждом вызове addScore результат должен присваиваться обратно в scores, чтобы изменения сохранялись даже если карта была nil.
• Программа должна прочитать n, затем в цикле for выполнить чтение n пар (playerName и delta) и для каждой пары вызвать addScore для обновления карты.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func addScore(scores map[string]int, playerName string, delta int) map[string]int {
	// TODO: Реализуйте обновление карты очков:
	// TODO: 1) Если scores == nil — инициализируйте карту строго через make(map[string]int).
	if scores == nil {
	    scores = make(map[string]int)
	}
	// TODO: 2) Прибавьте delta к текущему значению scores[playerName] строго через выражение:
	// TODO:    scores[playerName] = scores[playerName] + delta
	scores[playerName] = scores[playerName] + delta
	// TODO: 3) Обязательно верните (возможно, новую) карту scores.
	return scores
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var scores map[string]int // важно: nil-map, без make

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var playerName string
		var delta int
		fmt.Fscan(in, &playerName, &delta)

		// Важно присваивать результат обратно: карта могла быть nil.
		scores = addScore(scores, playerName, delta)
	}

	var queryName string
	fmt.Fscan(in, &queryName)

	fmt.Fprintln(out, scores[queryName])
	fmt.Fprintln(out, len(scores))
}