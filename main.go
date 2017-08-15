package main
import (
  "./oldLulu_008/finiteAutomata"
  "fmt"
)

func main(){
  bs := finiteAutomata.TestPatterns
  for _, b := range bs{
    fmt.Println("Try:", b)
    err := finiteAutomata.Validate(b)
    if err != nil{
      fmt.Println("-> Invalid:",err)
    }else{
      fmt.Println("-> Valid")
    }
  }
}
