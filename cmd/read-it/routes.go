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

	// All about articles
	r.HandleFunc("/articles", app.listArticlesHandler).Methods("GET")
	r.HandleFunc("/articles", app.createArticleHandler).Methods("POST")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.getArticleHandler).Methods("GET")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.updateArticleHandler).Methods("PUT")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.deleteArticleHandler).Methods("DELETE")

	return r
}
