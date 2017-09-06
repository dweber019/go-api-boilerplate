package app

import (
	"github.com/gorilla/mux"
	"github.com/dweber019/go-api-boilerplate/app/ctrl"
)

// NewRouter ...
func NewRouter() *mux.Router {

	//Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	/**
	 * meta-data
	 */
	mainRouter.Methods("GET").Path("/api/info").HandlerFunc(ctrl.GetAPIInfo)

	/**
	 * /users
	 */
	// usersRouter.HandleFunc("/", l.Use(c.GetAllUsersHandler, m.SaySomething())).Methods("GET")
	apiRouter.Methods("GET").Path("/users").HandlerFunc(ctrl.GetAllUsersHandler)
	apiRouter.Methods("POST").Path("/users").HandlerFunc(ctrl.CreateUserHandler)
	apiRouter.Methods("GET").Path("/users/{id}").HandlerFunc(ctrl.GetUserByIdHandler)
	apiRouter.Methods("PUT").Path("/users/{id}").HandlerFunc(ctrl.UpdateUserHandler)
	apiRouter.Methods("DELETE").Path("/users/{id}").HandlerFunc(ctrl.DeleteUserHandler)

	/**
	 * /items
	 */
	apiRouter.Methods("GET").Path("/items").HandlerFunc(ctrl.GetAllItemsHandler)

	return mainRouter
}
