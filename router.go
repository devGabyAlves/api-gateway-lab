package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"gateway/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	usersURL, _ := url.Parse("http://users:3001")
	adminURL, _ := url.Parse("http://admin:3002")

	usersProxy := httputil.NewSingleHostReverseProxy(usersURL)
	adminProxy := httputil.NewSingleHostReverseProxy(adminURL)

	mux.Handle(
		"/api/users/",
		middleware.Chain(
			usersProxy,              
			middleware.Logging,      
			middleware.Auth,         
		),
	)

	mux.Handle(
		"/api/admin/",
		middleware.Chain(
			adminProxy,              
			middleware.Logging,      
			middleware.Auth,         
			middleware.RequireAdmin, 
		),
	)

	return mux
}
