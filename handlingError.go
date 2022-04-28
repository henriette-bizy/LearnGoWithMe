package main

import (
	"error"
	"errors"
	"fmt"
)

func main(name string) (string,error) {

	if name = "" {
      return "", errors.New("empty name")

	}
 message := fmt.Sprintf("Hi %v",name)
 return message, nil




}