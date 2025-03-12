package main

import (
	"fmt"

	"defi-swap-backend/config"
	"defi-swap-backend/routes"
)

func main() {
	config.LoadConfig()
	r := routes.SetupRouter()

	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}
