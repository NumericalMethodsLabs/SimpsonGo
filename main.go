package main

import "math"

func main() {
	a := 0.0
	b := 3.1415926535
	n := 8
	e := 0.0001

	h := (b - a) / (2 * float64(n))
	x := make([]float64, 0)
	f := make([]float64, 0)
	for i := 0; i <= 2*n; i++ {
		x = append(x, a+float64(i)*h)
		f = append(f, math.Sin(a+float64(i)*h))
	}

	s1 := 0.0
	for i := 1; i <= 2*n-1; i++ {
		if i%2 == 1 {
			s1 = s1 + 4*f[i]
		}
	}
	s2 := 0.0
	for i := 1; i <= 2*n-1; i++ {
		if i%2 == 0 {
			s2 = s2 + 2*f[i]
		}
	}

	int1 := (h / 3) * (f[0] + s1 + s2 + f[2*n])
	n = n * 2
	h = (b - a) / (2 * float64(n))
	for i := 0; i < 2*n; i++ {
		x[i] = a + float64(i)*h
		f[i] = 1 / (x[i] + 2)
	}
	s1 = 0.0
	for i := 1; i <= 2*n-1; i++ {
		if i%2 == 1 {
			s1 = s1 + 4*f[i]
		}
	}
	s2 = 0.0
	for i := 1; i <= 2*n-1; i++ {
		if i%2 == 0 {
			s2 = s2 + 2*f[i]
		}
	}
	int2 := (h / 3) * (f[0] + s1 + s2 + f[2*n])
	pogr := math.Abs(int1 - int2)

}
