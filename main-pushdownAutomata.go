package main
import (
  "github.com/kecbigmt/go-white-and-black-doors/automata/oldKufu_033"
  "fmt"
)

func main(){
  bs := oldKufu_033.TestPatterns
  for _, b := range bs{
    fmt.Println("Try:", b)
    err := oldKufu_033.Validate(b)
    if err != nil{
      fmt.Println("-> Invalid:",err)
    }else{
      fmt.Println("-> Valid")
    }
  }
}
