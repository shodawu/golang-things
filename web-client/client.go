package webclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

// RunClient ...
func RunClient() {

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(iProcess int) {
			fmt.Println("ProcessID: ", iProcess)
			defer wg.Done()
			for {
				r, err := http.Post("http://localhost:1234/ping", "application/json", nil)

				if err != nil {
					fmt.Println(err)
					time.Sleep(5 * time.Second)
					continue
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				fmt.Println(string(body))
			}
		}(i)
	}
	wg.Wait()

}

func RunClientP2() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iProcess int) {
			fmt.Println("ProcessID: ", iProcess)
			defer wg.Done()
			cFileName := make(chan string)
			cUUID := make(chan string)

			go func(c chan string) {
				var err = errors.New("not conneted")
				var r *http.Response
				for err != nil {
					r, err = http.Post("http://localhost:1234/ping", "application/json", nil)
					if err == nil {
						break
					}
					fmt.Println(err)
					time.Sleep(5 * time.Second)
				}
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				fileName := fmt.Sprintf("%v.log", string(body))
				fmt.Println(iProcess, "API returns", fileName)

				go p2UUIDChan(c, fileName)
			}(cFileName)

			go func() {
				s := <-cUUID
				fmt.Println(iProcess, "Find log", s)
			}()

			select {
			case fName := <-cFileName:
				var erFile = errors.New("not exists")
				var f *os.File
				for erFile != nil {
					f, erFile = os.Open(fName)
					if erFile != nil {
						time.Sleep(100 * time.Millisecond)
					}
				}
				b, _ := ioutil.ReadAll(f)
				go p2UUIDChan(cUUID, string(b))
			}

		}(i)
	}
	wg.Wait()

}

func p2UUIDChan(c chan string, data string) {
	c <- data
}
