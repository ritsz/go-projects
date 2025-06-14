package handlers

import (
	"file-upload-go/services"
	"fmt"
	"log"
	"net/http"
)

/**
 * The interface that supports uploading a file using Mutlipart form data.
 */
type UploadHandler struct {
	service services.IUploadService
}

func (uh UploadHandler) FileUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("Reached upload.")

	//1. Param input for multipart file upload
	r.ParseMultipartForm(1 << 20)

	/**
	 * 2. Retrieve file from form-data
	 * <file> is the form key that we will read from.
	 * Client should use the same form key when uploading the file
	 */
	file, handler, err := r.FormFile("file")
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file %s\n", err)
		fmt.Fprintf(w, errStr)
		log.Println(errStr)
		return
	}

	/**
	 * 3. The upload service saves the file.
	 */
	result, err := uh.service.SaveFile(file, handler)
	fmt.Fprintf(w, result)

	if err != nil {
		// Error handling here
		log.Println("Error when saving file:", err)
		return
	}
}

func NewUploadHandler(service services.IUploadService) UploadHandler {
	return UploadHandler{
		service: service,
	}
}

/**
 * Create an upload handler that saves to disc.
 */
func NewUploadHandlerSaveToDisc() UploadHandler {
	return UploadHandler{
		service: services.LocalUploadService{},
	}
}
