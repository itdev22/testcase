package main

import "log"

func main() {

}

func no1() {
	ch := make(chan int)
	go func() { ch <- 1 }()
	go func() { ch <- 2 }()

	log.Println(<-ch, <-ch)
}

func no3() {
	var total int = 0
	for {

	}
}
