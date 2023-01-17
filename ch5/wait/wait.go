package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "использование: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.SetPrefix("wait: ")
		log.SetFlags(0)
		log.Fatalf("Сервер не работает: %v\n", err)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("Сервер не отвечает (%s); повтор...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("cервер %s не отвечает; вермя %s", url, timeout)
}
