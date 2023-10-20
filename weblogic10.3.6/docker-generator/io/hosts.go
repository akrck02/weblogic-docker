package io

import (
	models "akrck02/docker-generator/models"
	"log"
	"os"
	"strings"
)

const HOSTS_PATH = "../resources/hosts"

func ReadHosts() []models.Host {
	// Read the host file
	hostFile, err := os.ReadFile(HOSTS_PATH)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the host file to a string
	hostFileString := string(hostFile)

	// Split the host file string into an array of strings
	hostFileArray := strings.Split(hostFileString, "\n")

	// Create an array of hosts
	var hosts []models.Host

	// Iterate through the host file array
	for _, host := range hostFileArray {

		// Split the host file array into an array of strings
		hostArray := strings.Split(host, " ")

		// Create a host
		host := models.Host{Host: hostArray[0], Ip: hostArray[1]}

		// Append the host to the hosts array
		hosts = append(hosts, host)
	}

	// Return the hosts array
	return hosts
}
