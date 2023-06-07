package main

import (
	"fmt"
)

func main() {
	var f64 float64 = 0.5
	fmt.Println(f64)
	fl := 0.7
	fmt.Println(f64 + fl)
	fmt.Printf("%T, %T\n", f64, fl)

	var f32 float32 = 0.9
	fmt.Printf("%T\n", f32)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := 0.0 / zero
	fmt.Println(nan)

	a := 1.0 / 0.5
	fmt.Println(a)

	var c64 complex64 = -5 + 12i
	fmt.Println(c64)

}
