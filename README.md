# go-white-and-black-doors
川添 愛『[白と黒のとびら: オートマトンと形式言語をめぐる冒険](http://amzn.asia/ibyQLRz)』に登場する遺跡をGoで実装してみる

# How to use
小説に登場する文字列○と●を、それぞれ0と1で表現しています。

検査したい文字列を1文字ずつ入れた[]byteを`b`とすると、`finiteAutomata.Validate(b)`で有限オートマトンによる検査を実行できます。

## Script
```go
package main
import (
	"./OldLulu_001/finiteAutomata"
	"fmt"
)

func main(){
	b := []byte{0,1,0,0,1}
	fmt.Println("Try:", b)
	finiteAutomata.Validate(b)
}
```
## Output
```
Try: [0 1 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)
```

# Test Patterns
言語ごとにテスト用のスライスをひとつのスライスにまとめて定義しています。

`finiteAutomata.TestPatterns`に[][]byteとして格納されています。

## Script
```go
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
```
## Output
```
Try: [1 1 1 1 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 1 1 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 1 0 1 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 1 0 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 0 1 1 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 0 1 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 0 0 1 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [1 0 0 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [0 1 1 1 0 1]
->OK! state=&{}(*finiteAutomata.exit)
Try: [0 1 1 0 0 1]
->OK! state=&{}(*finiteAutomata.exit)

```
