package codewars
import (
  "fmt"
  "strings"
)
func DNAtoRNA(dna string) {
   fmt.Println(strings.ReplaceAll(dna,"T","U"))
  }