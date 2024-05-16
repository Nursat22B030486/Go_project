package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nursat22B030486/go_project/pkg/read-it/model"
	"github.com/Nursat22B030486/go_project/pkg/read-it/validator"
	"github.com/gorilla/mux"
)

func (app *application) listArticleCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["articleId"]

	article_id, err := strconv.Atoi(param)
	if err != nil || article_id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	var input struct {
		Body string
		model.Filters
	}

	v := validator.New()

	qs := r.URL.Query()
	input.Body = app.readString(qs, "body", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.Page_size = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "created_at")

	input.Filters.SortSafelist = []string{"id", "created_at", "-id", "-created_at"}

	if model.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	comments, metadata, err := app.models.Comments.GetAll(article_id, input.Body, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"comments": comments, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createArticlesCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["articleId"]

	article_id, err := strconv.Atoi(param)
	if err != nil || article_id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article ID")
		return
	}

	var input struct {
		UserId uint   `json:"user_id"`
		Body   string `json:"body"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment := &model.Comment{
		ArticleId: uint(article_id),
		UserId:    input.UserId,
		Body:      input.Body,
	}

	err = app.models.Comments.Insert(comment)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, comment)
}

func (app *application) getArticleCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param1 := vars["articleId"]
	param2 := vars["commentId"]

	article_id, err1 := strconv.Atoi(param1)
	comment_id, err2 := strconv.Atoi(param2)

	if err1 != nil || article_id < 1 || err2 != nil || comment_id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid article or comment ID")
		return
	}

	comment, err := app.models.Comments.Get(comment_id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, comment)
}

func (app *application) updateArticleCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["commentId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	comment, err := app.models.Comments.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Body *string `json:"body"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if input.Body != nil {
		comment.Body = *input.Body
	}

	err = app.models.Comments.Update(comment)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		fmt.Println(err)
		return
	}
	app.respondWithJSON(w, http.StatusOK, comment)
}

func (app *application) deleteArticleCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["commentId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid comment ID")
		return
	}

	err = app.models.Comments.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
