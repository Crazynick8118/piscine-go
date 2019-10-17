package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for k := i + 1; k <= '9'; k++ {
			for j := k + 1; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(k)
				z01.PrintRune(j)
			}

		}
	}
	z01.PrintRune('\n')
}
