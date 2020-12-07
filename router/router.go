package router

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/r/list/*filepath", http.Dir(svr.dataDir))
	router.GET("/r/status", status)
	router.POST("/r/upload/*filepath", upload) // support http gzip compressed
	router.GET("/r/download/*filepath", download)
	router.GET("/r/info/*filepath", info)
	router.GET("/r/clean/", clean)
	router.GET("/r/backup", backup)
	return router
}