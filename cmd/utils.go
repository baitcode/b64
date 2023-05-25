package cmd

func removeTrainlingNewline(input []byte) []byte {
	// Check if the slice is empty
	if len(input) <= 0 {
		return input
	}

	// Check if the last byte is a newline
	if input[len(input)-1] != '\n' {
		return input
	}

	return input[:len(input)-1]
}
