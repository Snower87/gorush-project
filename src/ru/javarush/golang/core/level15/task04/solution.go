package main

/*
= Задача №4. Команды телефонника =
Вы собираете мини‑CLI для телефонной книги, как будто это консольная утилита: добавили номер, посмотрели, удалили, проверили размер. Всё хранится в map[string]string, где ключ — имя, значение — телефон. Телефон всегда непустая строка, а вот отсутствие записи надо уметь показывать сообщением.
Поддерживаются команды set, get, del, len. В первой строке дано число q — сколько команд дальше придёт. Далее идут q команд:
- set name phone — записывает (или обновляет) телефон,
- get name — выводит телефон, а если записи нет, выводит MISSING,
- del name — удаляет запись (повторное удаление допускается),
- len — выводит текущее количество записей.
Карту начните с var book map[string]string (то есть nil‑map). Инициализируйте её только при первой команде set.
Требования:
• Программа должна прочитать целое число q, затем обработать ровно q команд, читая их из stdin по одной через fmt.Scan/fmt.Fscan.
• Программа должна хранить телефонную книгу в переменной book типа map[string]string, объявленной как var book map[string]string (изначально nil), и создавать её через make(map[string]string) только при первой команде set (проверка book == nil).
• При команде set name phone программа должна (при необходимости сначала инициализировать book) записать значение phone по ключу name, перезаписывая старое значение, если ключ уже был.
• При команде get name программа должна получить значение через book[name] и, если результат равен пустой строке "", вывести MISSING, иначе вывести телефон; вывод должен быть в stdout с новой строки.
• При команде del name программа должна вызывать delete(book, name) без предварительных проверок существования ключа (повторное удаление допускается и не должно ломать программу).
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)

	// По условию карта должна начинаться как nil-map.
	var book map[string]string

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "set":
			var name, phone string
			fmt.Fscan(in, &name, &phone)

			// TODO: Реализуйте команду set: при первом set инициализируйте book через make (ленивая инициализация),
			if book == nil {
			    book = make(map[string]string)
			}
			// TODO: затем сохраните/обновите phone по ключу name.
			book[name] = phone

		case "get":
			var name string
			fmt.Fscan(in, &name)

			// TODO: Реализуйте команду get: получите значение через book[name],
			// TODO: если результат равен пустой строке "", выведите MISSING, иначе выведите телефон.
			if book[name] == "" {
			    fmt.Fprintln(out, "MISSING")
			} else {
			    fmt.Fprintln(out, book[name])
			}

		case "del":
			var name string
			fmt.Fscan(in, &name)

			// TODO: Реализуйте команду del: удалите запись через delete(book, name) без дополнительных проверок.
			delete(book, name)

		case "len":
			// TODO: Реализуйте команду len: выведите текущее количество записей как len(book)
			// TODO: (в том числе для nil-map должно получаться 0).

			fmt.Fprintln(out, len(book))
		}
	}
}