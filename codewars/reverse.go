package codewars

import (
	"fmt"
)

func Solution(word string)  string {
	var result string
	for _,v := range word{
		result = string(v)+result
	}
	fmt.Println(result)
	return result
}