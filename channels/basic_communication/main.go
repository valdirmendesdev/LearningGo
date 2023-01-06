package main

import (
	"fmt"
	"math/rand"
	"time"
)

func download(url string) {
	r := rand.Intn(4)
	time.Sleep(time.Duration(r) * time.Second)
}

func main() {
	urls := []string{
		"http://files.com/file1.png",
		"http://files.com/file2.png",
		"http://files.com/file3.png",
	}

	finished := make(chan string)

	for _, url := range urls {
		// isso é necessário agora porque estamos usando uma closure
		url := url

		go func() {
			download(url)
			finished <- url
		}()
	}

	for range urls {
		url := <-finished
		fmt.Println("concluído:", url)
	}
}
