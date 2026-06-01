package main

/*
= Задача №20. След первой разницы =
Вы делаете автопроверку ответов: ожидаемая строка (expected) и фактическая (actual) часто отличаются не буквами, а «мусором» по краям — пробелами или переносами строк. Чтобы не спорить вслепую, вы хотите отчёт: как строки выглядели сырыми, как выглядят после TrimSpace, и в каком месте начинается первое отличие.
Программа читает две строки из stdin (сначала expected, затем actual, обе через ReadString('\n')). Для каждой строки сделайте два варианта: Raw (как прочитали) и Norm (после strings.TrimSpace). Если expectedNorm равна actualNorm, выведите ровно одну строку match.
Если не равны, выведите диагностический отчёт ровно из пяти строк в указанном порядке:
expectedRaw=<%q>
actualRaw=<%q>
expectedNorm=<%q>
actualNorm=<%q>
firstDiff=<описание первого отличия>
firstDiff должен описывать первое отличие на уровне байтов в нормализованных строках: индекс, байт expected (или -1 если строка закончилась), байт actual (или -1).
Требования:
• Программа должна прочитать ровно две строки из stdin в порядке expected затем actual, используя bufio.Reader и метод ReadString('\n').
• Для каждой входной строки программа должна сохранить два отдельные значения: Raw (как прочитано из ReadString) и Norm (результат нормализации), то есть всего четыре отдельные переменные expectedRaw, actualRaw, expectedNorm, actualNorm.
• Нормализованные строки expectedNorm и actualNorm должны получаться строго через strings.TrimSpace от соответствующих raw-строк.
• Программа должна сравнивать на равенство именно expectedNorm и actualNorm; если они равны — вывести ровно одну строку: match.
• Если expectedNorm и actualNorm не равны, программа должна вывести ровно пять строк и строго в таком порядке: expectedRaw=<%q> actualRaw=<%q> expectedNorm=<%q> actualNorm=<%q> firstDiff=<описание первого отличия> При этом значения raw/norm должны печататься с использованием формата %q.
• firstDiff должен находить первое отличие на уровне байтов в нормализованных строках, сравнивая []byte(expectedNorm) и []byte(actualNorm) без использования сторонних пакетов, и выводить: индекс отличия, байт expected (или -1 если expected закончилась), байт actual (или -1 если actual закончилась).
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	// По условию читаем ровно две строки именно через ReadString('\n').
	expectedRaw, _ := r.ReadString('\n')
	actualRaw, _ := r.ReadString('\n')

	// Norm строго через TrimSpace от raw-строк.
	expectedNorm := strings.TrimSpace(expectedRaw)
	actualNorm := strings.TrimSpace(actualRaw)

	if expectedNorm == actualNorm {
		fmt.Println("match")
		return
	}

	fmt.Printf("expectedRaw=<%q>\n", expectedRaw)
	fmt.Printf("actualRaw=<%q>\n", actualRaw)
	fmt.Printf("expectedNorm=<%q>\n", expectedNorm)
	fmt.Printf("actualNorm=<%q>\n", actualNorm)

	idx, expB, actB := firstDiff(expectedNorm, actualNorm)
	fmt.Printf("firstDiff=%d %d %d\n", idx, expB, actB)
}

func firstDiff(expectedNorm, actualNorm string) (idx int, expByte int, actByte int) {
	// TODO: Реализуйте поиск первого отличия на уровне байтов в нормализованных строках.
	// TODO: Нужно сравнивать именно []byte(expectedNorm) и []byte(actualNorm).
	// TODO: Верните индекс отличия, байт expected (или -1 если expected закончилась), байт actual (или -1 если actual закончилась).
	// TODO: Учтите случай, когда одна строка является префиксом другой (отличие только в длине).
	a := []byte(actualNorm)
	e := []byte(expectedNorm)

	n := len(a)

	if len(e) < n {
	    n = len(e)
	}

    for i := 0; i < n; i++ {
        if a[i] != e[i] {
            return i, int(e[i]), int(a[i])
        }
    }

    // Если одна строка — префикс другой, отличие в первом "лишнем" байте
	i := n
	exp := -1
	act := -1
	if i < len(e) {
		exp = int(e[i])
	}
	if i < len(a) {
		act = int(a[i])
	}
	return i, exp, act
}