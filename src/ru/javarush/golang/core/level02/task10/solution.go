package main

import "fmt"

func main() {
	// read
	var scoreA, scoreB, scoreC int
	fmt.Scan(&scoreA, &scoreB, &scoreC)

	// compute
	// TODO: Посчитайте сумму трёх целых чисел scoreA, scoreB, scoreC и сохраните её в переменную sumScores.
	sumScores := scoreA + scoreB + scoreC

	// print
	fmt.Println(sumScores)
}