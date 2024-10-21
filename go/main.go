package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("testing!!")
	godotenv.Load("./.env")
	fmt.Println(os.Getenv("TEST_ENV_FILE_VAR"))
}
