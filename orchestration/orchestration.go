package orchestration

import (
	"context"

	"github.com/teamjorge/betty/structures"

	"github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// GetRunningContainers finds a list of current Docker containers
func GetRunningContainers() []structures.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var result []structures.Container

	for _, container := range list {
		result = append(
			result,
			structures.Container{
				ID:     container.ID,
				Image:  container.Image,
				Names:  container.Names,
				Status: container.Status,
			},
		)
	}

	return result
}

// GetContainerByName finds a list of current Docker containers
func GetContainerByName(name string) []structures.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	filters := filters.NewArgs()
	filters.Add("name", name)

	list, err := cli.ContainerList(
		ctx,
		types.ContainerListOptions{
			Filters: filters,
		},
	)
	if err != nil {
		panic(err)
	}

	var result []structures.Container

	for _, container := range list {
		result = append(
			result,
			structures.Container{
				ID:     container.ID,
				Image:  container.Image,
				Names:  container.Names,
				Status: container.Status,
			},
		)
	}

	return result
}

// ListNetworks returns a list of current Docker networks
func ListNetworks() map[string]string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	list, err := cli.NetworkList(ctx, types.NetworkListOptions{})

	var networkMap map[string]string
	networkMap = make(map[string]string)

	for _, network := range list {
		networkMap[network.ID] = network.Name
	}

	return networkMap
}

// StartContainer will create and run a Docker container
func StartContainer(containerConfig structures.StartContainerRequest) container.ContainerCreateCreatedBody {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image:           containerConfig.Image,
			Tty:             true,
			Env:             containerConfig.Environment,
			NetworkDisabled: false,
		},
		&container.HostConfig{
			NetworkMode: container.NetworkMode(containerConfig.Network),
		},
		nil,
		containerConfig.Name,
	)

	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	return resp
}
