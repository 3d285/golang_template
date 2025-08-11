package main

import (
	"log"
	"golang_template/internal/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatalf("server start failed: %v", err)
	}
}
