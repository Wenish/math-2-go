package one

import (
	"fmt"
	"math"
)

func printMathPi() {
	fmt.Println("-------------------------------------------")
	fmt.Println("6 math.Pi ̸= π")
	fmt.Println("math.Pi", math.Pi)

	// 1. Beginne mit einem Intervall [A, B], für welches gilt: π ∈ [A, B]
	// 2. Falls A = B → Ende!
	// 3. Aus den Grenzen des Intervalls A und B berechne ein C, so dass A < C < B
	// 4. Falls π ∈ [A, C] → neues Intervall ist [A, C]
	// 5. Falls π ∈ [C, B] → neues Intervall ist [C, B]
	// 6. Gehe zu Punkt 2

	A := 3.0
	B := 3.2

	/* THIS IS SLOW AF AND DOSNT WORK
	for A != B {
		C := (A + B) / 2

		if math.Pi <= C {
			B = C
		} else {
			A = C
		}

		fmt.Println("[", A, C, "]")
	}
	*/
	fmt.Printf("π liegt im Intervall [%.64f, %.64f]\n", A, B)
	fmt.Println("-------------------------------------------")
}
