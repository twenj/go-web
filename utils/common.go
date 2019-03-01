package utils

import "fmt"

func Abs(data interface{}) interface{} {
  switch data.(type) {
	case int:
		fmt.Println("int")
		break
	default:
		fmt.Println("default")
	}
  return data
}
