package one

import (
	"fmt"
	"math"
)

func term1(x float64) {
	fmt.Println("- Input x:", x)
	var result = math.Pow(1.0+x, 4.0)
	fmt.Println("- Resultat:", result)
	fmt.Println()
}

func term2(x float64) {
	fmt.Println("- Input x:", x)
	var result = (math.Pow(x, 3.0) - math.Pow(3.0*x, 2.0) + (3.0 * x) - 1.0)
	fmt.Println("- Resultat:", result)
	fmt.Println()
}

func term3(x float64) {
	fmt.Println("- Input x:", x)
	var result = math.Cbrt(x)
	fmt.Println("- Resultat:", result)
	fmt.Println()
}

func term4(x float64) {
	fmt.Println("- Input x:", x)
	var result = math.Sqrt(x) / math.Pow(x+1.0, 2.0)
	fmt.Println("- Resultat:", result)
	fmt.Println()
}

func term5(x float64) {
	fmt.Println("- Input x:", x)
	var result = math.Sqrt(math.Sin(x) / math.Log(x))
	fmt.Println("- Resultat:", result)
	fmt.Println()
}

func printTerme() {
	var values = [5]float64{-2, 1.0 / 2.0, 1, 2, 4}
	fmt.Println("-------------------------------------------")
	fmt.Println("1 Terme")
	fmt.Println("Werte welche für x eingesetzt werden:")
	fmt.Println(values)
	fmt.Println("1. Term")
	for i := range values {
		term1(values[i])
	}

	fmt.Println("2. Term")
	for i := range values {
		term2(values[i])
	}

	fmt.Println("3. Term")
	for i := range values {
		term3(values[i])
	}

	fmt.Println("4. Term")
	for i := range values {
		term4(values[i])
	}

	fmt.Println("5. Term")
	for i := range values {
		term5(values[i])
	}
	fmt.Println("-------------------------------------------")
}
