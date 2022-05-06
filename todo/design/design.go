package design

import . "goa.design/goa/v3/dsl"

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service that manage todo.")
	Server("todo", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var UserName = ResultType("UserName", func() {
	Attributes(func() {
		Attribute("id", Int, "ID")
		Attribute("name", String, "Name")
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
	Method("show", func() {
		Payload(func() {
			Attribute("id", Int, "ID")
			Required("id")
		})
		Result(UserName)
		HTTP(func() {
			GET("/user/{id}")
			Response(StatusOK)
		})
	})
})
