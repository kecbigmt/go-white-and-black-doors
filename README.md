# go-white-and-black-doors
川添 愛『[白と黒のとびら: オートマトンと形式言語をめぐる冒険](http://amzn.asia/ibyQLRz)』に登場する遺跡をGoで実装してみます。

小説に登場する文字列○と●を、それぞれ0と1で表現しました。

# How to use
1. 使いたいオートマトン（遺跡）を`import`で読み込み（ここでは第1古代ルル語の有限オートマトン*1を読み込み）
2. 検査したい文字列を1文字ずつ入れた[]byteを作成（ここでは変数`b`に格納）
3. `finiteAutomata.Validate(b)`で検査を実行
4. `finiteAutomata.Validate(b)`の返り値がnilであれば受理、nilでなければ拒否を意味する*2

*1 小説に最初に登場する遺跡です（本書p.16参照）。

*2 nilでない場合は返り値には`common.ValidationError`が入っており、エラー内容を確認できます。

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
	err := finiteAutomata.Validate(b)
	if err != nil{
		fmt.Println("-> Invalid:", err)
	}else{
		fmt.Println("-> Valid")
	}
}
```
## Output
```
Try: [0 1 0 0 1]
-> Valid
```

# Test Patterns
言語ごとにテスト用のスライスをひとつのスライスにまとめて定義しています。

`finiteAutomata.TestPatterns`に[][]byteとして格納されています。

## Script
```go
package main
import (
  "./OldLulu_008/finiteAutomata"
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

```
## Output
```
Try: [0 0]
-> Valid
Try: [1 1]
-> Valid
Try: [0 0 1 1]
-> Valid
Try: [1 0 1 0]
-> Valid
Try: [0 1 1 1 0 1]
-> Valid
Try: [1 1 1 0 1 0]
-> Valid
Try: [0 0 1 1 1 0 1 0]
-> Valid
Try: []
-> Valid
Try: [0 0 1 1 1 0 1 0 1 0]
-> Invalid: failed to reach exit(101) state=*finiteAutomata.roomC, index=9, input=0
Try: [0 0 1 1 1 0]
-> Invalid: failed to reach exit(101) state=*finiteAutomata.roomC, index=5, input=0
```
