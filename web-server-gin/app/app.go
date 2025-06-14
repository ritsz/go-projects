package app

import (
	"file-upload-go/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 * This function starts the application, and performs a blocking listen on the
 * specified port.
 */
func StartApp() {
	// 1. Create the router
	router := mux.NewRouter()

	handler := handlers.NewUploadHandlerSaveToDisc()

	// 2. Add functions to router
	router.HandleFunc("/upload", handler.FileUpload).Methods(http.MethodPost)

	// 3. Start the server
	fmt.Println("Server started")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalln("Unable to start server:", err)
	}
}
