package main

/*
= Задача №19. Безопасный просмотр =
Вы пишете функцию "показать первые два элемента", как будто это превью в интерфейсе. Но вы не хотите, чтобы кто-то потом сделал append к этому превью и случайно испортил исходные данные в общем буфере. Поэтому превью должно быть именно “видом”, но с намеренно ограниченной ёмкостью.
Реализуйте функцию viewFirstTwo(all []string) []string, которая возвращает вид на первые два элемента слайса, но запрещает расширение этого вида через append в общий backing array.
Если длина all меньше 2, функция должна возвращать вид на весь all, тоже с ограниченным cap (то есть cap должен быть равен len). Если len(all) >= 2, функция должна возвращать all[:2:2].
В main создайте all как make([]string, 4, 4) со значениями "A", "B", "C", "D". Получите v := viewFirstTwo(all). Затем сделайте grown := append(v, "X") и выведите all и grown. После этого выполните v[0] = "Z" и выведите all и v, чтобы увидеть разницу между append (не должен трогать all) и изменением по индексу (должно отразиться в all).
// Примерный шаблон для main:
all := make([]string, 4, 4)
all[0], all[1], all[2], all[3] = "A", "B", "C", "D"
v := viewFirstTwo(all)
grown := append(v, "X")
fmt.Println(all)
fmt.Println(grown)
v[0] = "Z"
fmt.Println(all)
fmt.Println(v)
Требования:
• Нужно реализовать отдельную функцию viewFirstTwo(all []string) []string, которая возвращает “вид” (sub-slice) с намеренно ограниченной ёмкостью: при len(all) >= 2 возвращать именно all[:2:2], а при len(all) < 2 возвращать именно all[:len(all):len(all)] (то есть cap результата должен быть равен len результата).
• В main нужно создать all через make([]string, 4, 4) и заполнить его значениями "A", "B", "C", "D" в этом порядке.
• В main нужно получить v := viewFirstTwo(all), затем создать grown только одной операцией grown := append(v, "X") (без append от all и без дополнительных копирований/пересозданий grown).
• После выполнения grown := append(v, "X") требуется вывести all и grown так, чтобы было видно, что all остался [A B C D], а grown содержит добавленный элемент.
• После append нужно выполнить v[0] = "Z" и затем вывести all и v так, чтобы было видно, что all стал [Z B C D] (то есть изменения через индекс в view отражаются в исходном backing array).
*/

import "fmt"

// Слайс — это view (ptr+len+cap) на backing array.
// Если намеренно ограничить cap, то append не сможет "расти" в общий массив
// и будет вынужден выделить новый буфер.
func viewFirstTwo(all []string) []string {
	// TODO: Реализуйте безопасный "view" на первые два элемента (или на весь all, если len < 2).
	// TODO: Важно намеренно ограничить ёмкость результата так, чтобы cap(result) == len(result),
	// TODO: и append к result не мог изменить исходный backing array.
	//
	// Ниже оставлена упрощённая версия, которая НЕ выполняет требование по ограничению cap.
	if len(all) < 2 {
		return all[:len(all):len(all)]
	}
	return all[:2:2]
}

func main() {
	all := make([]string, 4, 4)
	all[0] = "A"
	all[1] = "B"
	all[2] = "C"
	all[3] = "D"

	v := viewFirstTwo(all)

	fmt.Printf("v(initial)=%v len=%d cap=%d\n", v, len(v), cap(v))
	fmt.Printf("all(initial)=%v\n", all)

	grown := append(v, "X")

	fmt.Printf("all(after append)=%v\n", all)
	fmt.Printf("grown(after append)=%v\n", grown)

	// Изменение по индексу — это всё тот же backing array, поэтому all изменится.
	v[0] = "Z"

	fmt.Printf("all(after edit)=%v\n", all)
	fmt.Printf("v(after edit)=%v len=%d cap=%d\n", v, len(v), cap(v))
}