/*
Wenn man in go bei der Zahl 2 die Wurzel zieh entsteht ein ungenaus Ergebnis. Da es ein Ergebnis ist, welches nicht in R ausgeben werden kann.
Wenn man dieses ungenau Ergebnis quadriert wird es nicht wieder Wurzel 2 ergeben.
*/

package one

import (
	"fmt"
	"math"
)

func printWurzeln() {
	fmt.Println("-------------------------------------------")
	fmt.Println("3 Wurzeln")

	equation := math.Pow(math.Sqrt(2), 2) == 2

	fmt.Println("Die Gleichung ist in go folgendes:", equation)
	fmt.Println("-------------------------------------------")
}
