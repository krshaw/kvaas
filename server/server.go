package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/krshaw/kvaas/database"
)

// the getHandler should be using a database package
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got get request")
	key := r.URL.Query().Get("key")
	// open file for searching
	// search backwards to get most recent entry of the key
	fmt.Printf("Getting the value for %s\n", key)
	v, err := database.Get(key)
	if err != nil {
		fmt.Fprintf(w, "Error getting key: %s", err)
	}
	w.Write(v)
}

// the createHandler should be using a database package
func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got CREATE request")
	// body should be json of the form
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "ERROR READING REQUEST BODY")
		return
	}
	err = database.Create(body)
	if err != nil {
		fmt.Fprintf(w, "ERROR WRITING PAIR\n%s", err)
		return
	}
	fmt.Fprintf(w, "Writing the k/v pair:\n%s\n", string(body))
}

func Start() {
	database.StringIndex = make(map[string]int)
	database.IntIndex = make(map[int64]int)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/create", createHandler)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
