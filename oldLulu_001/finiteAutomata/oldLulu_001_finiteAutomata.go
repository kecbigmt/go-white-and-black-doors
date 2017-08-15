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
    return &roomC{}
  case uint8(1):
    return &roomB{}
  default:
    return common.Error(r,v)
  }
}

type roomB struct{}
func (r *roomB) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomD{}
  case uint8(1):
    return &roomE{}
  default:
    return common.Error(r,v)
  }
}

type roomC struct{}
func (r *roomC) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return common.Error(r,v)
  case uint8(1):
    return &roomF{}
  default:
    return common.Error(r,v)
  }
}

type roomD struct{}
func (r *roomD) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &roomF{}
  case uint8(1):
    return common.Error(r,v)
  default:
    return common.Error(r,v)
  }
}

type roomE struct{}
func (r *roomE) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return &exit{}
  case uint8(1):
    return common.Error(r,v)
  default:
    return common.Error(r,v)
  }
}

type roomF struct{}
func (r *roomF) openDoor(v uint8) interface{}{
  switch v{
  case uint8(0):
    return common.Error(r,v)
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
    return common.Error(r,v)
  case uint8(1):
    return common.Error(r,v)
  default:
    return common.Error(r,v)
  }
}

/*func input(r *room, v uint8){
  o := r.OpenDoor(v)
  fmt.Println
}*/

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
    case *roomD:
      state = tmp
    case *roomE:
      state = tmp
    case *roomF:
      state = tmp
    case *entrance:
      state = tmp
    case *exit:
      state = tmp
    case error:
      fmt.Println(tmp)
      return
    case string:
      fmt.Println(tmp)
      return
    }
  }
  if _, ok:=state.(*exit);ok{
    fmt.Printf("->OK! state=%v(%T)\n", state, state)
  }else{
    fmt.Printf("->Failed. state=%v(%T)\n", state, state)
  }
}

var TestPatterns [][]byte = [][]byte{
  {1,1,0},
  {0,1,0,0,1},
  {1,0,0,1},
  {0,1,1,0},
  {0,0,1,1},
  {0,1,1,1},
}
