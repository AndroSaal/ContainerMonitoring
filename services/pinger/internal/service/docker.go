package service

import (
	"context"
	"log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	cli *client.Client
	log slog.Logger
}

func NewDockerCllient() *DockerClient {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic("cannot connect to docker client: " + err.Error())
	}
	return &DockerClient{
		cli: cli,
	}
}

func (d *DockerClient) GetAllContainersIP(ctx context.Context) ([]string, error) {
	containers, err := d.cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}
	var ips []string
	for _, container := range containers {
		containerJSON, err := d.cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			d.log.Error("upexpected error: ", "error", err)
			continue
		}

		for _, network := range containerJSON.NetworkSettings.Networks {
			ips = append(ips, network.IPAddress)
		}
	}

	return ips, nil
}
