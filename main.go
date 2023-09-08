package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func init() {
	setup()
}

func main() {
	fmt.Printf("Server Started on PORT: %v", os.Getenv("PORT"))
}
