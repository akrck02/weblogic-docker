package main

import (
	"akrck02/docker-generator/cmd"
	"akrck02/docker-generator/io"
	"fmt"
)

func main() {

	// Read the env file
	io.LoadEnv()

	fmt.Println("------------------------------------")
	fmt.Println(" Docker image ")
	fmt.Println("------------------------------------")
	// Create the image
	cmd.CreateImage()

	fmt.Println("------------------------------------")
	fmt.Println(" Docker container ")
	fmt.Println("------------------------------------")

	// Create the container
	cmd.CreateContainer()

}
