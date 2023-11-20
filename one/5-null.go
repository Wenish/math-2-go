package one

import (
	"fmt"
	"math"
)

func printNull() {
	fmt.Println("-------------------------------------------")
	fmt.Println("5 Null")

	var amount = 10 // Bei der 1024. teilung währe das resultat false

	// TODO: algorythmus der dann bei 0 automatisch stoppt
	for i := 1; i <= amount; i++ {
		var resultTry = 1.0 / math.Pow(2.0, float64(i))
		fmt.Println(i, "try:", resultTry, ">", 0, resultTry > 0)
	}

	var resultTry1023 = 1.0 / math.Pow(2.0, float64(1023))
	fmt.Println(1023, "try:", resultTry1023, ">", 0, resultTry1023 > 0)

	var resultTry1024 = 1.0 / math.Pow(2.0, float64(1024))
	fmt.Println(1024, "try:", resultTry1024, ">", 0, resultTry1024 > 0)
	fmt.Println("Bei der 1024. teilung ist das resultat false")
	fmt.Println("-------------------------------------------")
}
