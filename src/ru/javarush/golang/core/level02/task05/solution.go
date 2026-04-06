package main

import "fmt"

func main() {
	// Имя пользователя задано заранее по условию
	var userName string = "Ann"

	// TODO: Сформируйте строку greeting с помощью конкатенации (+) так,
	// TODO: чтобы итоговый вывод точно соответствовал условию задачи.
	var greeting string = "Hello, " + userName + "!"

	// Выводим ровно одну строку
	fmt.Println(greeting)
}