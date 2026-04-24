package main

/*
= Задача №12. Разбор команды =
Вы пишете консольный калькулятор «для инженеров»: пользователь вводит три токена (a op b), а вы должны сначала аккуратно разобрать ввод (включая ошибки), а потом вычислить результат (включая деление на ноль). Чтобы код было удобно читать, вы делаете отдельную функцию разбора с именами результатов прямо в сигнатуре.
Программа читает три токена: aToken, opToken, bToken. Сделайте функцию parseExpr(aS, opS, bS string) (a int, op string, b int, err error), где имена результатов показывают, что именно возвращается. Если оператор неизвестен или число не парсится — верните ненулевую ошибку. В parseExpr на ошибках используйте явный return со значениями (без naked return), чтобы поведение было однозначным. Отдельно сделайте calc(a int, op string, b int) (int, error): она считает результат и возвращает ошибку при делении на ноль или неизвестном операторе, и тоже использует только явные return со значениями.
Формат вывода:
- при успехе: result: <число>
- при ошибке: error: <текст ошибки>
Требования:
• В коде должна быть реализована функция `parseExpr(aS, opS, bS string) (a int, op string, b int, err error)`, где имена возвращаемых значений соответствуют смыслу результата (a, op, b, err).
• Функция `parseExpr` должна преобразовывать `aS` и `bS` в `int` с помощью `strconv.Atoi`; при ошибке парсинга должна возвращать ненулевую ошибку через `err`.
• Функция `parseExpr` должна считать допустимым оператором только один из: `+`, `-`, `*`, `/`; при любом другом значении `opS` должна возвращать ошибку.
• Во всех ошибочных ветках внутри `parseExpr` должен использоваться только явный `return` с конкретными возвращаемыми значениями (запрещён `return` без значений).
• В коде должна быть реализована функция `calc(a int, op string, b int) (int, error)` (без именованных результатов), которая выполняет вычисление по оператору и возвращает ошибку при неизвестном операторе.
• Функция `calc` должна возвращать ошибку, если `op == "/"` и `b == 0`.
*/

import (
	"errors"
	"fmt"
	"strconv"
)

// parseExpr разбирает три токена "a op b" в типизированное выражение.
// Имена результатов в сигнатуре специально показывают смысл возвращаемых значений.
func parseExpr(aS, opS, bS string) (a int, op string, b int, err error) {
	aVal, err := strconv.Atoi(aS)
	if err != nil {
		return 0, "", 0, fmt.Errorf("cannot parse integer: %s", aS)
	}

	bVal, err := strconv.Atoi(bS)
	if err != nil {
		return 0, "", 0, fmt.Errorf("cannot parse integer: %s", bS)
	}

	// TODO: Проверьте, что opS — один из: +, -, *, /.
	// TODO: Если оператор неизвестен — верните ошибку (явным return со значениями).
	if opS != "+" && opS != "-" && opS != "*" && opS != "/" {
	    return 0, "", 0, fmt.Errorf("unknown operator: %s", opS)
	}

	return aVal, opS, bVal, nil
}

// calc выполняет вычисление; неизвестный оператор и деление на ноль — ошибки.
func calc(a int, op string, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
        return a - b, nil
	case "*":
        return a * b, nil
	case "/":
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func main() {
	var aToken, opToken, bToken string
	fmt.Scan(&aToken, &opToken, &bToken)

	a, op, b, err := parseExpr(aToken, opToken, bToken)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	res, err := calc(a, op, b)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %d\n", res)
}