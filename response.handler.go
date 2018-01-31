package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseHandler interface {
	JSON(c *gin.Context, status int, model interface{})
	JSONToken(c *gin.Context, status int, model interface{}, token string)
	Errors(c *gin.Context, status int, errorObjects []*ErrorObject)
	Error(c *gin.Context, title string, status int, detail string)
	ValidationErrors(c *gin.Context, errors error)
	InternalServerError(c *gin.Context)
	GenerateTokenError(c *gin.Context)
	NotFound(c *gin.Context)
	MalformedJSON(c *gin.Context)
	NoRoute(c *gin.Context)
	Unauthorised(c *gin.Context)
	jwtAbort(c *gin.Context, msg string)
	downloadFile(c *gin.Context, filename string)
}

type APIResponseHandler struct{}

func NewResponseHandler() ResponseHandler {
	return &APIResponseHandler{}
}

func (*APIResponseHandler) JSON(c *gin.Context, status int, model interface{}) {
	c.JSON(200, gin.H{
		"code": status,
		"msg":  "success",
		"data": model,
	})
}
func (*APIResponseHandler) JSONToken(c *gin.Context, status int, model interface{}, token string) {
	c.JSON(200, gin.H{
		"code":  status,
		"msg":   "success",
		"token": token,
		"data":  model,
	})
}
func (*APIResponseHandler) Errors(c *gin.Context, status int, errorObjects []*ErrorObject) {
	c.JSON(200, gin.H{
		"code": status,
		"msg":  errorObjects,
	})
}

func (r *APIResponseHandler) Error(c *gin.Context, title string, status int, detail string) {
	r.Errors(c, status, []*ErrorObject{
		&ErrorObject{title, detail},
	})
}

func (r *APIResponseHandler) InternalServerError(c *gin.Context) {
	r.Error(c, InternalServerError, http.StatusInternalServerError, "Something went wrong")
}

func (r *APIResponseHandler) NotFound(c *gin.Context) {
	r.Error(c, NotFound, http.StatusNotFound, "Resource does not exist")
}

func (r *APIResponseHandler) MalformedJSON(c *gin.Context) {
	r.Error(c, MalformedJson, http.StatusBadRequest, "Request contains invalid JSON")
}

func (r *APIResponseHandler) NoRoute(c *gin.Context) {
	r.Error(c, NotFound, http.StatusNotFound, "No route found")
}

func (r *APIResponseHandler) ValidationErrors(c *gin.Context, err error) {
	r.Error(c, ValidationError, http.StatusForbidden, err.Error())
}

func (r *APIResponseHandler) Unauthorised(c *gin.Context) {
	r.Error(c, Unauthorised, http.StatusUnauthorized, "You don't have permission for this resource")
}

func (r *APIResponseHandler) GenerateTokenError(c *gin.Context) {
	r.Error(c, InternalServerError, http.StatusBadRequest, "token 生成失败")
}

func (r *APIResponseHandler) jwtAbort(c *gin.Context, msg string) {
	r.Error(c, Unauthorised, http.StatusUnauthorized, msg)
	c.Abort()
}

func (r *APIResponseHandler) downloadFile(c *gin.Context, filename string) {

	c.Header("Content-Type", "application/vnd.ms-excel;charset=UTF-8")
	c.Header("Pragma", "public")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	c.Header("Content-Type", "application/force-download")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Type", "application/download")
	c.Header("Content-Disposition", "attachment;filename="+time.Now().Format("2006-01-02 15:04:05")+".xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

}
