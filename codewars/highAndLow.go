package codewars

import (
"fmt"
"strings"
"sort"
"strconv"
)


func HighAndLow(s string) string	{
	
	strA:= strings.Fields(s)
	sort.Strings(strA)
	var t2 = []int{}
	for _, i := range strA {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        t2 = append(t2, j)
    }
	sort.Ints(t2)
	min := t2[0]
	max := t2[0]
	for _, value := range t2 {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
		sample := fmt.Sprint(max,min)
	// fmt.Println(reflect.TypeOf(sample))

		return sample
		
}