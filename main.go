package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

// Main function
func main() {
	// Load server configuration and password from .tmux-docker-env file
	server, password := getServerFromEnv()
	// Connect to the remote server via SSH
	client, session, err := connectToServer(server, password)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()
	defer session.Close()
	// Execute Docker command and get output
	output, err := getDockerInfo(session)
	if err != nil {
		log.Fatalf("Failed to get Docker information: %v", err)
	}
	// Parse Docker output and count containers in different states
	totalContainers, upContainers, downContainers, diedContainers := parseDockerOutput(output)
	// Display information in tmux status bar
	displayToTmux(totalContainers, upContainers, downContainers, diedContainers)
}

// Function to load user, server IP, and password from .tmux-docker-env
func getServerFromEnv() (string, string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user's home directory: %v", err)
	}

	// Load variables from .tmux-docker-env file
	err = godotenv.Load(filepath.Join(homeDir, ".tmux-docker-env"))
	if err != nil {
		log.Fatalf("Error loading .tmux-docker-env file: %v", err)
	}

	// Get user, server IP, and password from .tmux-docker-env
	user := os.Getenv("DOCKER_USER")
	serverIP := os.Getenv("REMOTE_SERVER_IP")
	password := os.Getenv("DOCKER_PASSWORD")
	if user == "" || serverIP == "" || password == "" {
		log.Fatalf("Environment variables DOCKER_USER, REMOTE_SERVER_IP, or DOCKER_PASSWORD are not set in .tmux-docker-env file")
	}

	// Return the format user@serverIP and password
	return user + "@" + serverIP, password
}

// Function to connect to server via SSH
func connectToServer(serverWithUser, password string) (*ssh.Client, *ssh.Session, error) {
	// Separate user from server IP
	parts := strings.Split(serverWithUser, "@")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid format for serverWithUser: %s", serverWithUser)
	}
	user := parts[0]
	serverIP := parts[1]

	sshConfig := &ssh.ClientConfig{
		User: user, // Use only the user
		Auth: []ssh.AuthMethod{
			ssh.Password(password), // We use the password from .tmux-docker-env file
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect using only the server IP
	client, err := ssh.Dial("tcp", serverIP+":22", sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

// Function to execute Docker command on remote server via SSH
func getDockerInfo(session *ssh.Session) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("docker ps -a --format '{{.ID}} {{.Status}}'"); err != nil {
		return "", err
	}
	return b.String(), nil
}

// Function to parse Docker output and count containers
func parseDockerOutput(output string) (int, int, int, int) {
	lines := strings.Split(output, "\n")
	total := len(lines) - 1 // Remove the last empty line
	up, down, died := 0, 0, 0

	for _, line := range lines {
		if strings.Contains(line, "Up") {
			up++
		} else if strings.Contains(line, "Exited") {
			down++
		} else if strings.Contains(line, "Dead") {
			died++
		}
	}

	return total, up, down, died
}

// Function to display results in tmux status bar
func displayToTmux(total, up, down, died int) {
	cmd := exec.Command("tmux", "set-option", "-g", "status-right", fmt.Sprintf("Total: %d | Up: %d | Down: %d | Died: %d", total, up, down, died))
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to update tmux: %v", err)
	}
}
