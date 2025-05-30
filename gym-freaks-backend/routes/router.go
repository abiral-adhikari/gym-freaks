package router

import (
	"gym-freaks-backend/handlers"
	"gym-freaks-backend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	//  Use Route Group to group protected paths
	authrouter := &RouteGroup{
		router: router,
		prefix: "",
	}

	// Food Handler
	authrouter.HandleFunc("/food", handlers.FoodHandlers.Create).Methods("POST")
	authrouter.HandleFunc("/food", handlers.FoodHandlers.Search).Methods("GET")
	authrouter.HandleFunc("/food/{id}", handlers.FoodHandlers.Update).Methods("PATCH")
	authrouter.HandleFunc("/food/{id}", handlers.FoodHandlers.GetOne).Methods("GET")
	authrouter.HandleFunc("/food/{id}", handlers.FoodHandlers.Delete).Methods("DELETE")
	return router

	
	// router.Handle4Func("/exercise", ExerciseHandler).Methods("GET")

}

// RouteGroup struct with middleware applied to all routes it registers
type RouteGroup struct {
	router *mux.Router
	prefix string
}

func (g *RouteGroup) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	return g.router.HandleFunc(g.prefix+path, func(w http.ResponseWriter, r *http.Request) {
		middleware.AuthMiddleware(http.HandlerFunc(handler)).ServeHTTP(w, r)
	})
}
