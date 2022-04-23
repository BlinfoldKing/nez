package handler

import (
	"errors"
	"fmt"
)

func Ping(request struct{}) Result {
	fmt.Println("Coba")
	return Ok(true)
}

func PingErr(request struct{}) Result {
	return Err(errors.New("hello error"))
}

func PingBody(request struct {
	Name string `json:"name" validate:"required"`
}) Result {
	return Ok(request.Name)
}
