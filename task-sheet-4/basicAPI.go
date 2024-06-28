package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type session struct {
	id     string
	userID int
	expiry int64
}

var sessions = map[string]session{}

const sessionDuration = 1200

func createSession(userID int) (string, error) {

	sessionID := genUniqueID()
	expiry := time.Now().Add(sessionDuration * time.Second).Unix()

	sessions[sessionID] = session{
		id: sessionID, userID: userID, expiry: expiry,
	}
	return sessionID, nil
}

func getSession(sessionID string) (session, bool) {

	sess, exists := sessions[sessionID] //session , exists
	if !exists || sess.expiry < time.Now().Unix() {
		return session{}, false
	}
	return sess, true

}
func invalidSession(sessionID string) {
	delete(sessions, sessionID)
}

func genUniqueID() string {
	// var id []byte
	id := make([]byte, 32)
	rand.Read(id)
	return base64.URLEncoding.EncodeToString(id)
}

// Helper Functions  ^^^

func loginHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Bad request: ID should be int", http.StatusBadRequest)
		return
	}
	sessionid, err := createSession(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error please try again ")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionid,
		Expires: time.Now().Add(sessionDuration * time.Second),
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged in successfully")
}

func logOutHandler(w http.ResponseWriter, r *http.Request) {
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

// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token != "Bearer Secret" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Unauthorized"))
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

type toDo struct {
	task        string
	id          string
	description string
}

type toDolist struct {
	listToDo []toDo
}

var listOfTask toDolist

func (t *toDolist) add(name string, id string, description string) {
	toDoTask := toDo{task: name, id: id, description: description}
	t.listToDo = append(t.listToDo, toDoTask)
}

func (t *toDolist) read(id string) string {
	for _, v := range t.listToDo {
		if v.id == id {
			return "Task created successfully"
		}
	}
	return "Task not created"
}

func (t *toDolist) del(id string) {
	for i, v := range t.listToDo {
		if v.id == id {
			arr1 := t.listToDo[:i]
			arr2 := t.listToDo[i+1:]
			t.listToDo = append(arr1, arr2...)
			return
		}
	}
	fmt.Println("Task not found! ", id)
}
func handleDelTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	listOfTask.del(id)
	fmt.Fprintf(w, "Task with id %s deleted", id)
}

func (t *toDolist) edit(task string, id string, description string) {

	for i, v := range t.listToDo {
		if v.id == id {
			t.listToDo[i].task = task
			t.listToDo[i].description = description
			return
		}

	}
	fmt.Println("Task not found: ", id)
}
func handleEditTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task := r.URL.Query().Get("task")
	description := r.URL.Query().Get("description")

	listOfTask.edit(task, id, description)
	fmt.Fprintln(w, "Task changed successfully!")
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	task := r.URL.Query().Get("task")
	description := r.URL.Query().Get("description")

	listOfTask.add(task, id, description)
	fmt.Fprintln(w, "Task created successfully!")
}
func handleReadTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	listtask := listOfTask.read(id)
	fmt.Fprintln(w, listtask)
}
func handleReadAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, listOfTask.listToDo)
}
func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logOutHandler)
	http.Handle("/tasks/create", sessionMiddleaware(http.HandlerFunc(handleCreateTask)))
	http.Handle("/tasks/read", sessionMiddleaware((http.HandlerFunc(handleReadTask))))
	http.Handle("/tasks/edit", sessionMiddleaware(http.HandlerFunc(handleEditTask)))
	http.Handle("/tasks/readall", sessionMiddleaware(http.HandlerFunc(handleReadAll)))
	http.Handle("/tasks/del", sessionMiddleaware(http.HandlerFunc(handleDelTask)))

	port := ":8080"

	fmt.Println("Starting server at", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
