package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"strconv"

	"github.com/Nursat22B030486/go_project/pkg/read-it/model"
	"github.com/Nursat22B030486/go_project/pkg/read-it/validator"
	"github.com/gorilla/mux"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
}

func (app *application) listArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Genre string
		model.Filters
	}

	// Initialize a new Validator instance.
	v := validator.New()

	// Call r.URL.Query() to get the url.Values map containing the query string data.
	qs := r.URL.Query()

	// / Use our helpers to extract the title and genres query string values, falling back
	// to defaults of an empty string and an empty slice respectively if they are not
	// provided by the client.
	input.Title = app.readString(qs, "title", "")

	input.Genre = app.readString(qs, "genre", "")

	// Get the page and page_size query string values as integers. Notice that we set
	// the default page value to 1 and default page_size to 20, and that we pass the
	// validator instance as the final argument here.
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.Page_size = app.readInt(qs, "page_size", 20, v)

	// Extract the sort query string value, falling back to "id" if it is not provided
	// by the client (which will imply a ascending sort on movie ID).
	input.Filters.Sort = app.readString(qs, "sort", "id")

	// We can sort by this fields
	input.Filters.SortSafelist = []string{"id", "title", "created_at", "-id", "-title", "-created_at"}

	if model.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	articles, metadata, err := app.models.Articles.GetAll(input.Title, input.Genre, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send articles as a Jason format
	err = app.writeJSON(w, http.StatusOK, envelope{"articles": articles, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		AuthorId uint `json:"author_id"`
		Genre string `json:"genre"`
		Body  string `json:"body"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	article := &model.Article{
		Title: input.Title,
		AuthorId: input.AuthorId,
		Genre: input.Genre,
		Body:  input.Body,
	}

	err = app.models.Articles.Insert(article)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		fmt.Println(err)
		return
	}

	app.respondWithJSON(w, http.StatusCreated, article)
}

func (app *application) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["articleId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	article, err := app.models.Articles.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, article)
}

func (app *application) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["articleId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	article, err := app.models.Articles.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	// TODO: add users or authors
	var input struct {
		Title *string `json:"title"`
		// AuthorId *uint   `json:"author_id"`
		Genre *string `json:"genre"`
		Body  *string `json:"body"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if input.Title != nil {
		article.Title = *input.Title
	}

	// if input.AuthorId != nil {
	// 	article.AuthorId = *input.AuthorId
	// }

	if input.Genre != nil {
		article.Genre = *input.Genre
	}

	if input.Body != nil {
		article.Body = *input.Body
	}

	err = app.models.Articles.Update(article)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		fmt.Println(err)
		return
	}

	app.respondWithJSON(w, http.StatusOK, article)
}

func (app *application) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["articleId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	err = app.models.Articles.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
