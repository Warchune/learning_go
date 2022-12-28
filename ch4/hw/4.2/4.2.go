package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "хеш длиной 256, 384 или 512")

func main() {
	var answ []byte
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	switch *width {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(b))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(b))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(b))
	default:
		log.Fatal("Неверная длина хеша")
	}
	fmt.Printf("%x\n", answ)
}
