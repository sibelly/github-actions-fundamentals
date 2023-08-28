package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("./.env")
	fmt.Println(os.Getenv("ENV_VAR"))
}
