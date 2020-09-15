package main

import (
	"errors"
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	calc(value1, value2, operation)
}

func readTask() (interface{}, interface{}, interface{}) {
	return 34.4, true, "/"
}

func calc(v1, v2, op interface{}) {
	var (
		err       error
		result    float64
		v1float64 float64
		v2float64 float64
		opStr     string
	)

	v1float64, floatErr := toFloat64(v1)
	if floatErr != nil {
		err = floatErr
	}

	v2float64, floatErr = toFloat64(v2)
	if floatErr != nil {
		err = floatErr
	}

	opStr, signErr := toStringMathSign(op)
	if signErr != nil {
		err = signErr
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	switch opStr {
	case "+":
		result = v1float64 + v2float64
	case "-":
		result = v1float64 - v2float64
	case "*":
		result = v1float64 * v2float64
	case "/":
		result = v1float64 / v2float64
	default:
		panic("How can it be?")
	}

	fmt.Printf("%.4f", result)
	return
}

func toFloat64(v interface{}) (float64, error) {
	var (
		result float64
		err    error
	)

	result, ok := v.(float64)
	if ok == false {
		err = errors.New(fmt.Sprintf("value=%v: %T", v, v))
	}

	return result, err
}

func toStringMathSign(sign interface{}) (string, error) {
	var (
		result string
		errMsg = "неизвестная операция"
	)

	result, ok := sign.(string)
	if ok == false {
		return result, errors.New(errMsg)
	}

	var mathSigns = [4]string{"+", "-", "*", "/"}
	for _, sign := range mathSigns {
		if sign == result {
			return result, nil
		}
	}

	return "", errors.New(errMsg)
}
