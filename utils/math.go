package utils

import (
	"errors"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var operations = map[string]func(float64, float64) (float64, error){
	"+": func(x, y float64) (float64, error) {
		return x + y, nil
	},
	"-": func(x, y float64) (float64, error) {
		return x - y, nil
	},
	"*": func(x, y float64) (float64, error) {
		return x * y, nil
	},
	"/": func(x, y float64) (float64, error) {
		if y == 0 {
			return 0, errors.New("Can't divide by 0")
		}
		return x / y, nil
	},
}

var operationsByPrecedence = [2][2]string{
	{"*", "/"},
	{"+", "-"},
}

func processOperations(input []string, ops []string) ([]string, error) {
	for i := 0; i < len(input); i++ {
		item := input[i]

		if slices.Contains(ops, item) {

			left, err := parseFloat(input[i-1])
			if err != nil {
				return nil, err
			}

			right, err := parseFloat(input[i+1])
			if err != nil {
				return nil, err
			}

			if operations[item] == nil {
				return nil, err
			}

			res, err := operations[item](left, right)

			if err != nil {
				return nil, err
			}

			result := Splice(input, i-1, i+1, strconv.FormatFloat(res, 'f', -1, 64))

			return processOperations(result, ops)
		}
	}

	return input, nil
}

func ProcessInput(s string) (float64, error) {
	exp := strings.Split(s, " ")
	first, err := processOperations(exp, operationsByPrecedence[0][:])

	if err != nil {
		return 0, err
	}

	second, err := processOperations(first, operationsByPrecedence[1][:])

	if err != nil {
		return 0, err
	}

	if len(second) > 1 {
		return 0, errors.New("Unsupported operations were passed on " + s)
	}

	return parseFloat(second[0])
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}
