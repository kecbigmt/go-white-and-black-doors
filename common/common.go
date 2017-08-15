package common
import (
  "errors"
  "fmt"
)

func Error(r interface{}, v uint8) error{
  s := fmt.Sprintf("->Error! state=%v(%T), input=%v(%T).\n",r, r, v, v)
  return errors.New(s)
}
