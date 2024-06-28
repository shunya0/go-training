package main

import (
	"bufio"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type session struct {
	id     string
	userID int
	expiry int64
}

var sessions = map[string]session{}

const (
	saveFile       = "./save"
	sessionDuratio = 1200
)

func getUniqueID() string {
	id := make([]byte, 32)
	rand.Read(id)
	return base64.URLEncoding.EncodeToString(id)
}
func createSession(userID int) (string, error) {

	sessionID := getUniqueID()
	expiry := time.Now().Add(sessionDuratio * time.Second).Unix()

	sessions[sessionID] = session{
		id: sessionID, userID: userID, expiry: expiry,
	}
	return sessionID, nil

}
func getSession(sessionID string) (session, bool) {
	sess, exists := sessions[sessionID]
	if !exists || sess.expiry < time.Now().Unix() {
		return session{}, false
	}
	return sess, true
}
func invalidSession(sessionID string) {
	delete(sessions, sessionID)
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	// id := r.URL.Query().Get("id")

	filePath := filepath.Join(saveFile, "file.txt")

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	dataByte := []byte("\n" + text)
	_, errS := writer.Write(dataByte)
	if errS != nil {
		http.Error(w, "File not saved", http.StatusInternalServerError)
		return
	}
	if errs := writer.Flush(); errs != nil {
		http.Error(w, "Error Flushing file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File Created"))

}

func readHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")

	// if id != FILE_STRUCT.id {
	// 	http.Error(w, "Incorrect id", http.StatusUnauthorized)
	// 	return
	// }
	filename := "file.txt"
	filepath := filepath.Join(saveFile, filename)
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, "Error opening file to read data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w.Write([]byte(scanner.Text() + "\n"))
	}

}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	// if id != FILE_STRUCT.id {
	// 	http.Error(w, "WRONG ID", http.StatusUnauthorized)
	// 	return
	// }
	filename := "file.txt"
	filepath := filepath.Join(saveFile, filename)

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, "Error can't reset", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		http.Error(w, "Unable to reset data", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Reseting file: DONE"))

}

func downloadHander(w http.ResponseWriter, r *http.Request) {
	filename := "file.txt"
	filepath := filepath.Join(saveFile, filename)

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, "Error can't reset", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=\"file.txt\"")
	w.Header().Set("Content-Type", "text/Plain")
	// w.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size()))

	_, ERR := io.Copy(w, file)
	if ERR != nil {
		http.Error(w, "Unable to download file", http.StatusInternalServerError)
		return
	}

}

func sessionMiddleaware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		sessionID := cookie.Value
		sess, valid := getSession(sessionID)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		r.Header.Set("user_id", strconv.Itoa(sess.userID))
		next.ServeHTTP(w, r)

	})
}
func deCompress(inputFile string, outPutfileName string) error {
	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return err
	}
	defer readFile.Close()

	outputeFile, err := os.Create(outPutfileName)
	if err != nil {
		fmt.Println("Error while creting output file: ", err)
		return err
	}
	defer outputeFile.Close()

	gzipReader, err := gzip.NewReader(readFile)
	if err != nil {
		fmt.Println("Error while reading the output file: ", err)
		return err
	}
	defer gzipReader.Close()

	if _, err := io.Copy(outputeFile, gzipReader); err != nil {
		fmt.Println("Error , while copying file: ", err)
		return err
	}
	return nil
}

func compress(inputFile string, outPutfileName string) error {

	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return err
	}
	defer readFile.Close()

	outputFile, err := os.Create(outPutfileName)
	if err != nil {
		fmt.Println("Error while creating output file: ", err)
		return err
	}
	defer outputFile.Close()

	buf := make([]byte, 4096)

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	for {
		data, err := readFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error while rading buffer file: ", err)
			return err
		}
		if data == 0 {
			break
		}
		if _, err := gzipWriter.Write(buf[:data]); err != nil {
			fmt.Println("Error while copying file: ", err)
			return err
		}
		if err := gzipWriter.Flush(); err != nil {
			fmt.Println("Error while flushing: ", err)
			return err
		}
	}

	return nil

}

func loginHander(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Error id should be int", http.StatusBadRequest)
		return
	}
	sessionid, err := createSession(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionid,
		Expires: time.Now().Add(sessionDuratio * time.Second),
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged in successfully")
}

func logoutHander(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	invalidSession(cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged out successfully")
}
func compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		task := r.URL.Query().Get("task")
		if task == "compress" {
			inputFile := filepath.Join(saveFile, "file.txt")
			outputFile := filepath.Join(saveFile, "file.txt.gz")
			if err := compress(inputFile, outputFile); err != nil {
				http.Error(w, "Failed to compress file", http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, outputFile)
			return
		} else {
			inputFile := filepath.Join(saveFile, "file.txt.gz")
			outputFile := filepath.Join(saveFile, "file.txt")
			if err := deCompress(inputFile, outputFile); err != nil {
				http.Error(w, "Failed to decompress file", http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, outputFile)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	http.HandleFunc("/login", loginHander)
	http.HandleFunc("/logout", logoutHander)

	http.Handle("/add", sessionMiddleaware(http.HandlerFunc(addHandler)))
	http.Handle("/read", sessionMiddleaware(http.HandlerFunc(readHandler)))
	http.Handle("/reset", sessionMiddleaware(http.HandlerFunc(resetHandler)))
	http.Handle("/download", compressionMiddleware(http.HandlerFunc(downloadHander)))

	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// /add , /read , /reset , /download
// /add: text as input and save to a txt file
// /read: return all content of txt
// /reset: del content
// download: download all the contents as a txt file
