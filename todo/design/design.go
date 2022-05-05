package design

import . "goa.design/goa/v3/dsl"

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service that manage todo.")
	Server("todo", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})
var _ = Service("todo", func() {
	Description("Service that manage todo.")
	Method("hello", func() {
		Payload(func() {
			Attribute("name", String, "Name")
			Required("name")
		})
		Result(String)
		HTTP(func() {
			GET("/hello/{name}")
			Response(StatusOK)
		})
	})
})
