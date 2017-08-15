package main
import (
  "./OldLulu_047/finiteAutomata"
  "fmt"
)

func main(){
  bs := finiteAutomata.TestPatterns
  for _, b := range bs{
    fmt.Println("Try:", b)
    finiteAutomata.Validate(b)
  }
}
