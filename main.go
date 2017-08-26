package main
import (
  "github.com/kecbigmt/go-white-and-black-doors/automata/oldLulu_001"
  "fmt"
)

func main(){
  bs := oldLulu_001.TestPatterns
  for _, b := range bs{
    fmt.Println("Try:", b)
    err := oldLulu_001.Validate(b)
    if err != nil{
      fmt.Println("-> Invalid:",err)
    }else{
      fmt.Println("-> Valid")
    }
  }
}
