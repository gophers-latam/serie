package math

func Compute(x int) int {

	result := 1

	for i := 1; i <= x; i++ {

		y := i % 2

		if y == 1 {
			result = result + i
		} else {
			result = result * i
		}

	}

	return result

}
