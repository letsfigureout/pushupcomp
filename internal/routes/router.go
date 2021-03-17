package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIRoute struct {
	Log *log.Logger
}

func API(log *log.Logger) http.Handler {

	api := APIRoute{
		Log: log,
	}

	router := mux.NewRouter()
	router.HandleFunc("/", api.IndexRoute)

	return router
}

func (api *APIRoute) IndexRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Index")
}
