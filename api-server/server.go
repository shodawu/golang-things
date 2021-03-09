package apiserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

var mu sync.Mutex
var cUUID chan string

// RunServer ...
func RunServer() {

	cUUID = make(chan string)
	go storeUUID2(cUUID)

	s := &http.Server{
		Addr:           ":1234",
		Handler:        http.HandlerFunc(handlePing2),
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

func handlePing2(w http.ResponseWriter, r *http.Request) {
	u, _ := uuid.NewUUID()
	go func(c chan string, data string) {
		c <- data
	}(cUUID, u.String())

	fmt.Fprintf(w, u.String())
}

func storeUUID2(c chan string) {
	for {
		s := <-c
		time.Sleep(2000 * time.Millisecond)
		ioutil.WriteFile(fmt.Sprintf("%v.log", s), []byte(s), 0775)
	}
}
