package codewars

import (
	"strings"
	//"strconv"
	"fmt"
	//"reflect"
)


func Points(games []string) int{

	var res int 
	for _, game := range games{
		str := strings.Split(game,":")
		x,y := str[0],str[1]
		if x > y{res+=3}else if x==y{
			res+=1
		}else{
			res+=0
		}
	}	
	fmt.Println(res)
	return res
}
  