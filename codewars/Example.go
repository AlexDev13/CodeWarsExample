package codewars

import (
	"fmt"
	//"math"
)

func Example(a int) {
	var result int
	var n = 2
	var sum int
	for k := 1; k <= a; k++ {
		// if result == 0{
		// 	continue
		// }
		
		result = a / Pow(n, k)
		//	result += result
		sum += result

	}
	fmt.Println(sum)
}

// func Pow(a, b int) int {
// 	res := math.Pow(float64(a), float64(b))
// 	return int(res)
// }
