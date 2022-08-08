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
	"github.com/alehechka/buf-connect-playground/utils/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

	api := http.NewServeMux()
	api.Handle(usersv1connect.NewUsersServiceHandler(users.NewServer()))

	mux := http.NewServeMux()
	mux.Handle("/api/", otelhttp.NewHandler(http.StripPrefix("/api", api), "grpc", otelhttp.WithTracerProvider(otel.OpenTelTracer)))

	listenOn := ":" + os.Getenv("PORT")
	fmt.Println("Listening on ", listenOn)
	http.ListenAndServe(listenOn, h2c.NewHandler(grpc.NewCORS().Handler(mux), &http2.Server{}))
}
