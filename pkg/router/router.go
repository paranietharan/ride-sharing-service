package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	//router.HandleFunc("/books", handlers.GetBooks).Methods("GET")

	return router
}

func StartServer() {
	router := InitializeRoutes()
	http.Handle("/", router)

	// create goroutines &
	// Start the server
	server := &http.Server{Addr: ":8080"}

	fmt.Println("Server started on port 8080...")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server stopped:", err)
		}
	}()
}
