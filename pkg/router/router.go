package router

import (
	"net/http"

	"github.com/auth0-developer-hub/api_standard-library_golang_hello-world/pkg/middleware"
)

func Router() http.Handler {

	router := http.NewServeMux()

	router.HandleFunc("/", middleware.NotFoundHandler)
	router.HandleFunc("/api/messages/public", middleware.PublicApiHandler)
	router.Handle("/api/messages/protected", middleware.ValidateJWT()(http.HandlerFunc(middleware.ProtectedApiHandler)))
	router.Handle("/api/messages/admin", middleware.ValidateJWT()(http.HandlerFunc(middleware.AdminApiHandler)))
 
	return middleware.HandleCacheControl(router)
}
