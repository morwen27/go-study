package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type OperationSet struct {
	handler func([]int) float64
	message string
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	operation := getOperation()
	operands := getOperands()
	getOperationResult(operation, operands)
}

func getOperationResult(operation string, operands []int) {
	operationsMap := map[string]OperationSet{
		"avg": {handler: getAvarage, message: "Среднее арифметическое %v для равно: %.2f\n"},
		"sum": {handler: getSum, message: "Сумма всех чисел из %v равна: %.2f\n"},
		"med": {handler: getMedium, message: "Медиана для %v равна: %.2f\n"},
	}

	result := operationsMap[operation].handler(operands)
	fmt.Printf(operationsMap[operation].message, operands, result)
}

func getMedium(operands []int) float64 {
	operandsLength := len(operands)
	slices.Sort(operands)

	if operandsLength%2 == 0 {
		leftIndex := operandsLength/2 - 1
		rightIndex := operandsLength / 2
		sum := float64(operands[leftIndex] + operands[rightIndex])

		return sum / 2
	} else {
		return float64(operands[operandsLength/2])
	}
}

func getSum(operands []int) float64 {
	sum := 0.0

	for i := range operands {
		sum += float64(operands[i])
	}

	return sum
}

func getAvarage(operands []int) float64 {
	return getSum(operands) / float64(len(operands))
}

func getOperands() []int {
	fmt.Print("Введите числа (целые или дробные) через запятую: ")
	userInput := bufio.NewScanner(os.Stdin)

	if userInput.Scan() {
		text := userInput.Text()
		items := strings.Split(text, ",")

		if len(items) <= 1 {
			showError(errors.New("Вы ввели недостаточное количество операндов."))
		}

		operands := make([]int, len(items), cap(items))

		for i := range items {
			operand, incorrectParsing := strconv.Atoi(strings.TrimSpace(items[i]))
			showError(incorrectParsing)
			operands[i] = operand
		}

		return operands
	}

	showError(errors.New("Ввод не был считан"))

	return make([]int, 1)
}

func getOperation() string {
	var operation string = ""

	fmt.Println("Доступные операции: AVG - нахождение среднего, SUM - нахождение суммы, MED - получние медианы.")
	fmt.Printf("Введите код операции: ")
	if _, errorMessage := fmt.Scan(&operation); errorMessage != nil {
		showError(errors.New("Input is incorrect!"))
	}

	checkOperation(operation)

	return operation
}

func checkOperation(operation string) {
	operations := [3]string{"AVG", "SUM", "MED"}
	isMatch := false

	for i := range operations {
		if strings.EqualFold(operation, operations[i]) {
			isMatch = true

			break
		}
	}

	if !isMatch {
		showError(errors.New("Вы ввели неподдерживаемую операцию"))
	}
}

func showError(errorMessage error) {
	if errorMessage != nil {
		panic(errorMessage)
	}
}
