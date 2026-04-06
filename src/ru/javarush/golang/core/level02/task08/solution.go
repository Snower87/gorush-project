package main

import "fmt"

func main() {
	// TODO: Задайте строковые переменные firstName, lastName, role значениями из условия задачи.
	firstName := "Ivan"
	lastName := "Petrov"
	role := "Junior Gopher"

	// TODO: Соберите fullName конкатенацией (через +), добавив ровно один пробел между именем и фамилией.
	fullName := firstName + " " + lastName

	// TODO: Выведите ровно 4 строки через fmt.Println в точном формате из условия (включая пробелы после двоеточия).
	fmt.Println("First name:", firstName)
	fmt.Println("Last name:", lastName)
	fmt.Println("Role:", role)

	// TODO: Используйте len(fullName) как число при выводе (не склеивайте строку и число через +).
	fmt.Println("Full name length:", len(fullName))

	_ = fullName
}