package main

import (
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	"github.com/Nerzal/gocloak/v13"
	"github.com/charmmtech/identity/gen/charmmtech/identity/v1/v1connect"
	"github.com/charmmtech/identity/internal/data/repositories"
	"github.com/charmmtech/identity/internal/servers"
	"github.com/charmmtech/identity/internal/services"
	"github.com/charmmtech/nextools"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/net/http2/h2c"
)

var (
	cfg = nextools.NewConfig(
		nextools.WithEnvValue("APP.NAME", os.Getenv("APP.NAME")),
		nextools.WithEnvValue("APP.VERSION", os.Getenv("APP.VERSION")),
		nextools.WithEnvValue("APP.ENV", os.Getenv("APP.ENV")),
		nextools.WithEnvValue("APP.DEBUG", os.Getenv("APP.DEBUG")),
	)

	logger = nextools.NewLoggerBuilder().Build()
)

func init() {
	logger.Banner()
}

func main() {
	serverAddr := cfg.ServerAddr()

	authenticator, err := nextools.NewAuthenticator()
	if err != nil {
		logger.Error("👮🏽Error connecting to authenticator:", nextools.Err(err))
		os.Exit(2)
	}

	contextHelper := nextools.NewContextHelperBuilder().
		WithAuthenticator(authenticator).
		Build()
	middleware := nextools.NewMiddlewareBuilder().
		WithContextHelper(contextHelper).
		WithAuthenticator(authenticator).
		WithLogger(logger).
		Build()

	goCloak := gocloak.NewClient(os.Getenv("KEYCLOAK.URL"))
	repo := repositories.NewKeycloakIdentityRepository(goCloak)
	service := services.NewIdentityService(repo)
	server := servers.NewIdentityServer(service)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(v1connect.IdentityServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	validateInterceptor := validate.NewInterceptor()
	loggingInterceptor := middleware.UnaryLoggingInterceptor()
	authenticationInterceptor := middleware.UnaryTokenInterceptor()

	interceptors := connect.WithInterceptors(authenticationInterceptor, validateInterceptor, loggingInterceptor)

	path, handler := v1connect.NewIdentityServiceHandler(server, interceptors)
	mux.Handle(path, middleware.CorsMiddleware(handler))

	logger.Info("🚀🚀 gRPC server started on port: 🚀🚀", nextools.F("serverAddr", serverAddr))
	if err := http.ListenAndServe(serverAddr, h2c.NewHandler(mux, cfg.Http2())); err != nil {
		log.Println("failed to start server:", err)
		os.Exit(1)
	}
}
