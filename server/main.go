package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"buf-connect-playground/proto/gen/users/v1/usersv1connect"

	"buf-connect-playground/services/users"

	"buf-connect-playground/utils"

	"buf-connect-playground/utils/database"

	"buf-connect-playground/utils/grpc"

	"buf-connect-playground/utils/middleware"

	"buf-connect-playground/utils/otel"

	"github.com/gin-gonic/gin"
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

	listenOn := ":" + os.Getenv("PORT")
	fmt.Println("Listening on ", listenOn)
	http.ListenAndServe(listenOn, NewHandler())

}

type Handler struct {
	ginHandler  *gin.Engine
	grpcHandler http.Handler
}

func NewHandler() *Handler {
	router := gin.Default()
	router.GET("/rest/hello", func(ctx *gin.Context) { ctx.Data(http.StatusOK, "text/plain", []byte("world")) })
	router.POST("/rest/generate/:numUsers", users.GenerateUsers)

	api := middleware.ServeConnect(usersv1connect.NewUsersServiceHandler(users.NewServer()))
	fs := http.FileServer(http.Dir("./client"))
	otelServiceName := os.Getenv("OTEL_SERVICE_NAME")

	mux := http.NewServeMux()

	// serves default grpc endpoint and falls back to serving client
	mux.Handle("/", middleware.AttachConnectFallback(api, fs))

	// serves `/api/` prefixed grpc endpoint for client
	mux.Handle("/api/", http.StripPrefix("/api", api))

	grpcHandler := h2c.NewHandler(grpc.NewCORS().Handler(middleware.AttachOpenTelemetry(mux, otelServiceName)), &http2.Server{})

	return &Handler{
		ginHandler:  router,
		grpcHandler: grpcHandler,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/rest") {
		h.ginHandler.ServeHTTP(w, req)
		return
	}
	h.grpcHandler.ServeHTTP(w, req)
}
