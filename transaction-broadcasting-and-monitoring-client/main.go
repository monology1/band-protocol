package main

import (
	"example/band-protocol/routes"
	"example/band-protocol/services"
)

func main() {
	r := routes.Router()
	r.POST("/broadcast", routes.NewGin(services.BroadcastAndMonitorTransaction))
	err := r.Run()
	if err != nil {
		return
	}
}
