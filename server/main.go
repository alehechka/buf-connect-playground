package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alehechka/buf-connect-playground/proto/gen/users/v1/usersv1connect"
	"github.com/alehechka/buf-connect-playground/services/users"
	"github.com/alehechka/buf-connect-playground/utils"
	"github.com/alehechka/buf-connect-playground/utils/database"
	"github.com/alehechka/buf-connect-playground/utils/grpc"
	"github.com/alehechka/buf-connect-playground/utils/middleware"
	"github.com/alehechka/buf-connect-playground/utils/otel"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	shutdownTracer, err := otel.InitializeOpenTelTracer()
	utils.Check(err)
	defer shutdownTracer()

	disconnect, err := database.InitializeMongoDB(os.Getenv("MONGODB_CONNECTION_STRING"), "users")
	utils.Check(err)
	defer disconnect()

	api := middleware.ServeConnect(usersv1connect.NewUsersServiceHandler(users.NewServer()))
	fs := http.FileServer(http.Dir("./client"))
	otelServiceName := os.Getenv("OTEL_SERVICE_NAME")

	mux := http.NewServeMux()

	// serves default grpc endpoint and falls back to serving client
	mux.Handle("/", middleware.AttachOpenTelemetry(middleware.AttachConnectFallback(api, fs), otelServiceName))

	// serves `/api/` prefixed grpc endpoint for client
	mux.Handle("/api/", middleware.AttachOpenTelemetry(http.StripPrefix("/api", api), otelServiceName))

	listenOn := ":" + os.Getenv("PORT")
	fmt.Println("Listening on ", listenOn)
	http.ListenAndServe(listenOn, h2c.NewHandler(grpc.NewCORS().Handler(mux), &http2.Server{}))
}
