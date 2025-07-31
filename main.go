package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	fmt.Print("Выберите операцию (AVG/SUM/MED): ")
	fmt.Scanln(&operation)

	fmt.Print("Введите числа через запятую: ")
	var input string
	fmt.Scanln(&input)

	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	switch strings.ToUpper(operation) {
	case "AVG":
		fmt.Printf("Среднее: %.2f\n", calculateAvg(numbers))
	case "SUM":
		fmt.Printf("Сумма: %.2f\n", calculateSum(numbers))
	case "MED":
		fmt.Printf("Медиана: %.2f\n", calculateMedian(numbers))
	default:
		fmt.Println("Неизвестная операция. Допустимые: AVG, SUM, MED")
	}

}

func parseNumbers(input string) ([]float64, error) {
	parts := strings.Split(input, ",")
	var numbers []float64

	for _, part := range parts {
		numStr := strings.TrimSpace(part)
		if numStr == "" {
			continue
		}

		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, fmt.Errorf("'%s' не является числом", numStr)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("не введено ни одного числа")
	}

	return numbers, nil
}

func calculateAvg(numbers []float64) float64 {
	sum := calculateSum(numbers)
	return sum / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 1 {
		return numbers[n/2]
	}
	return (numbers[n/2-1] + numbers[n/2]) / 2
}
