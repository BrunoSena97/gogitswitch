package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	//get homedir path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding the user's home directory")
	}
	//define envpath
	envPath := filepath.Join(homeDir, "gitswitch/.env")

	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 3 {
		log.Fatal("Please specify the service (gitlab/github) and the key of the account that you want to use to switch credentials.")
	}

	service := os.Args[1]
	accountType := os.Args[2]
	envKey := fmt.Sprintf("%s_%s", service, accountType)
	creds := os.Getenv(envKey)

	if creds == "" {
		log.Fatalf("No credentials found for %s account type %s. Please check your .env file.", service, accountType)
	}

	// Remove quotes if present
	creds = strings.Trim(creds, `"`)

	// split username and password
	credsArray := strings.Split(creds, ":")
	if len(credsArray) != 2 {
		log.Fatal("Credentials format is invalid. It should be username:password")
	}

	username := credsArray[0]
	password := credsArray[1]

	// Set the host based on the service
	var host string
	switch service {
	case "gitlab":
		host = "gitlab.com"
	case "github":
		host = "github.com"
	default:
		log.Fatal("Invalid service. Use 'gitlab' or 'github'.")
	}

	// First, delete any existing credentials for the specified host
	deleteExistingCreds(host)

	// Set the new git credentials using the `git credential approve` command
	cmd := exec.Command("git", "credential", "approve")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("protocol=https\nhost=%s\nusername=%s\npassword=%s\n", host, username, password))
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to set credentials: %v", err)
	}

	fmt.Printf("Switched to %s account type %s.\n", service, accountType)
}

func deleteExistingCreds(host string) {
	// Use the `git credential reject` command to delete existing credentials
	cmd := exec.Command("git", "credential", "reject")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("protocol=https\nhost=%s\n", host))
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to delete existing credentials: %v", err)
	}
}
