package codewars
import "fmt"
import "math"

func MultipleOfIndex (ints []int) {
	var result []int
	for i:= 1;i<len(ints);i++{
		if ints[i]%i != 0{
			continue
		}
		result = append(result,ints[i])
	}
	
	
	fmt.Println(result)
	fmt.Println(math.Trunc(15.55))
}