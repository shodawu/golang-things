package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

var mu sync.Mutex

// RunServer ...
func RunServer() {

	s := &http.Server{
		Addr:           ":1234",
		Handler:        http.HandlerFunc(handlePing),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getUUID())
}

func getUUID() string {
	mu.Lock()
	time.Sleep(500 * time.Millisecond)
	u, _ := uuid.NewUUID()
	mu.Unlock()
	return u.String()

}
