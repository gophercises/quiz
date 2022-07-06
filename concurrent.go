package main

import (
	"fmt"
	"time"
)


func con(){
	fmt.Println("Conncurrent example")

	go concurrent(5)
	go concurrent(5)

	fmt.Scanln()
}


func concurrent (n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}