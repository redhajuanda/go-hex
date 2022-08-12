package main

import (
	"go-hex/cmd"
)

// @title Go Hex RESTful APIs
// @description This is a documentation for Go Hex RESTful APIs. <br>

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

func main() {
	cmd.Run()
}
