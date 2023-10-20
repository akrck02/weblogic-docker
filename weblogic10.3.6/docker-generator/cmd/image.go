package cmd

import (
	"akrck02/docker-generator/io"
	"akrck02/docker-generator/models"
	"akrck02/docker-generator/templates"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const PATH = ".."

func CreateImage() {

	var cmd = createImageCommand()

	// Execute the command
	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)

	print("Executing: " + args[0] + " " + strings.Join(args[1:], " "))

	out, err := command.CombinedOutput()

	if err != nil {
		fmt.Println()
		fmt.Println(string(out))
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))

	// Write the command to the console

}

func createImageCommand() string {

	cmd := templates.Format(templates.DOCKER_IMAGE_BUILD_CMD, map[string]string{
		"arguments": getArguments(),
		"image":     io.Params.IMAGE_NAME,
		"path":      PATH,
	})

	return cmd
}

func getArguments() string {

	args := ""

	argList := []models.Arg{
		{Name: "eHDIV_ENABLE", Value: io.Params.HDIV_ENABLE},
		{Name: "eHDIV_MAIL", Value: io.Params.HDIV_MAIL},
		{Name: "eHDIV_CONSOLE_URL", Value: io.Params.HDIV_CONSOLE_URL},

		{Name: "eHOSTNAME", Value: io.Params.HOST_NAME},
		{Name: "eADMIN_PORT_SSL", Value: io.Params.CONTAINER_HTTPS_PORT},
		{Name: "eADMIN_PORT", Value: io.Params.CONTAINER_HTTP_PORT},
		{Name: "eDEBUG_PORT", Value: io.Params.CONTAINER_DEBUG_PORT},
		{Name: "eMEMORY", Value: io.Params.MEMORY},
	}

	for _, arg := range argList {
		args += arg.ToString() + " "
	}

	return args
}
