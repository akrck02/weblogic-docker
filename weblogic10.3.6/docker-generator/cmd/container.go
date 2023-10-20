package cmd

import (
	"akrck02/docker-generator/io"
	"akrck02/docker-generator/models"
	"akrck02/docker-generator/templates"
	"fmt"
)

func CreateContainer() {

	var cmd = createContainerCommand()

	// Write the command to the console
	fmt.Println(cmd)

}

func createContainerCommand() string {
	// generate the docker container create command
	var cmd = templates.Format(templates.DOCKER_CONTAINER_CREATE_CMD, map[string]string{
		"name":     getName(),
		"hostname": getHostname(),
		"ports":    getPortFlags(),
		"hosts":    getHostFlags(),
		"mounts":   getMountFlags(),
		"image":    io.Params.IMAGE_NAME,
	})

	return cmd
}

// Parse the port flags
func getPortFlags() string {

	// Create Ports
	ports := []models.Port{
		{HostPort: io.Params.HOST_HTTP_PORT, ContainerPort: io.Params.CONTAINER_HTTP_PORT},
		{HostPort: io.Params.HOST_HTTPS_PORT, ContainerPort: io.Params.CONTAINER_HTTPS_PORT},
		{HostPort: io.Params.HOST_DEBUG_PORT, ContainerPort: io.Params.CONTAINER_DEBUG_PORT},
	}
	flags := ""

	for _, port := range ports {
		flags += port.ToString() + " "
	}

	return flags
}

// Parse the host flags
func getHostFlags() string {

	// Read the hosts file
	hosts := io.ReadHosts()
	flags := ""

	for _, host := range hosts {
		flags += host.ToString() + " "
	}

	return flags
}

// Parse the mount flags
func getMountFlags() string {

	// Create Mounts

	mounts := []models.Mount{
		{HostPath: io.Params.HOST_WEBLOGIC_LOGS_PATH, ContainerPath: io.Params.CONTAINER_WEBLOGIC_LOGS_PATH},
		{HostPath: io.Params.HOST_WEBLOGIC_DOMAIN_PATH, ContainerPath: io.Params.CONTAINER_WEBLOGIC_DOMAIN_PATH},
		{HostPath: io.Params.HOST_WEBLOGIC_DEPLOY_PATH, ContainerPath: io.Params.CONTAINER_WEBLOGIC_DEPLOY_PATH},
		{HostPath: io.Params.HOST_WEBLOGIC_DATA_PATH, ContainerPath: io.Params.CONTAINER_WEBLOGIC_DATA_PATH},
	}

	flags := ""

	for _, mount := range mounts {
		mount.HostPath = templates.Format(mount.HostPath, map[string]string{"appcode": io.Params.APP_CODE})
		flags += mount.ToString() + " "
	}

	return flags
}

// Parse the name
func getName() string {
	return templates.Format(io.Params.CONTAINER_NAME,
		map[string]string{
			"appcode": io.Params.APP_CODE,
		},
	)
}

// Parse the hostname
func getHostname() string {
	return templates.Format(io.Params.HOST_NAME,
		map[string]string{
			"appcode": io.Params.APP_CODE,
		},
	)
}
