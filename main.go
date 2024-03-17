package main

import (
	"flag"
	"log"
	"os"

	"github.com/ayn2op/discordo/cmd"
	"github.com/ayn2op/discordo/internal/constants"
	"github.com/zalando/go-keyring"
)

func main() {
	// Fetch config from environment variables
	tokenEnv := os.Getenv("DISCORDO_TOKEN")

	// Fetch config from CLI flags
	tokenFlag := flag.String("token", "", "The authentication token.")
	flag.Parse()

	// Merge configs
	var token string
	if tokenEnv != "" { token = tokenEnv }
	if *tokenFlag != "" { token = *tokenFlag }

	// If no token was provided, look it up in the keyring
	if token == "" {
		t, err := keyring.Get(constants.Name, "token")
		if err != nil {
			log.Println("Authentication token not found in keyring:", err)
		} else {
			token = t
		}
	}

	if err := cmd.Run(token); err != nil {
		log.Fatal(err)
	}
}
