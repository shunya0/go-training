package main

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static/index.html")

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to home!")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the main page!")
}
func genshinHandler(w http.ResponseWriter, r *http.Request) {
	mimeType := mime.TypeByExtension(filepath.Ext("static/Genshin.png"))
	fmt.Println(mimeType)

	http.ServeFile(w, r, "static/Genshin.png")
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/Betty_Boop.mp4")
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/index", helloHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/genshin", genshinHandler)
	http.HandleFunc("/video", videoHandler)

	portAdd := ":8080"

	fmt.Println("Starting server at: ", portAdd)
	if err := http.ListenAndServe(portAdd, nil); err != nil {
		fmt.Println("Error in starting server: ", err)
		return
	}
}
