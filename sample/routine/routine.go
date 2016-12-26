package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("Boring !", c)
	for i := 0; i < 10; i++ {
		fmt.Printf("You say : %q\n", <-c)
	}
	fmt.Println("Main leaving")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(100000000)))
	}
}
