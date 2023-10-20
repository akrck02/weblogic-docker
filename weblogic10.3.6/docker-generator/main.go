package main

import (
	"akrck02/docker-generator/cmd"
	"akrck02/docker-generator/io"
)

func main() {

	// Read the env file
	io.LoadEnv()

	// Create the image
	cmd.CreateImage()

	// Create the container
	cmd.CreateContainer()

}
