package utils

import (
	"errors"
	cu "github.com/tatiananeda/calculator/utils"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

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

			if cu.Operations[item] == nil {
				return nil, err
			}

			res, err := cu.Operations[item](left, right)

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
