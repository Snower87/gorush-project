package main

/*
= Задача №11. Права редактора =
Вы настраиваете доступ к панели управления контентом. Пользователь может редактировать, только если он не заблокирован,
и при этом его роль — admin или editor. Важно не перепутать приоритеты операторов, поэтому часть с || нужно явно взять в скобки — так код читается без двусмысленностей.
Программа читает role (строка) и bannedStr (строка "yes"/"no"). Затем вычисляет:
isBanned := bannedStr == "yes"
isAdmin := role == "admin"
isEditor := role == "editor"
Разрешение на редактирование: (!isBanned) && (isAdmin || isEditor). Если условие истинно — вывести EDIT, иначе — READONLY.
• Программа должна прочитать из stdin ровно два значения (role и bannedStr) как строки, используя fmt.Scan.
• Программа должна вычислить isBanned как (bannedStr == "yes"), isAdmin как (role == "admin") и isEditor как (role == "editor").
• Разрешение на редактирование должно определяться выражением (!isBanned) && (isAdmin || isEditor), обязательно используя операторы !, && и ||.
• В итоговом логическом выражении должны быть скобки вокруг (isAdmin || isEditor), чтобы приоритет операций был явно задан.
*/

import "fmt"

func main() {
	var role string
	var bannedStr string

	// По условию читаем ровно два значения через fmt.Scan
	fmt.Scan(&role, &bannedStr)

	// TODO: Вычислите булевые флаги isBanned, isAdmin, isEditor по входным строкам (через сравнение строк).
	isBanned := bannedStr == "yes"
	isAdmin := role == "admin"
	isEditor := role == "editor"

	// TODO: Соберите итоговое логическое условие для права редактирования, используя !, &&, ||.
	// TODO: Явно поставьте скобки вокруг части с ||, чтобы приоритет читался однозначно.

	canEdit := (!isBanned) && (isAdmin || isEditor)

	// Один if/else без вложенности
	if canEdit {
		fmt.Print("EDIT")
	} else {
		fmt.Print("READONLY")
	}
}