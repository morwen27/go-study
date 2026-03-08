package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	initialCurrency, targetCurrency, amount, errorMessage := getUserInput()
	if errorMessage != nil {
		fmt.Println(errorMessage)

		return
	}

	result := calculateCurrency(initialCurrency, targetCurrency, amount)
	fmt.Printf("%.2f %s = %.2f %s", amount, initialCurrency, result, targetCurrency)
}

func getRatiosMap() *map[string]map[string]float64 {
	ratios := map[string]map[string]float64{
		"USD": {
			"EUR": 0.84,
			"RUB": 75.25,
		},
	}
	ratiosPtr := &ratios

	currencyRelatedWithUSD := []string{}
	for key := range (*ratiosPtr)["USD"] {
		currencyRelatedWithUSD = append(currencyRelatedWithUSD, key)
	}

	for i, key := range currencyRelatedWithUSD {
		(*ratiosPtr)[key] = map[string]float64{
			"USD": 1 / (*ratiosPtr)["USD"][key],
		}

		for j, restKey := range currencyRelatedWithUSD {
			if i != j {
				(*ratiosPtr)[key][restKey] = 1 / ((*ratiosPtr)["USD"][key] / (*ratiosPtr)["USD"][restKey])
			}
		}
	}

	return ratiosPtr
}

func calculateCurrency(source string, target string, amount float64) float64 {
	ratios := getRatiosMap()

	return amount * (*ratios)[source][target]
}

func getUserInput() (string, string, float64, error) {
	initialCurrency, errorMessage := defineByUser("исходную валюту", "USD, EUR, RUB")
	if errorMessage != nil {
		return "", "", 0, errorMessage
	}

	amount, errorMessage := defineByUser("сумму", "")
	if errorMessage != nil {
		return "", "", 0, errorMessage
	}

	targetCurrency, errorMessage := defineByUser("целевую валюту", getTargetCurrencyTip(initialCurrency))
	if errorMessage != nil {
		return "", "", 0, errorMessage
	}

	if targetCurrency == initialCurrency {
		return "", "", 0, errors.New("CURRENCIES ARE THE SAME")
	}

	convertedAmount, convertError := strconv.ParseFloat(amount, 64)

	return initialCurrency, targetCurrency, float64(convertedAmount), convertError
}

func getTargetCurrencyTip(initialCurrency string) string {
	switch initialCurrency {
	case "RUB":
		return "USD, EUR"
	case "USD":
		return "EUR, RUB"
	case "EUR":
		return "USD, RUB"
	default:
		return ""
	}
}

func defineByUser(param string, tip string) (string, error) {
	var input string

	fmt.Printf("Введите %s (%s): ", param, tip)
	_, errorMessage := fmt.Scan(&input)

	hasError := errorMessage != nil || defineInputError(tip == "", input)

	if hasError {
		fmt.Println("Вы ошиблись. Попробуйте ввести еще раз. ")
		fmt.Printf("Введите %s (%s): ", param, tip)

		if _, errorMessage := fmt.Scan(&input); errorMessage != nil || defineInputError(tip == "", input) {
			return "", errors.New("INPUT IS INCORRECT")
		}
	}

	return input, nil
}

func checkStringInput(input string) bool {
	return input != "USD" && input != "RUB" && input != "EUR"
}

func checkAmoutInput(input string) bool {
	_, parsingError := strconv.ParseFloat(input, 64)

	return parsingError != nil
}

func defineInputError(contidion bool, input string) bool {
	if contidion {
		return checkAmoutInput(input)
	} else {
		return checkStringInput(input)
	}
}
