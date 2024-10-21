package utils

func ValidateLuhn(number string) bool {
	var sum int
	parity := len(number) % 2

	for i, digit := range number {
		n := int(digit - '0')
		if i%2 == parity {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}

	return sum%10 == 0
}
