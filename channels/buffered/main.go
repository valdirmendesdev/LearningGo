// source: https://dev.to/tutorial_livre/golang-desmistificando-channels-buffered-channels-16d0
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	finished := make(chan bool, 10)
	for {
		finished <- true

		fmt.Println(time.Now())
		go func() {
			client := http.Client{}

			resp, _ := client.Get("https://www.google.com")
			resp.Body.Close()

			fmt.Println(resp.Status)
			<-finished
		}()
	}
}
