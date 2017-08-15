package finiteAutomata
import (
  "fmt"
  "../../common"
)

type room interface{
  openDoor(i int,v uint8) (interface{}, error)
}

type entrance struct{}
func (r *entrance) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomA{}, nil
  case uint8(1):
    return &roomB{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomA struct{}
func (r *roomA) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &entrance{}, nil
  case uint8(1):
    return &roomC{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomB struct{}
func (r *roomB) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomC{}, nil
  case uint8(1):
    return &entrance{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomC struct{}
func (r *roomC) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomB{}, nil
  case uint8(1):
    return &roomA{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

func Validate(b []byte) error{
  var state room
  state = &entrance{}
  for i,v := range b{
    result, err := state.openDoor(i, uint8(v))
    if err != nil{
      return err
    }
    switch tmp := result.(type){
    case *roomA:
      state = tmp
    case *roomB:
      state = tmp
    case *roomC:
      state = tmp
    case *entrance:
      state = tmp
    case error:
      fmt.Println(tmp)
    }
  }
  if _, ok:=state.(*entrance);ok{
    return nil
  }else{
    return common.ValidationError{"failed to reach exit", 101, state, len(b)-1, uint8(b[len(b)-1])}
  }
}

var TestPatterns [][]byte = [][]byte{
  {0,0},
  {1,1},
  {0,0,1,1},
  {1,0,1,0},
  {0,1,1,1,0,1},
  {1,1,1,0,1,0},
  {0,0,1,1,1,0,1,0},
  {},
  {0,0,1,1,1,0,1,0,1,0},
  {0,0,1,1,1,0},
}
