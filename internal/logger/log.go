package logger

import (
	"log"
	"os"
)

func Initialize() {

	log.SetOutput(os.Stdout)

	log.SetFlags(log.Ldate | log.Ltime)

	log.Println("✓ Logger Initialized")
}