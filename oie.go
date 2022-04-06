package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 5.5

	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferencia é", area)

	const (
		a = 5
		b = 90
	)
	var (
		c = 7
		d = 1
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
