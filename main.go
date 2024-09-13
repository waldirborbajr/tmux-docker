package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

// Função principal
func main() {
	// Carregar as configurações do servidor e senha do arquivo .env
	server, password := getServerFromEnv()

	// Conectar ao servidor remoto via SSH
	client, session, err := connectToServer(server, password)
	if err != nil {
		log.Fatalf("Falha ao conectar ao servidor: %v", err)
	}
	defer client.Close()
	defer session.Close()

	// Executar o comando Docker e obter a saída
	output, err := getDockerInfo(session)
	if err != nil {
		log.Fatalf("Falha ao obter informações do Docker: %v", err)
	}

	// Analisar a saída do Docker e contar os containers em diferentes estados
	totalContainers, upContainers, downContainers, diedContainers := parseDockerOutput(output)

	// Exibir as informações na barra de status do tmux
	displayToTmux(totalContainers, upContainers, downContainers, diedContainers)
}

// Função para carregar o usuário, IP do servidor e senha a partir do .env
func getServerFromEnv() (string, string) {
	// Carregar as variáveis do arquivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Pegar o usuário, IP do servidor e senha do .env
	user := os.Getenv("DOCKER_USER")
	serverIP := os.Getenv("REMOTE_SERVER_IP")
	password := os.Getenv("DOCKER_PASSWORD")

	if user == "" || serverIP == "" || password == "" {
		log.Fatalf("Variáveis de ambiente DOCKER_USER, REMOTE_SERVER_IP ou DOCKER_PASSWORD não estão definidas no arquivo .env")
	}

	// Retornar o formato user@serverIP e a senha
	return user + "@" + serverIP, password
}

// Função para conectar ao servidor via SSH
func connectToServer(serverWithUser, password string) (*ssh.Client, *ssh.Session, error) {
	// Separar o usuário do IP do servidor
	parts := strings.Split(serverWithUser, "@")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("formato inválido para serverWithUser: %s", serverWithUser)
	}
	user := parts[0]
	serverIP := parts[1]

	sshConfig := &ssh.ClientConfig{
		User: user, // Usar apenas o usuário
		Auth: []ssh.AuthMethod{
			ssh.Password(password), // Usamos a senha do arquivo .env
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Conectar usando apenas o IP do servidor
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

// Função para executar o comando Docker no servidor remoto via SSH
func getDockerInfo(session *ssh.Session) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("docker ps -a --format '{{.ID}} {{.Status}}'"); err != nil {
		return "", err
	}
	return b.String(), nil
}

// Função para analisar a saída do Docker e contar os containers
func parseDockerOutput(output string) (int, int, int, int) {
	lines := strings.Split(output, "\n")
	total := len(lines) - 1 // Remove a última linha vazia
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

// Função para exibir os resultados na status bar do tmux
func displayToTmux(total, up, down, died int) {
	cmd := exec.Command("tmux", "set-option", "-g", "status-right", fmt.Sprintf("Total: %d | Up: %d | Down: %d | Died: %d", total, up, down, died))
	if err := cmd.Run(); err != nil {
		log.Fatalf("Falha ao atualizar o tmux: %v", err)
	}
}
