package piscine

func IterativeFactorial(nb int) int {
	if nb > -20 && nb < 20 {
		if nb == 1 || nb == 0 {
			return 1
		}
		if nb > 1 {
			return nb * IterativeFactorial(nb-1)

		}
	}
	return 0
}
