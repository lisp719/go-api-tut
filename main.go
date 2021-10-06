package main

import (
	"go-api-tut/core"
)

func main() {
	core.SetupDb()

	r := core.SetupRouter()
	r.Run(":8080")
}
