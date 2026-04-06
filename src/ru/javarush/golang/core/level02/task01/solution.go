package main

import "fmt"

func main() {
	var startBalance int = 1000
	var income int = 500
	var purchase int = 175

	finalBalance := 0
	// TODO: Вычислите итоговый баланс по формуле из условия, используя + и -, и сохраните в finalBalance.
	finalBalance = startBalance + income - purchase

	fmt.Println(finalBalance)
}