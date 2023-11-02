package one

import (
	"fmt"
	"math"
)

func printIeee754() {
	fmt.Println("-------------------------------------------")
	fmt.Println("2 IEEE 754")
	fmt.Println("float32 zu Bitfolgen")
	fmt.Println()
	var values = [7]float32{0, 1.0 / 3.0, -(1.0 / 3.0), 1, -2, 65536, -65536}
	for i := range values {
		myFloat := float32(values[i])
		bits := math.Float32bits(myFloat)
		fmt.Printf("Float32: %f\n", myFloat)
		fmt.Printf("Bitmuster: %b\n", bits)
		fmt.Println()
	}
	fmt.Println("Bitfolgen zu float32")
	fmt.Println()

	// 0 -> binärzahl 0
	// 1024 -> binärzahl 10000000000
	// 1000000  -> binärzahl 11110100001001000000

	// TODOOOOOO: binärzahl zu float32 printen
	fmt.Println("-------------------------------------------")
}
