package view

import "net/http"

type ApiResponse struct {
	Headers map[string]string `json:"-"`
	Code    int
	Data    interface{} `json:"data,omitempty"`
}

func Ok(data interface{}) ApiResponse {
	return ApiResponse{Code: http.StatusOK, Data: data}
}
