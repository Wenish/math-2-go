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
		fmt.Printf("Bitmuster: %032b\n", bits)
		fmt.Println()
	}
	fmt.Println("Bitfolgen zu float32")
	fmt.Println()

	fmt.Printf("Zahl: -0 von Bitmuster zu Float32: %f\n", math.Float32frombits(0b_1_0000_0000_000_0000_0000_0000_0000_0000))
	fmt.Printf("Zahl: 1024 von Bitmuster zu Float32: %f\n", math.Float32frombits(0b_0_1000_1001_000_0000_0000_0000_0000_0000))
	fmt.Printf("Zahl: 1000000 von Bitmuster zu Float32: %f\n", math.Float32frombits(0b_0_1001_0010_111_0100_0010_0100_0000_0000))
	fmt.Println("-------------------------------------------")
}
