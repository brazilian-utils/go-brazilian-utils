package helpers

func numberToWeightArray(weight, length int) []int {
	var array []int

	for i := 0; i < length; i++ {
		array = append(array, weight-i)
	}

	return array
}

func createChecksum(input string, weigths []int) int {
	var total int

	for index, char := range input {
		total = total + int(char-'0')*weigths[index]
	}

	return total
}

// GenerateChecksum creates a checksum based on the input
func GenerateChecksum(input string, weights []int) int {
	numericInput := OnlyNumbers(input)

	if len(weights) == 1 {
		return createChecksum(
			numericInput,
			numberToWeightArray(weights[0], len(numericInput)),
		)
	}

	return createChecksum(numericInput, weights)
}
