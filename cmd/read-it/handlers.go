package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string `json:"title"`
		AuthorId  string `json:"author_id"`
		Genre     string `json:"genre"`
		Body      string `json:"text"`
	}
	
}
