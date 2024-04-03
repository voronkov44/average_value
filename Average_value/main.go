package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Введите свои оценки (для завершения введите 'q'):")

	sum := 0.0
	count := 0

	for {
		var input string
		fmt.Scanln(&input)

		if input == "q" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка ввода числа:", err)
			os.Exit(1)
		}

		sum += num
		count++
	}

	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Среднее число: %.2f\n", average)
		fmt.Printf("Какую оценку поставят: %.2f\n", math.Round(average))
	} else {
		fmt.Println("Нет чисел для подсчета среднего")
	}
}
