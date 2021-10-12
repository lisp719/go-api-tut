package main

import (
	"go-api-tut/core"
	"go-api-tut/router"
)

func main() {
	core.SetupDb()

	r := router.SetupRouter()
	r.Run(":8080")
}
