package main
import (
  "github.com/kecbigmt/go-white-and-black-doors/automata/oldLulu_047"
  "fmt"
)

func main(){
  bs := oldLulu_047.TestPatterns
  for _, b := range bs{
    fmt.Println("Try:", b)
    err := oldLulu_047.Validate(b)
    if err != nil{
      fmt.Println("-> Invalid:",err)
    }else{
      fmt.Println("-> Valid")
    }
  }
}
