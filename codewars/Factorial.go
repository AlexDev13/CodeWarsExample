
package codewars

import (
	"fmt"
	"math"
	"strconv"
)

func FactorialExample(f int) string {
	var result int
	var sum int
	var str string

	for i := 2; i <= f; i++ {
		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 {

				isPrime = false
			}
		}

		if isPrime == true {

			for k := 1; k <= f; k++ {

				result = f / Pow(i, k)

				sum += result

			}
			str1 := strconv.Itoa(i)
			str2 := strconv.Itoa(sum)
		
			str = str1 + "^" + str2 + " * "
			fmt.Print(str)

		}
		sum = 0
	}
	return str
}

func Pow(a, b int) int {
	res := math.Pow(float64(a), float64(b))
	return int(res)
}
