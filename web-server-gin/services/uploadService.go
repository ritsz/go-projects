package services

import (
	"file-upload-go/common"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

/**
 * Interface that takes the multipart file and, using the File handle,
 * saves the file to storage.
 */
type IUploadService interface {
	SaveFile(multipart.File, *multipart.FileHeader) (string, error)
}

/**
 * An implementation of IUploadService that saves the file to local storage.
 */
type LocalUploadService struct {
}

func (lus LocalUploadService) SaveFile(file multipart.File, handler *multipart.FileHeader) (string, error) {
	/**
	 * 1. Retrieve file from form-data
	 * <file> is the form key that we will read from.
	 * Client should use the same form key when uploading the file
	 */
	defer file.Close()

	log.Println("File name:", handler.Filename)
	log.Println("File size:", handler.Size)
	for header, value := range handler.Header {
		log.Println("File header::", header, ":", value)
	}

	/**
	 * 2. Create a temporary file to our directory
	 */
	tempFolderPath := fmt.Sprintf("%s%s", common.RootPath, "/tempFiles")
	tempFileName := fmt.Sprintf("upload-%s-*.%s", common.FileNameWithoutExtension(handler.Filename), filepath.Ext(handler.Filename))

	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
	if err != nil {
		errStr := fmt.Sprintf("Error in creating the file %s\n", err)
		fmt.Println(errStr)
		return errStr, err
	}

	defer tempFile.Close()

	/**
	 * 3. Write upload file bytes to your new file
	 */
	buffer := make([]byte, 4096)

	for {
		bytesread, err := file.Read(buffer)
		buffer = buffer[:bytesread]

		if bytesread != 4096 {
			log.Println("Read bytes:", bytesread)
		}

		if err != nil {
			if err != io.EOF {
				log.Println("Error while reading file:", err)
				return "Error while reading file", err
			}

			log.Println("Finished writing to temp file: ", tempFile.Name())
			break
		}

		tempFile.Write(buffer)
	}

	return "Successfully uploaded\n", nil

}

func NewLocalUploadService() LocalUploadService {
	return LocalUploadService{}
}
