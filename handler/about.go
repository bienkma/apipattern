package handler

import (
	"net/http"
	"apipattern/view"
)

// Intro About me. So it's not necessary BaseHandler
func About(_ BaseHandler) view.ApiResponse{
	return view.ApiResponse{Code:http.StatusOK, Data:"Welcome to system!"}
}