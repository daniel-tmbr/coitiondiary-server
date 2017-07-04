package routes

import (
	"net/http"

	"coitiondiary.tmbr.me/utils"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	var routes Routes
	routes = append(routes, authRoutes...)
	routes = append(routes, coitionsRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, locationsRoutes...)
	routes = append(routes, tagsRoutes...)

	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = utils.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
