package codewars

import (
"fmt"
)


func MonkeyCount(n int) []int {
			var arr []int
		for i:=1;i<=n;i++{
			arr = append(arr,i)
		}
		
		return arr
 }