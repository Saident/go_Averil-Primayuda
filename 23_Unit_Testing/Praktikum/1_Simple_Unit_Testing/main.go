package main

func Addition(a, b int) int {
	res := a + b
	return res
}

func Substract(a, b int) int {
	res := a - b
	return res
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, nil
	}
	res := a / b
	return res, nil
}

func Multiple(a, b int) int {
	res := a * b
	return res
}
