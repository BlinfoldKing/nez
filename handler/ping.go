package handler

import "errors"

func Ping() Result {
	return Ok(true)
}

func PingErr() Result {
	return Err(errors.New("hello error"))
}
