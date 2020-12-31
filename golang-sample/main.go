package main

import (
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("[init] hello")
}

func main() {
	log.Println("[main] hello")
}
