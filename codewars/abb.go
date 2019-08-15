package codewars

	import (
		"fmt"
	"strings"
	)


	func AbbrevName(name string) string{
		n:= strings.Split(name," ")
		res := string(n[0][0])+"."+string(n[1][0])
		fmt.Println(res)
		strings.ToUpper(res)
		fmt.Println(res)

		return res

	}