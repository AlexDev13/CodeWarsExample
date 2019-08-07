package codewars

import (
	"fmt"
	"strings"
)

func PartList(arr []string) string{
	onlyStr :=strings.Join(arr," ") 
	result := strings.Fields(onlyStr)
	var all string
	var a string

	for i:=1; i<len(result); i++ {
		all +=	fmt.Sprintf("(%v%v%v)",strings.Join(result[:i], " "),", ",strings.Join(result[i:], " "))
	}
	fmt.Println(a)
	return all	
}