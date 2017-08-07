package handler

import (
	"apipattern/view"
	"net/http"
)

type BaseHandler struct {
	w http.ResponseWriter
	r *http.Request
}

type HandleFunc func(ctx BaseHandler) view.ApiResponse

func GetBaseHandler(w http.ResponseWriter, r *http.Request) BaseHandler {
	return BaseHandler{w, r}
}
