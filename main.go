package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Scope struct {
	Project string
	Area    string
}

type Note struct {
	Title string
	Tags  []string
	Text  string
	Scope Scope
}

var mdbClient *mongo.Client

func main() {
	const serverAddr string = "127.0.0.1:8081"
	// TODO: Replace with your connection string
	const connStr string = "mongodb+srv://yourusername:yourpassword@notekeeper.xxxxxx.mongodb.net/?retryWrites=true&w=majority&appName=NoteKeeper"

	fmt.Println("Hola Caracola")

	ctxBg := context.Background()
	var err error
	mdbClient, err = mongo.Connect(ctxBg, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mdbClient.Disconnect(ctxBg); err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP Caracola"))
	})
	http.HandleFunc("POST /notes", createNote)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Note: %+v", note)
}
