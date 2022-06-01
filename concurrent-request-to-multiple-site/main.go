package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	InfoColor    = "\033[1;32m [%d] %s %s \n \033[0m"
	NoticeColor  = "\033[1;36m [%d] %s %s \n \033[0m"
	WarningColor = "\033[1;33m [%d] %s %s \n \033[0m"
	ErrorColor   = "\033[1;31m %s %s  %s \n \033[0m"
	DebugColor   = "\033[0;36m [%d] %s %s \n \033[0m"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func sendRequest(url string) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf(ErrorColor, url, time.Since(start), err.Error())
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	if res.StatusCode == 200 {
		fmt.Printf(InfoColor, res.StatusCode, url, time.Since(start))
		return
	}
	fmt.Printf(WarningColor, res.StatusCode, url, time.Since(start))

}
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <file.txt>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			for _, url := range strings.Split(string(buf[:n]), "\n") {
				go sendRequest("https://" + url)
				wg.Add(1)
			}
		}
	}

	wg.Wait()
}
