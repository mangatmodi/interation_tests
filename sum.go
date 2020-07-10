package main

import (
	"github.com/pkg/errors"
	"math"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Divide by 0")
	}
	return a / b, nil
}

func log(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return -1, errors.New("Can't log 0")
	}
	return math.Log(a) / math.Log(b), nil
}
