package common
import "fmt"

type ValidationError struct{
  Msg string
  Code uint8
  State interface{}
  Index int
  Input uint8
}

func (e ValidationError) Error() string{
  return fmt.Sprintf("%v(%v) state=%T, index=%v, input=%v", e.Msg, e.Code, e.State, e.Index, e.Input)
}
