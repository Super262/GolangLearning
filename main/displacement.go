package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		s := 0.5*a*t*t + v0*t + s0
		return s
	}
}

func main() {
	var a, v0, s0 float64
	fmt.Println("Set the acceleration. ")
	_, _ = fmt.Scan(&a)
	fmt.Println("Set the initial velocity. ")
	_, _ = fmt.Scan(&v0)
	fmt.Println("Set the initial displacement. ")
	_, _ = fmt.Scan(&s0)
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
