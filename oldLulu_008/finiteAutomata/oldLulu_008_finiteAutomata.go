package finiteAutomata
import (
  "fmt"
  "../../common"
)

type room interface{
  openDoor(v uint8) interface{}
}

type entrance struct{}
func (r *entrance) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomA{}
  case uint8(1):
    return &roomB{}
  default:
    return common.Error(r,v)
  }
}

type roomA struct{}
func (r *roomA) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &entrance{}
  case uint8(1):
    return &roomC{}
  default:
    return common.Error(r,v)
  }
}

type roomB struct{}
func (r *roomB) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomC{}
  case uint8(1):
    return &entrance{}
  default:
    return common.Error(r,v)
  }
}

type roomC struct{}
func (r *roomC) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomB{}
  case uint8(1):
    return &roomA{}
  default:
    return common.Error(r,v)
  }
}

func Validate(b []byte){
  var state room
  state = &entrance{}
  for _,v := range b{
    switch tmp := state.openDoor(uint8(v)).(type){
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
    fmt.Printf("->OK! state=%v(%T)\n", state, state)
  }else{
    fmt.Printf("->Failed. state=%v(%T)\n", state, state)
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
