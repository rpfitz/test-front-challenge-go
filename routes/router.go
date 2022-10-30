package routes

import (
	"frontendmod/controller"
	"frontendmod/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func MyRouter() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(middleware.NotFound)

	router.Handle("/", middleware.ValidateAuthentication(controller.GetLogin)).Methods("GET", http.MethodOptions)
	router.Handle("/login", middleware.ValidateAuthentication(controller.GetLogin)).Methods("GET", http.MethodOptions)
	router.HandleFunc("/login", controller.PostLogin).Methods("POST")

	router.HandleFunc("/google/login", controller.GoogleLogin).Methods("POST")
	router.HandleFunc("/google/callback", controller.GoogleCallback)

	router.Handle("/sign-up", middleware.ValidateAuthentication(controller.GetSignUp)).Methods("GET", http.MethodOptions)
	router.HandleFunc("/sign-up", controller.PostSignUp).Methods("POST")

	router.Handle("/edit-profile", middleware.ValidateAuthentication(controller.GetEditProfile)).Methods("GET", http.MethodOptions)
	router.Handle("/edit-profile", middleware.ValidateAuthentication(controller.PostEditProfile)).Methods("POST", http.MethodOptions)

	router.Handle("/profile", middleware.ValidateAuthentication(controller.Profile)).Methods("GET", http.MethodOptions)

	router.HandleFunc("/logout", controller.Logout).Methods("POST", http.MethodOptions)
	return router
}
