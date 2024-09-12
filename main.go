package main

import (
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: tmux-docker <remote_host>")
		os.Exit(1)
	}
	remoteHost := os.Args[1]
	fmt.Print(getContainerCounts(remoteHost))
}

func getContainerCounts(host string) string {
	cli, err := client.NewClientWithOpts(
		client.WithHost(fmt.Sprintf("tcp://%s", host)),
		client.WithAPIVersionNegotiation(),
		client.WithTimeout(5*time.Second),
	)
	if err != nil {
		return fmt.Sprintf("Docker conn error")
	}
	defer cli.Close()

	ctx := context.Background()

	runningContainers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return fmt.Sprintf("Docker list error")
	}
	running := len(runningContainers)

	allContainers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return fmt.Sprintf("Docker list all error")
	}
	total := len(allContainers)

	diedFilter := filters.NewArgs()
	diedFilter.Add("status", "exited")
	diedContainers, err := cli.ContainerList(ctx, container.ListOptions{All: true, Filters: diedFilter})
	if err != nil {
		return fmt.Sprintf("Docker list died error")
	}
	died := len(diedContainers)

	down := total - running - died

	return fmt.Sprintf("R:%d D:%d X:%d", running, down, died)
}
