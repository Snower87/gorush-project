package main

/*
= Задача №20. Онлайн и лидер =
У вас есть журнал событий по пользователям, как в чате или игре: кто-то зашёл (IN), кто-то вышел (OUT). После обработки всех событий нужно ответить на два вопроса: кто остался онлайн прямо сейчас, и кто чаще всего заходил (то есть у кого больше всего событий IN).
Каждое событие — пара: action и userName. action бывает IN или OUT, userName — строка без пробелов.
Ввод: число m, затем m пар (action, userName).
Вывод:
- В первой строке число onlineCount — сколько пользователей сейчас онлайн.
- Далее onlineCount строк с именами онлайн‑пользователей в отсортированном порядке.
- Затем строка maxIn — максимальное число входов среди всех пользователей.
- Затем имена всех пользователей с числом входов = maxIn, в отсортированном порядке (по одному имени в строке).
Для онлайн‑списка используйте set на map[string]struct{}, для входов — счётчик map[string]int.
Требования:
• Программа должна прочитать целое число m, затем ровно m пар значений (action, userName), где action — строка IN или OUT, а userName — строка без пробелов.
• Текущее множество онлайн-пользователей должно храниться в map[string]struct{} и использоваться как set (наличие ключа означает “онлайн”).
• При действии IN программа должна добавить userName в online-set и увеличить счётчик входов userName в map[string]int; при действии OUT программа должна удалить userName из online-set с помощью delete независимо от того, был ли ключ в map.
• После обработки всех событий программа должна вывести onlineCount (количество ключей в online-set), затем вывести имена всех онлайн-пользователей по одному в строке в лексикографическом порядке (через сбор ключей в слайс и сортировку).
• Программа должна вычислить maxIn как максимум по значениям map[string]int одним проходом по map, затем вывести maxIn и вывести всех пользователей, у которых число входов равно maxIn, в лексикографическом порядке по одному имени в строке.
*/

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scan(&m)

	// online — set через map[string]struct{}
	online := map[string]struct{}{}

	// inCount — счётчик входов (IN) по пользователям
	inCount := map[string]int{}

	for i := 0; i < m; i++ {
		var action, userName string
		fmt.Scan(&action, &userName)

		if action == "IN" {
			// TODO: Обработайте вход пользователя: добавьте userName в online-set и увеличьте счётчик входов в inCount.
			online[userName] = struct{}{}
			inCount[userName]++
		} else if action == "OUT" {
			// TODO: Обработайте выход пользователя: удалите userName из online-set через delete (даже если ключа нет).
			delete(online, userName)
		}
	}

	// TODO: Сформируйте список онлайн-пользователей из online (ключи map), отсортируйте и выведите:
	// - onlineCount
	// - onlineCount строк с именами пользователей (лексикографически)
	//
	// Временный код ниже оставлен, чтобы шаблон был минимально запускаемым.
	onlineUsers := make([]string, 0, len(online))
	for name := range online {
	    onlineUsers = append(onlineUsers, name)
	}
	sort.Strings(onlineUsers)

	fmt.Println(len(onlineUsers))
	// TODO: Выведите имена онлайн-пользователей (по одному в строке) после строки с onlineCount.
	for _, user := range onlineUsers {
	    fmt.Println(user)
	}

	// TODO: Найдите maxIn как максимум по значениям inCount и выведите maxIn.
	maxIn := 0
	for _, k := range inCount {
	    if k > maxIn { maxIn = k }
	}
	fmt.Println(maxIn)

	// TODO: Сформируйте список лидеров (всех пользователей, у кого входов == maxIn),
	// отсортируйте и выведите по одному имени в строке.
	leaders := make([]string, 0)
	for v, k := range inCount {
	    if k == maxIn {
	        leaders = append(leaders, v)
	    }
	}
	sort.Strings(leaders)
	for ind, leader := range leaders {
	    fmt.Println(leader)
	}
}