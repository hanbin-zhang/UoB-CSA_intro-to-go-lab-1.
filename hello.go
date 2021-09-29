package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go superHelloWorld(i)

	}
	time.Sleep(1 * time.Second)
	fmt.Println("a")
}

func superHelloWorld(int2 int) {

	fmt.Println(byte(255))
}
