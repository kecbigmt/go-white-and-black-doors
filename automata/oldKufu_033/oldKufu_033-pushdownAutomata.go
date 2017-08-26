package oldKufu_033
import (
  "github.com/kecbigmt/go-white-and-black-doors/automata/common"
  "fmt"
)

type room interface{
  openDoor(s []byte, i int, v uint8) (interface{}, []byte, error)
}

type entrance struct{}
func (r *entrance) openDoor(s []byte, i int, v uint8) (interface{}, []byte, error){
  if s[0] == '$'{
    return &roomA{}, s, nil
  }else{
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  }
}

type roomA struct{}
func (r *roomA) openDoor(s []byte, i int, v uint8) (interface{}, []byte, error){
  fmt.Println("number", len(s), i+1)
  switch {
  case len(s) == i+1 && v == uint8(0):
    fmt.Println("go to RoomB")
    return &roomB{}, s[1:], nil
  case len(s) == i+1 && v == uint8(1):
    fmt.Println("go to RoomB")
    return &roomB{}, s[1:], nil
  case len(s) != i+1 && v == uint8(0):
    return &roomA{}, append([]byte{0}, s...), nil
  case len(s) != i+1 && v == uint8(1):
    return &roomA{}, append([]byte{1}, s...), nil
  default:
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  }
}

type roomB struct{}
func (r *roomB) openDoor(s []byte, i int, v uint8) (interface{}, []byte, error){
  if s[0] == '$' {
    return &exit{}, s, nil
  }
  switch v{
  case uint8(0):
    return &roomB{}, s[1:], nil
  case uint8(1):
    return &roomB{}, s[1:], nil
  default:
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  }
}

type exit struct{}
func (r *exit) openDoor(s []byte, i int, v uint8) (interface{}, []byte, error){
  switch v{
  case uint8(0):
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  case uint8(1):
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  default:
    return 0, s, common.ValidationError{"invalid input", 100, r, i, v, s}
  }
}

func Validate(b []byte) error{
  var (
    stack []byte = []byte{'$'}
    state room
    result interface{}
    err error
  )
  state = &entrance{}
  for i,v := range b{
    fmt.Println(i, v, stack)
    result, stack, err = state.openDoor(stack, i, uint8(v))
    if err != nil{
      return err
    }
    switch tmp := result.(type){
    case *entrance:
      state = tmp
    case *roomA:
      state = tmp
    case *roomB:
      state = tmp
    case *exit:
      state = tmp
    }
  }
  if _, ok:=state.(*exit);ok{
    return nil
  }else{
    return common.ValidationError{"failed to reach exit", 101, state, len(b)-1, uint8(b[len(b)-1]), stack}
  }
}

var TestPatterns [][]byte = [][]byte{
  {1,1,1,1,0,1},
  {1,1,1,0,0,1},
  {1,1,0,1,0,1},
  {1,1,0,0,0,1},
  {1,0,1,1,0,1},
  {1,0,1,0,0,1},
  {1,0,0,1,0,1},
  {1,0,0,0,0,1},
  {0,1,1,1,0,1},
  {0,1,1,0,0,1},
  {0,1,1,0,0,0},
}
