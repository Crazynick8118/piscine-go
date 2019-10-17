package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i < '9'; i++ {
		for j := '0'; i < 58; j++ {
			for k := '0'; i < 58; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)

				}

			}
		}
	}
	z01.PrintRune('\n')
}
