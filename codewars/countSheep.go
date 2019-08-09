package codewars

import (
	
	"strconv"
)

func CountSheep(num int) string {
	var sheep string
	for i:=1;i<=num;i++{
		sheep += strconv.Itoa(i)+"sheep"
	}
	return sheep
}	


// package kata

// import (
//   "fmt"
//   "strings"
// )

// func countSheep(num int) string {
//   var sb strings.Builder

//   for count := 1; count <= num; count++ {
//         fmt.Fprintf(&sb, "%d sheep...", count)
//   }
  
//   return sb.String()
// }