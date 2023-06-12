package main

import (
	"log"
	"omni-integration-service/src/service"
)

func main() {
	log.Printf("Application is running")

	productService := service.NewOrderService()
	productService.ConsumeOrderMessages()
}
