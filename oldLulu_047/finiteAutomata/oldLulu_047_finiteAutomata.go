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
    return &entrance{}
  default:
    return common.Error(r,v)
  }
}

type roomA struct{}
func (r *roomA) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomA{}
  case uint8(1):
    return &exit{}
  default:
    return common.Error(r,v)
  }
}

type exit struct{}
func (r *exit) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomA{}
  case uint8(1):
    return &entrance{}
  default:
    return common.Error(r,v)
  }
}

func Validate(b []byte){
  var state room
  state = &entrance{}
  for _,v := range b{
    switch tmp := state.openDoor(uint8(v)).(type){
    case *entrance:
      state = tmp
    case *roomA:
      state = tmp
    case *exit:
      state = tmp
    case error:
      fmt.Println(tmp)
    }
  }
  if _, ok:=state.(*exit);ok{
    fmt.Printf("->OK! state=%v(%T)\n", state, state)
  }else{
    fmt.Printf("->Failed. state=%v(%T)\n", state, state)
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
}
