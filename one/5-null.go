package one

import (
	"fmt"
	"math"
)

func printNull() {
	fmt.Println("-------------------------------------------")
	fmt.Println("5 Null")

	var result = 1.0

	// Bei der 1024. teilung währe das resultat false
	for i := 1; result > 0; i++ {
		var resultTry = 1.0 / math.Pow(2.0, float64(i))
		result = resultTry
		fmt.Println(i, "try:", result, ">", 0, result > 0)
	}

	fmt.Println("Bei der 1024. teilung ist das resultat false")
	fmt.Println("-------------------------------------------")
}
