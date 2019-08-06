package codewars


func Hero(bullets, dragons int) bool {
	if (bullets-(dragons*2)) < 0 {
		return false
	  }
	return true
}
   