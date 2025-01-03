package grpcerrors

import "fmt"

type Error struct {
	Code     string
	Function string
	Err      error
	notFound bool
}

func (e Error) Error() string {
	return fmt.Sprintf("error in function %s with code %s: %s", e.Function, e.Code, e.Err.Error())
}
