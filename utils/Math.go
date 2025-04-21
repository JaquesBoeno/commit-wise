package utils

func ArithmeticMod(value, mod int) int {
	result := value % mod
	if result < 0 {
		return mod + result
	}

	return result
}
