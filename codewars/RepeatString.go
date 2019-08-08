package codewars

import (
	"fmt"
)

func RepeatStr(repetitions int, value string) string {
	var str string
	for i:=0; i<repetitions;i++{
			str += value
	}
	fmt.Println(str)
	return str
  }