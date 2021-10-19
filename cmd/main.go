package main

import (
	"go-api-tut/pkg/core"
	"go-api-tut/pkg/router"
)

func main() {
	core.SetupDb()

	r := router.SetupRouter()

	if err := r.Run(); err != nil {
		panic("failed to start server")
	}
}
