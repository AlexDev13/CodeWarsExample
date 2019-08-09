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
