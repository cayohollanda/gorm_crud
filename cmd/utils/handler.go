package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(w *gin.Context, status int, payload interface{}) {
	var resp Response
	resp.Status = status
	resp.Data = payload

	switch status {
	case 404:
		resp.Message = "Resource not found"
	case 200:
		resp.Message = "Successful request"
	}

	w.JSON(status, resp)
}
