package config

import (
	"fmt"
	"log"
	"os"
)

var API_GATEWAY_SERVER_HOST,
	API_GATEWAY_SERVER_PORT,

	SERVICE_LOGIN_HOST,
	SERVICE_LOGIN_PORT,

	SERVICE_MONTIR_HOST,
	SERVICE_MONTIR_PORT,

	SERVICE_USER_HOST,
	SERVICE_USER_PORT string

// Config server Host and Port
func SetEnvironmentVariables() {
	API_GATEWAY_SERVER_HOST = os.Getenv("API_GATEWAY_SERVER_HOST")
	API_GATEWAY_SERVER_PORT = os.Getenv("API_GATEWAY_SERVER_PORT")

	SERVICE_LOGIN_HOST = os.Getenv("SERVICE_LOGIN_HOST")
	SERVICE_LOGIN_PORT = os.Getenv("SERVICE_LOGIN_PORT")

	SERVICE_MONTIR_HOST = os.Getenv("SERVICE_MONTIR_HOST")
	SERVICE_MONTIR_PORT = os.Getenv("SERVICE_MONTIR_PORT")

	SERVICE_USER_HOST = os.Getenv("SERVICE_USER_HOST")
	SERVICE_USER_PORT = os.Getenv("SERVICE_USER_PORT")

	log.Print(`ENVIRONMENT VARIABLE`)
	fmt.Println()

	log.Println(`API_GATEWAY_SERVER_HOST=`, API_GATEWAY_SERVER_HOST)
	log.Println(`API_GATEWAY_SERVER_PORT=`, API_GATEWAY_SERVER_PORT)
	fmt.Println()

	log.Println(`SERVICE_LOGIN_HOST=`, SERVICE_LOGIN_HOST)
	log.Println(`SERVICE_LOGIN_PORT=`, SERVICE_LOGIN_PORT)
	fmt.Println()

	log.Println(`SERVICE_MONTIR_HOST=`, SERVICE_MONTIR_HOST)
	log.Println(`SERVICE_MONTIR_PORT=`, SERVICE_MONTIR_PORT)
	fmt.Println()

	log.Println(`SERVICE_USER_HOST=`, SERVICE_USER_HOST)
	log.Println(`SERVICE_USER_PORT=`, SERVICE_USER_PORT)
	fmt.Println()
}
