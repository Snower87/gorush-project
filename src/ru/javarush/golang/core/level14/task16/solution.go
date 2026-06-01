package main

/*
= Задача №16. Байтовый отчёт =
Вы подключили модуль, который отдаёт данные как байты, а не как строки (например, кусок «сырых» данных из протокола). Разделитель внутри — |, вокруг элементов бывают пробелы, а иногда встречаются пустые элементы. Нужно аккуратно почистить вход и построить отчёт: и одну строку‑сводку, и многострочный список.
В программе задана переменная raw типа []byte со значением []byte(" alpha| beta | |gamma| "). Обрежьте пробелы по краям всего raw через bytes.TrimSpace, разбейте по []byte("|") через bytes.Split, у каждой части снова обрежьте пробелы, удалите пустые элементы, затем соберите line через bytes.Join с разделителем []byte(",").
После этого соберите многострочный report через bytes.Buffer в точном формате:
Items:
- alpha
- beta
- gamma
Line: alpha,beta,gamma
и выведите report как строку.
Требования:
• Программа должна содержать переменную raw типа []byte со значением []byte(" alpha| beta | |gamma| ") и не должна читать данные из stdin.
• Программа должна обрезать пробелы по краям всего raw с помощью bytes.TrimSpace, затем разбить результат через bytes.Split с разделителем, заданным строго как []byte("|").
• После Split программа должна для каждого элемента снова применить bytes.TrimSpace и удалить элементы, которые стали пустыми после обрезки пробелов.
• Программа должна собрать байтовую строку line из очищенных непустых элементов через bytes.Join, используя разделитель []byte(",") (результат должен соответствовать "alpha,beta,gamma").
• Программа должна собрать итоговый report с помощью bytes.Buffer в точном формате "Items:\n- alpha\n- beta\n- gamma\nLine: alpha,beta,gamma\n", используя WriteString для текстовых частей и Write для добавления байтовой строки line.
*/

import (
	"bytes"
	"fmt"
)

func main() {
	// Заданный "сырой" вход: байты, разделитель внутри '|', пробелы и пустые элементы возможны.
	raw := []byte("  alpha| beta | |gamma|  ")

	// Чистим общий ввод и режем по байтовому разделителю.
	trimmed := bytes.TrimSpace(raw)
	parts := bytes.Split(trimmed, []byte("|"))

    clean := make([][]byte, 0, len(parts))

	// TODO: Пройдитесь по parts, у каждого элемента сделайте bytes.TrimSpace,
	// TODO: удалите элементы, которые стали пустыми, и соберите итоговый слайс clean.
	for _, v := range parts {
	    v = bytes.TrimSpace(v)
	    if v == "" {
	        continue
	    }
        clean = append(clean, v)
	}

	// TODO: Соберите сводную байтовую строку line из clean через bytes.Join с разделителем []byte(",").
	var line []byte
	line = bytes.Join(clean, []byte(","))

	// Собираем многострочный отчёт через bytes.Buffer.
	var buf bytes.Buffer
	buf.WriteString("Items:\n")

	for _, item := range clean {
		buf.WriteString("- ")
		// TODO: Добавьте байты item и перевод строки '\n' через методы bytes.Buffer (не через конкатенацию строк).
		buf.Write(item)
		buf.WriteByte("\n")
	}

	buf.WriteString("Line: ")
	buf.Write(line)
	buf.WriteByte('\n')

	// Единичный финальный вывод.
	fmt.Print(buf.String())
}