package finiteAutomata
import (
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
    return &roomC{}, nil
  case uint8(1):
    return &roomB{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomB struct{}
func (r *roomB) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomD{}, nil
  case uint8(1):
    return &roomE{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomC struct{}
func (r *roomC) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  case uint8(1):
    return &roomF{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomD struct{}
func (r *roomD) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &roomF{}, nil
  case uint8(1):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomE struct{}
func (r *roomE) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return &exit{}, nil
  case uint8(1):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type roomF struct{}
func (r *roomF) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  case uint8(1):
    return &exit{}, nil
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

type exit struct{}
func (r *exit) openDoor(i int,v uint8) (interface{}, error){
  switch v{
  case uint8(0):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  case uint8(1):
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  default:
    return 0, common.ValidationError{"invalid input", 100, r, i, v}
  }
}

/*func input(r *room, v uint8){
  o := r.OpenDoor(v)
  fmt.Println
}*/

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
    }
  }
  if _, ok:=state.(*exit);ok{
    return nil
  }else{
    return common.ValidationError{"failed to reach exit", 101, state, len(b)-1, uint8(b[len(b)-1])}
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
