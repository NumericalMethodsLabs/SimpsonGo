package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x * x * x
}

func main() {
	a := 0.0
	b := 1.0
	n := 2
	eps := 10e-5

	h := (b - a) / float64(n)
	s0, s1, s2 := f(a)+f(b), 0.0, 0.0

	for i := 1; i <= n/2; i++ {
		s1 += f(a + (2*float64(i)-1)*h)
	}

	for i := 1; i < n/2; i++ {
		s2 += f(a + 2*float64(i)*h)
	}

	var i1, i2 float64 = h / 3 * (s0 + 4*s1 + 2*s2), h / 3 * (s0 + 4*s1 + 2*s2)

	i1 = i2
	n *= 2
	h = (b - a) / float64(n)
	s2 = s1 + s2
	s1 = 0
	for i := 1; i <= n/2; i++ {
		s1 += f(a + (2*float64(i)-1)*h)
	}
	i2 = h / 3 * (s0 + 4*s1 + 2*s2)
	for math.Abs(i2-i1) > 15.0/16.0*eps {
		i1 = i2
		n *= 2
		h = (b - a) / float64(n)
		s2 = s1 + s2
		s1 = 0
		for i := 1; i <= n/2; i++ {
			s1 += f(a + (2*float64(i)-1)*h)
		}
		i2 = h / 3 * (s0 + 4*s1 + 2*s2)
	}
	fmt.Println("Количество:", n)
	fmt.Println("I:", i2)
}
