package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	for _, address := range os.Args[1:] {
		//a := address
		go func(address string) {
			conn, err := net.Dial("tcp", address)
			//conn, err := net.Dial("tcp", a)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		}(address)
	}
	time.Sleep(10 * time.Second)
	fmt.Println()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
