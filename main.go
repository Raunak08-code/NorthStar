package main

import (
	"fmt"
	"log"

	"github.com/Raunak08-code/NorthStar/internal/config"
	"github.com/Raunak08-code/NorthStar/internal/logger"
)

func banner() {
	fmt.Println("NorthStar v1")
	fmt.Println(" Insights for Every Development Journey.")
	fmt.Println("========================================")
}

func main() {

	banner()

	cfg := config.Load()

	logger.Initialize()

	log.Println("Configuration Loaded")

	log.Printf("Server Port : %s\n", cfg.Port)

	log.Println("NorthStar Started")
}
