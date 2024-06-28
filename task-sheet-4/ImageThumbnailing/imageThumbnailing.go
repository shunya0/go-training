//simply run index.html upload file and then you can find file in the specified dir

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const (
	uploadPath      = "./uploads"
	thumbnailPath   = "./thumbnails"
	thumbnailWidth  = 100
	thumbnailHeight = 100
)

func genThumbnail(filePath, fileName string) error {

	image, err := imaging.Open(filePath)
	if err != nil {
		return fmt.Errorf("Error while opening IMG: ", err)
	}

	thumbnail := imaging.Thumbnail(image, thumbnailWidth, thumbnailHeight, imaging.Lanczos)
	thumbFilePath := filepath.Join(thumbnailPath, fileName)
	errs := imaging.Save(thumbnail, thumbFilePath)
	if errs != nil {
		return fmt.Errorf("Error saving thumbnail: ", err)
	}

	return nil
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing data file size should be smaller than 10Mb: ", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retriving the file: ", http.StatusBadRequest)
		return
	}

	defer file.Close()

	filePath := filepath.Join(uploadPath, handler.Filename)

	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file: ", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	io.Copy(out, file)

	//To Do := Create Thumbnail

	errs := genThumbnail(filePath, handler.Filename)
	if errs != nil {
		http.Error(w, "Error generating thumbnail: ", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File Uploaded successfully")
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)

	filePath := filepath.Join(thumbnailPath, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File Not Found: ", http.StatusNotFound)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/thumbnails/", thumbnailHandler)
	http.HandleFunc("/healthCheck", healthHandler)
	// port := 8080
	fmt.Println("Starting server on: 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
