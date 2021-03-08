package webclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
