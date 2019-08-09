package codewars

import (
   "fmt"
)
func Century(year int) int {
	
      cen := year/100+1
      if year >=10000{
         cen+=10
      }
      fmt.Println(cen)
      return cen
   }