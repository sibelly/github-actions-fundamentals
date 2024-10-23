package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hueheuheuehueheu hotfix 1")
	godotenv.Load("./.env")
	fmt.Println(os.Getenv("TEST_ENV_FILE_VAR"))
}
