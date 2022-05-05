// Code generated by goa v3.7.3, DO NOT EDIT.
//
// todo client
//
// Command:
// $ goa gen todo/design

package todo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "todo" service client.
type Client struct {
	HelloEndpoint goa.Endpoint
}

// NewClient initializes a "todo" service client given the endpoints.
func NewClient(hello goa.Endpoint) *Client {
	return &Client{
		HelloEndpoint: hello,
	}
}

// Hello calls the "hello" endpoint of the "todo" service.
func (c *Client) Hello(ctx context.Context, p *HelloPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.HelloEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
