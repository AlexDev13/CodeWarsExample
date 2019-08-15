// условие задачи гласит: найти все множители факториала без его конечного решения.

package codewars

import (
	"fmt"
	"math"
)

func FactorialExample(f int) {
	for i := 2; i <= f; i++ {
		isPrime := true

		for j := 2; j < i; j++ {

			if i%j == 0 ||  math.Sqrt(float64(i))!= 0{

				isPrime = false
			}
		}

		if isPrime == true {
			
			
				fmt.Println(i)
			
			//res:= f/i
			
			//fmt.Println(i, res)
		}
	}

}

// func Example(a int) {
// 	//count :=1
// 	for i := 2; i < a; i++ {
// 		for n := 1; n < i; n++ {
// 			res := a/Pow(i, n)
// 			if res == 0 {
// 				continue
// 				}else{
// 					fmt.Println(res)
// 				}	
// 			fmt.Println(i," - ",res)
// 		}

// 	//fmt.Println(count)
// 	}

// }

// func Pow(a, b int) int {
// 	x :=math.Pow(float64(a), float64(b))
// 	return int(x)
// }
func Ex(f int){
	for i:=2;i<=f;i++{
		fmt.Println(i," - ",f/i)
	}
}