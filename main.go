package main

import (
	"os"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("")
)
}
