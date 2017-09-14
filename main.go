package main

import (
	"voting"
)

func main() {
	println("Hello world")

	service := voting.NewService()
	service.Find("test")
}
