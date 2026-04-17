package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
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

	reporter = nextools.NewErrorReporterFromEnv(os.Getenv("APP.NAME"), os.Getenv("APP.VERSION"), os.Getenv("APP.ENV"))
	logger   = nextools.NewLoggerBuilder().WithErrorReporter(reporter).Build()
)

func init() {
	logger.Banner()
}

func main() {
	defer reporter.Close()
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
	mux.Handle(grpchealth.NewHandler(middleware.HealthChecker(v1connect.IdentityServiceName)))

	validateInterceptor := validate.NewInterceptor()
	loggingInterceptor := middleware.UnaryLoggingInterceptor()
	authenticationInterceptor := middleware.UnaryTokenInterceptor()

	interceptors := connect.WithInterceptors(authenticationInterceptor, validateInterceptor, loggingInterceptor)

	path, handler := v1connect.NewIdentityServiceHandler(server, interceptors)
	mux.Handle(path, middleware.CorsMiddleware(handler))

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: h2c.NewHandler(mux, cfg.Http2()),
	}
	go func() {
		logger.Info("🚀🚀 gRPC server started on port: 🚀🚀", nextools.F("serverAddr", serverAddr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("❌ failed to start server", nextools.Err(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server gracefully...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", nextools.Err(err))
	}
}
