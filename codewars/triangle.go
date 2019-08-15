package codewars

import (
	"math"
)

func EquableTriangle(a, b, c int) bool {
	P := a+b+c
	p := (a+b+c)/2
	h := 2*math.Sqrt(float64(p*(p-a)*(p-b)*(p-c)))
	ha:= int(h)/a
	S := (a*int(ha))/2
	if P == S{
		return true
		}

	return false
}