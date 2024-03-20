/*
Durch einen Rundungsfehler wird die Zahl auf 0 gerundet.
Was natürlich ein mathematischer Fehler ist, jedoch nicht vermeidbar mit dem IEEE 754 standard.
*/

package one

import (
	"fmt"
	"math"
)

func printNull() {
	fmt.Println("-------------------------------------------")
	fmt.Println("5 Null")

	var result = 1.0
	var count = 0

	// Bei der 1024. teilung währe das resultat false
	for i := 1; result > 0; i++ {
		var resultTry = 1.0 / math.Pow(2.0, float64(i))
		result = resultTry
		count++
		//fmt.Println(i, "try:", result, ">", 0, result > 0)
	}

	fmt.Printf("Bei der %d. teilung ist das resultat false\n", count)
	fmt.Println("-------------------------------------------")
}
