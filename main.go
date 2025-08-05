package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operations := map[string]func([]float64) float64{
		"AVG": calculateAvg,
		"SUM": calculateSum,
		"MED": calculateMedian,
	}

	var operation string
	fmt.Print("Выберите операцию (AVG/SUM/MED): ")
	fmt.Scanln(&operation)

	calcFunc, exists := operations[strings.ToUpper(operation)]
	if !exists {
		fmt.Println("Неизвестная операция. Допустимые:", getAvailableOperations(operations))
		return
	}

	fmt.Print("Введите числа через запятую: ")
	var input string
	fmt.Scanln(&input)

	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	result := calcFunc(numbers)
	fmt.Printf("%s: %.2f\n", strings.ToUpper(operation), result)
}

func getAvailableOperations(ops map[string]func([]float64) float64) string {
	keys := make([]string, 0, len(ops))
	for k := range ops {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
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
