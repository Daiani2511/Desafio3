package main

import(
	"fmt"
	"time"
)

func ping(ch chan string){
	for {
		ch <- "ping"
		time.Sleep(1 * time.Second)
	}
}

func pong(ch chan string) {
	for {
		ch <- "pong"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan string)

	go ping(ch)
	go pong(ch)

	for {
		msg := <- ch
		fmt.Println(msg)
	}
}