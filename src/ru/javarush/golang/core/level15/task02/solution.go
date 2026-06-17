package main

/*
= Задача №2. Контактный черновик =
Вы пишете черновик телефонной книги “на один контакт” — просто чтобы проверить, как работают запись, перезапись и удаление в map. Сначала пользователь диктует имя и первый телефон, потом вспоминает, что номер был другой, и диктует второй вариант.
Программа читает три строки: contactName, firstPhone, secondPhone. Затем создаёт пустую карту, записывает firstPhone по ключу contactName, после этого перезаписывает номер на secondPhone. Дальше выведите итоговый номер и количество записей в карте.
После этого пользователь просит “стереть контакт”, и вы удаляете ключ contactName через delete, а затем печатаете новое количество записей.
Требования:
• Программа должна прочитать три строки (contactName, firstPhone, secondPhone) через fmt.Scan или fmt.Fscan.
• Программа должна создать пустую карту phonebook строго через make(map[string]string) с типом map[string]string.
• Программа должна сначала записать firstPhone по ключу contactName, затем перезаписать значение по тому же ключу на secondPhone, используя присваивание вида phonebook[contactName] = ....
• Программа должна вывести итоговый телефон (то есть secondPhone, сохранённый в карте по contactName) и количество записей в карте до удаления, используя только len(phonebook).
*/

import "fmt"

func main() {
	var contactName, firstPhone, secondPhone string
	fmt.Scan(&contactName, &firstPhone, &secondPhone)

	// Пустая телефонная книга на один контакт.
	phonebook := make(map[string]string)

	// TODO: Запишите firstPhone по ключу contactName в phonebook.
	// TODO: Перезапишите значение по ключу contactName на secondPhone.
	phonebook[contactName] = firstPhone
	phonebook[contactName] = secondPhone

	fmt.Println(phonebook[contactName])
	fmt.Println(len(phonebook))

	// TODO: Удалите контакт contactName из phonebook через delete(phonebook, contactName).
	delete(phonebook, contactName)

	fmt.Println(len(phonebook))
}