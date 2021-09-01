package helpers

import (
	"errors"
	"strconv"
	"strings"
)

const (
	ErrorConvertStringToUint64Array = "convert string to uint64 array failed"
)

func StringToUint64Array(str string) ([]uint64, error) {
	if str == "" {
		return nil, errors.New(ErrorConvertStringToUint64Array)
	}

	strLocal := strings.ReplaceAll(strings.ReplaceAll(str, "{", ""), "}", "")

	numbersStr := strings.Split(strLocal, ",")

	var numbers []uint64

	if len(numbersStr) == 0 {
		return nil, errors.New(ErrorConvertStringToUint64Array)
	}

	for _, n := range numbersStr {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, errors.New(ErrorConvertStringToUint64Array)
		}

		numbers = append(numbers, uint64(num))
	}

	return numbers, nil
}

func Uint64ArrayToString(numbers []uint64) (string, error) {
	if len(numbers) == 0 {
		return "", errors.New("convert uint64 array to string failed")
	}

	var str string

	str += "{"

	for i := 0; i < len(numbers); i++ {
		if len(numbers)-1 == i {
			str += strconv.Itoa(int(numbers[i]))
		} else {
			str += strconv.Itoa(int(numbers[i])) + ","
		}
	}

	str += "}"

	return str, nil
}
