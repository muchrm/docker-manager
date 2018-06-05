package main

import (
	"context"
	"fmt"
	"time"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
)

var masterNode = []string{}
var slaveNode = []string{}

func genContainerList(cli *docker.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		networkMode := container.HostConfig.NetworkMode
		if networkMode == "default" {
			networkMode = "bridge"
		}
		network := container.NetworkSettings.Networks[networkMode]
		aliase := ""
		if len(network.Aliases) > 0 {
			aliase = network.Aliases[0]
		}
		fmt.Printf("%s %s %s %s %s\n",
			container.ID[:10],
			container.Names[0][1:],
			networkMode,
			aliase,
			network.IPAddress,
		)
	}
}
func main() {
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	eventChan, _ := cli.Events(context.Background(), types.EventsOptions{})
	for {
		event := <-eventChan
		if event.Type == "container" && event.Action == "start" {
			fmt.Printf("event: %s container: %s\ncurrent container:\n", event.Action, event.Actor.ID[:10])
			genContainerList(cli)
		} else if event.Type == "network" && event.Action == "disconnect" {
			fmt.Printf("event: %s container: %s\ncurrent container:\n", event.Action, event.Actor.ID[:10])
			time.Sleep(100 * time.Millisecond)
			genContainerList(cli)
		}
	}
}
