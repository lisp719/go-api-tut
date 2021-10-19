package main

import (
	"go-api-tut/pkg/core"
	"go-api-tut/pkg/router"
)

func main() {
	core.SetupDb()

	r := router.SetupRouter()
	r.Run()
}
