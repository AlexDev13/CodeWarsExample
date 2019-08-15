package codewars
import "math"
import "fmt"
func SquareOrSquareRoot(arr []int) []int{
   var result []int
    for _, el := range arr{
      y := math.Sqrt(float64(el))
      if int(y) %1==0{
        result = append(result,int(y))
      }else{
      result = append(result,int(y*y))
      }
	}
	fmt.Println(result)
    return result
}