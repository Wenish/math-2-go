/*
Durch Rundungsfehler innerhalb der Berechung führt es dazu dass die 2 Terme nicht mehr gleich sind.

Beide Typen float64 und int64 belegen 64 Bits im Speicher.

Der Unterschied ist, dass int64 nur Ganzzahlen ohne Dezimalstellen darstellt.

Float64 kann Gleitkommazahlen mit Hilfe des IEEE 754-Standard darstellen.
Float64 hat eine begrenzte Genauigkeit. Deswegen können Rundungsfehler bei der Berechnung auftreten.
Wenn man eine höhere Genauigkeit brauch, muss man eine Library einsetzen.
*/

package one

import (
	"fmt"
	"math"
)

func equation(a float64, b float64, c float64) bool {
	return (a+b)+c == a+(b+c)
}

func printAssoziativ() {
	fmt.Println("-------------------------------------------")
	fmt.Println("4 Assoziativ")

	var a = math.Pow(10, 25)
	var b = math.Pow(-10, 25)
	var c = 1.0

	fmt.Println("Die Gleichung ist in go folgendes:", equation(a, b, c))
	fmt.Println("-------------------------------------------")
}
