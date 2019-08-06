package main

 import "fmt"

func Hero(bullets, dragons int) bool {
	if (bullets-(dragons*2)) < 0 {
		return false
	  }
	return true
}
   
func main(){
fmt.Println(Hero(100,40))
}