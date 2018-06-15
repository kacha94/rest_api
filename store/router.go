package store


import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"quadlets-server-api/store/controllers"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
	Route {
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route {
		"Index",
		"GET",
		"/",
		controller.Index,
	},

	Route {
		"GetJob",
		"GET",
		"/jobs/{id}",
		controller.GetProduct,
	},
	Route {
		"GetJobsInRange",
		"GET",
		"/Search/{query}",
		controller.SearchProduct,
	},
	Route {
		"AddJob",
		"POST",
		"/AddJob",
		controllers.AuthenticationMiddleware(controller.AddProduct),
	},
	Route {
		"UpdateJob",
		"PUT",
		"/UpdateJob",
		controllers.AuthenticationMiddleware(controller.UpdateProduct),
	},
	Route{
		"DeleteJob",
		"DELETE",
		"/deleteProduct/{id}",
		controllers.AuthenticationMiddleware(controller.DeleteProduct),
	}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}