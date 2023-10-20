package io

import (
	"os"

	"github.com/joho/godotenv"
)

const ENV_PATH = "../resources/.env"

type GlobalConfiguration struct {
	WEBLOGIC_VERSION   string
	GENERIC_IMAGE_NAME string

	CONTAINER_NAME string
	HOST_NAME      string
	IMAGE_NAME     string
	APP_CODE       string

	CONTAINER_HTTP_PORT  string
	CONTAINER_HTTPS_PORT string
	CONTAINER_DEBUG_PORT string

	HOST_HTTP_PORT  string
	HOST_HTTPS_PORT string
	HOST_DEBUG_PORT string

	CONTAINER_WEBLOGIC_LOGS_PATH string
	HOST_WEBLOGIC_LOGS_PATH      string

	CONTAINER_WEBLOGIC_DEPLOY_PATH string
	HOST_WEBLOGIC_DEPLOY_PATH      string

	CONTAINER_WEBLOGIC_DATA_PATH string
	HOST_WEBLOGIC_DATA_PATH      string

	HDIV_ENABLE      string
	HDIV_MAIL        string
	HDIV_CONSOLE_URL string
	MEMORY           string
}

var Params GlobalConfiguration

func LoadEnv() {

	// load env file
	err := godotenv.Load(ENV_PATH)
	if err != nil {
		panic(err)
	}

	Params = GlobalConfiguration{

		WEBLOGIC_VERSION:   os.Getenv("WEBLOGIC_VERSION"),
		GENERIC_IMAGE_NAME: os.Getenv("GENERIC_IMAGE_NAME"),

		CONTAINER_NAME: os.Getenv("CONTAINER_NAME"),
		HOST_NAME:      os.Getenv("HOST_NAME"),
		IMAGE_NAME:     os.Getenv("IMAGE_NAME"),
		APP_CODE:       os.Getenv("APP_CODE"),

		CONTAINER_HTTP_PORT:  os.Getenv("CONTAINER_HTTP_PORT"),
		CONTAINER_HTTPS_PORT: os.Getenv("CONTAINER_HTTPS_PORT"),
		CONTAINER_DEBUG_PORT: os.Getenv("CONTAINER_DEBUG_PORT"),

		HOST_HTTP_PORT:  os.Getenv("HOST_HTTP_PORT"),
		HOST_HTTPS_PORT: os.Getenv("HOST_HTTPS_PORT"),
		HOST_DEBUG_PORT: os.Getenv("HOST_DEBUG_PORT"),

		CONTAINER_WEBLOGIC_LOGS_PATH: os.Getenv("CONTAINER_WEBLOGIC_LOGS_PATH"),
		HOST_WEBLOGIC_LOGS_PATH:      os.Getenv("HOST_WEBLOGIC_LOGS_PATH"),

		CONTAINER_WEBLOGIC_DEPLOY_PATH: os.Getenv("CONTAINER_WEBLOGIC_DEPLOY_PATH"),
		HOST_WEBLOGIC_DEPLOY_PATH:      os.Getenv("HOST_WEBLOGIC_DEPLOY_PATH"),

		CONTAINER_WEBLOGIC_DATA_PATH: os.Getenv("CONTAINER_WEBLOGIC_DATA_PATH"),
		HOST_WEBLOGIC_DATA_PATH:      os.Getenv("HOST_WEBLOGIC_DATA_PATH"),

		HDIV_ENABLE:      os.Getenv("HDIV_ENABLE"),
		HDIV_MAIL:        os.Getenv("HDIV_MAIL"),
		HDIV_CONSOLE_URL: os.Getenv("HDIV_CONSOLE_URL"),
		MEMORY:           os.Getenv("MEMORY"),
	}

}
