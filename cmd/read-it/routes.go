package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	// Convert the app.notFoundResponse helper to a http.Handler using the http.HandlerFunc()
	// adapter, and then set it as the custom error handler for 404 Not Found responses.
	r.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)

	// Convert app.methodNotAllowedResponse helper to a http.Handler and set it as the custom
	// error handler for 405 Method Not Allowed responses
	r.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	//healthCheck
	r.HandleFunc("/healthcheck", app.healthcheckHandler).Methods("GET")

	// All about users
	r.HandleFunc("/register", app.registerUserHandler).Methods("POST")
	r.HandleFunc("/activate", app.activateUserHandler).Methods("PUT")
	r.HandleFunc("/login", app.createAuthenticationTokenHandler).Methods("POST")

	// All about articles
	r.HandleFunc("/articles", app.requirePermission("articles:read", app.listArticlesHandler)).Methods("GET")
	r.HandleFunc("/articles", app.requirePermission("articles:write", app.createArticleHandler)).Methods("POST")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.requirePermission("articles:read", app.getArticleHandler)).Methods("GET")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.requirePermission("articles:write", app.updateArticleHandler)).Methods("PUT")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.requirePermission("articles:write", app.deleteArticleHandler)).Methods("DELETE")

	// All about comments
	r.HandleFunc("/articles/{articleId:[0-9]+}/comments", app.requirePermission("articles:read", app.listArticleCommentsHandler)).Methods("GET")
	r.HandleFunc("/articles/{articleId:[0-9]+}/comments", app.requirePermission("articles:read", app.createArticlesCommentHandler)).Methods("POST")
	r.HandleFunc("/articles/{articleId:[0-9]+}/comments/{commentId:[0-9]+}", app.requirePermission("articles:read", app.getArticleCommentHandler)).Methods("GET")
	r.HandleFunc("/articles/{articleId:[0-9]+}/comments/{commentId:[0-9]+}", app.requirePermission("articles:read", app.updateArticleCommentHandler)).Methods("PUT")
	r.HandleFunc("/articles/{articleId:[0-9]+}/comments/{commentId:[0-9]+}", app.requirePermission("articles:read", app.deleteArticleCommentHandler)).Methods("DELETE")

	// Wrap the router with the panic recovery middleware.
	return app.authenticate(r)
}
