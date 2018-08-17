// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"go-swagger-sample/restapi/operations"
	"go-swagger-sample/handlers"
	"go-swagger-sample/common"
	"go-swagger-sample/models"
	"time"
)

//go:generate swagger generate server --target .. --name  --spec ../swagger.yml

func configureFlags(api *operations.ATodoListApplicationAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ATodoListApplicationAPI) http.Handler {
	api.ServeError = errors.ServeError

	// get handlers.
	handlers.InitializeHandlers(api)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.yml document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.yml document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return authenticationMiddleware(handler)
}

func authenticationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		consumerKey := r.Header.Get("x-consumer-key")
		consumerId := r.Header.Get("x-consumer-id")

		// if consumerKey and consumerId is empty throw error.
		if consumerKey == "" || consumerId == "" {
			common.ErrorResponse(
				w,
				4005,
				"Wrong consumer identifier and access key combination.",
				http.StatusUnauthorized,
				"Unauthorized",
			)

			// if consumerKey and consumerId is not empty and consumer exists keep continue.
		} else if consumerKey != "" && consumerId != "" {
			if len(common.GetCacheValueByKey("session_key")) > 0 {
				handler.ServeHTTP(w, r)
			} else {
				db, _ := common.InitializeDb()
				var consumer models.Consumer

				if err := db.Where("id = ? AND `key` = ?", consumerId, consumerKey).First(&consumer).Error; err != nil {
					common.ErrorResponse(
						w,
						4005,
						"Wrong consumer identifier and access key combination.",
						http.StatusUnauthorized,
						"Unauthorized",
					)
				} else {
					common.Cache{
						Key:      "session_token",
						Value:    consumerKey,
						ExprTime: 60 * time.Second}.SetCache()

					handler.ServeHTTP(w, r)
				}
			}
		}
	})
}
