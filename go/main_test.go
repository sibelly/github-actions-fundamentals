package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGeEnvFromGithub(t *testing.T) {
	godotenv.Load("../.env")
	test := os.Getenv("ENV_VAR")
	fmt.Println(test)
	require.NotNil(t, test)
}
