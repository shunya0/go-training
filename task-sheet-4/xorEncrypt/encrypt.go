// open index.html to access encrpter via webPage
//To encrypt you will need to provide file and key
//key must be an int b/w 0 to 255

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	uploadPath         = "./uploads"
	encryptDecryptPath = "./endcrypt"
)

func encryptDecrypt(arr []byte, key byte) ([]byte, error) {
	// var xorKey byte
	// xorKey := 0xAA
	xorKey := key

	for i, v := range arr {
		arr[i] = v ^ xorKey
	}
	return arr, nil

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File size should <less than 10Mb: ", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retriving the file ", http.StatusBadRequest)
		return
	}

	defer file.Close()

	filePath := filepath.Join(uploadPath, handler.Filename)

	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	io.Copy(out, file)
	out.Close()
	// Encrypt file
	outputPath, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, "Error output file not created", http.StatusBadRequest)
		return
	}
	defer outputPath.Close()

	reader := bufio.NewReader(outputPath)
	fileByte, err := io.ReadAll(reader)
	// fmt.Println(len(fileByte))
	if err != nil {
		http.Error(w, "Error reading file: ", http.StatusInternalServerError)
	}

	key := r.FormValue("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}
	intKey, err := strconv.Atoi(key)
	if err != nil {
		http.Error(w, "Invalid key key range should be 0 to 255", http.StatusBadRequest)
		return
	}
	keyXor := byte(intKey)
	encryptedData, err := encryptDecrypt(fileByte, keyXor)
	if err != nil {
		http.Error(w, "File not encrypted: ", http.StatusInternalServerError)
		return
	}
	if len(encryptedData) == 0 {
		http.Error(w, "Data not encrypted: ", http.StatusInternalServerError)
		return
	}
	encryptedFilePath := filepath.Join(encryptDecryptPath, handler.Filename)
	encryptedFile, errss := os.Create(encryptedFilePath)
	if errss != nil {
		http.Error(w, "File not created", http.StatusInternalServerError)
		return
	}

	// outputFile, err := os.Open(encryptedFilePath)

	defer encryptedFile.Close()
	writer := bufio.NewWriter(encryptedFile)

	if _, err := writer.Write(encryptedData); err != nil {
		fmt.Println("File not saved: ", err)
		return
	}
	http.ServeFile(w, r, encryptedFilePath)
	// fmt.Println(encryptedFilePath)
	fmt.Fprintln(w, "File encrypted seccuessfully!")

}

func encryptedDecrypted(w http.ResponseWriter, r *http.Request) {

	fileName := filepath.Base(r.URL.Path)
	filePath := filepath.Join(encryptDecryptPath, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found: ", http.StatusBadRequest)
		return
	}
	defer file.Close()

	http.ServeFile(w, r, filePath)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/xor/key", encryptedDecrypted)

	fmt.Println("Starting server at: 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
