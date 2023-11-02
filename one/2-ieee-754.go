package one

import (
	"fmt"
	"math"
)

func PrintIeee754() {
	fmt.Println("-------------------------------------------")
	fmt.Println("2 IEEE 754")
	fmt.Println(math.Float64bits(1))
	fmt.Println("-------------------------------------------")
}
