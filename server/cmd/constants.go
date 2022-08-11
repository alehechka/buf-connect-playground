package cmd

import (
	"buf-connect-playground/utils"
	"fmt"
)

const UserID string = "62f51d6513c7398035bbca9f"

var GrpcServerHost string = utils.GetEnv("GRPC_SERVER_HOST", "127.0.0.1:3000")

var ConnectServerHost string = fmt.Sprintf("http://%s/api", utils.GetEnv("GRPC_SERVER_HOST", "127.0.0.1:3000"))
