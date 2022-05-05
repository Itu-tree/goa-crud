// Code generated by goa v3.7.3, DO NOT EDIT.
//
// todo service
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"
)

// Service that manage todo.
type Service interface {
	// Hello implements hello.
	Hello(context.Context, *HelloPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "todo"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"hello"}

// HelloPayload is the payload type of the todo service hello method.
type HelloPayload struct {
	// Name
	Name string
}
