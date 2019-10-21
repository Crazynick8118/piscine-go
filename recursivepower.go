package piscine

func RecursivePower(nb int, power int) int {
	if nb > 20 || nb < -20 || power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}