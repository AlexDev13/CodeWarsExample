package codewars

import (
	"fmt"
)


func SumEvenFibonacci(limit int) int {
	if limit == 0{
		return 0
	}
	if limit ==1{
		return 1
	}
	if limit == 2{
		return 2
	}
	var a, b int = 1, 1
	sum :=0
		for i :=0; i <= limit; i++{
			a, b = b, a+b
	
			if a%2!=0{
				continue
			}
			
			
			if a >limit {
				break
			}
			sum +=a
		//	fmt.Println(a)
			// fmt.Println(a)
			fmt.Println(sum)
		}

return sum

}