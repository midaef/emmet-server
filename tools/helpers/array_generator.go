package helpers

import (
	"errors"
	"strconv"
	"strings"
)

func StringToUint64Array(str string) ([]uint64, error) {
	strLocal := strings.ReplaceAll(strings.ReplaceAll(str, "{", ""), "}", "")

	numbersStr := strings.Split(strLocal, ",")

	var numbers []uint64

	if len(numbersStr) == 0 {
		return []uint64{}, nil
	}

	for _, n := range numbersStr {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, errors.New("convert string to uint64 array failed")
		}

		numbers = append(numbers, uint64(num))
	}

	return numbers, nil
}
