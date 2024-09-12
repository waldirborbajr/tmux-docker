package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tmux-docker <remote_ip:port>")
		os.Exit(1)
	}

	remoteHost := os.Args[1]

	for {
		fmt.Print(getContainerNames(remoteHost))
	}
}

func getContainerNames(host string) string {
	cli, err := client.NewClientWithOpts(
		client.WithHost(fmt.Sprintf("tcp://%s", host)),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return fmt.Sprintf("Error connecting to Docker: %v", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return fmt.Sprintf("Error listing containers: %v", err)
	}

	var names []string
	for _, container := range containers {
		names = append(names, strings.TrimPrefix(container.Names[0], "/"))
	}

	if len(names) == 0 {
		return "No running containers"
	}

	return strings.Join(names, ", ")
}
