package cmd

import "fmt"

func CreateImage() {

	var cmd = createImageCommand()

	// Write the command to the console
	fmt.Println(cmd)

}

func createImageCommand() string {
	return ""
}
