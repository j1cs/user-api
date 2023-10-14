package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	cm "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/j1cs/api-user/internal/v1/port"
	m "github.com/j1cs/api-user/internal/v1/service/middleware"
)

const VERSION = "/v1"
const AUX = "/aux"

func RunServer(server *port.Server, logger *zerolog.Logger) error {
	// parent
	r := chi.NewRouter()

	// aux for health
	auxRouter := chi.NewRouter()
	auxRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(nil)
	})

	// api v1
	apiRouter := chi.NewRouter()
	middleWares(apiRouter, logger)
	apiRouter.Group(func(r chi.Router) {
		strictMode := port.NewStrictHandlerWithOptions(server, nil, configOptions())
		port.HandlerFromMux(strictMode, r)
	})

	apiPort := os.Getenv("API_PORT")
	// checking port
	checkPort(apiPort, logger)

	// mounting APIs
	r.Mount(AUX, auxRouter)
	r.Mount(VERSION, apiRouter)
	logger.Info().Msg(fmt.Sprintf("server start at PORT: %s ", apiPort))
	return http.ListenAndServe(":"+apiPort, r)
}

func checkPort(port string, logger *zerolog.Logger) {
	if len(port) <= 0 {
		logger.Err(errors.New("please add API_PORT with valid value. i.e: \"8080\"")).Msg("Config Error")
		os.Exit(1)
	}
}

// https://github.com/deepmap/oapi-codegen/pull/459
func writeJSONError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := port.Error{
		Code:    statusCode,
		Message: strings.TrimSpace(err.Error()),
	}
	json.NewEncoder(w).Encode(errResponse)
}

func configOptions() port.StrictHTTPServerOptions {
	return port.StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			writeJSONError(w, http.StatusBadRequest, err)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			writeJSONError(w, http.StatusInternalServerError, err)
		},
	}
}

func addCors(router *chi.Mux, logger *zerolog.Logger) {
	allowed := strings.Split(os.Getenv("CORS_ALLOWED_ORIGIN"), ";")
	if len(allowed) == 0 {
		logger.Warn().Msg("Cors origins empty")
	}
	headersAllowed := strings.Split(os.Getenv("CORS_ALLOWED_HEADER"), ";")
	if len(headersAllowed) == 0 {
		logger.Warn().Msg("Cors header empty")
	}

	isDebug := false
	if strings.EqualFold(os.Getenv("CORS_DEBUG"), "true") {
		isDebug = true
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: allowed,
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: headersAllowed,
		Debug:          isDebug,
	})
	router.Use(corsMiddleware.Handler)
}

func middleWares(router *chi.Mux, logger *zerolog.Logger) {

	router.Use(middleware.RequestID)
	router.Use(m.NewStructuredLogger(*logger))
	router.Use(requestValidatorInit(logger))
	router.Use(middleware.StripSlashes)
	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("Strict-Transport-Security", "max-age=31536000; includeSubDomains"),
		middleware.SetHeader("X-Frame-Options", "SAMEORIGIN"),
		middleware.SetHeader("X-XSS-Protection", "1; mode=block"),
	)
	addCors(router, logger)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.NoCache)
}

func requestValidatorInit(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	swagger, err := port.GetSwagger()
	if err != nil {
		logger.Err(err).Msg("Getting Swagger Error")
		os.Exit(1)
	}
	op := &cm.Options{
		SilenceServersWarning: true,
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			errResponse := port.Error{
				Code:    statusCode,
				Message: strings.TrimSpace(message),
			}
			json.NewEncoder(w).Encode(errResponse)
		},
	}
	return cm.OapiRequestValidatorWithOptions(swagger, op)
}
