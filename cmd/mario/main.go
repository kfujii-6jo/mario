package main

import (
	"mario/internal/app/routers"
)

func main() {
    internal := routers.SetupRouter()
    internal.Run(":8080")
}
