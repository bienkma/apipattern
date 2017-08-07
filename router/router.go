package router

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"

	"apipattern/handler"
	"apipattern/view"
	"apipattern/log"
	myMiddleware "apipattern/router/middleware"
)

func Register(r *chi.Mux) {

	// Setup log
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	// Add Middleware for router
	r.Use(chiMiddleware.Compress(2, "gzip"))
	r.Use(myMiddleware.CORS)

	r.Use(log.NewStructuredLogger(logger))

	//
	r.Group(func(r chi.Router) {
		r.Get("/api/about", makeHandler(handler.About))
	})
}

func makeHandler(handlerFunc handler.HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := handler.GetBaseHandler(w, r)
		res := handlerFunc(h)
		view.RenderJson(w, res)
	}
}
