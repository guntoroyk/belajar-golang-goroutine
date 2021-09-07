package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Guntoro Yudhy Kusuma"
		fmt.Println("selesai mengirim data ke channel")
	}()

	var data = <-channel
	fmt.Println("data", data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	//time.Sleep(2 * time.Second)
	channel <- "Guntoro Yudhy Kusuma"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	fmt.Println("satu")

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println("data", data)
	fmt.Println("dua")

}

// OnlyIn function with parameter send-only channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Guntoro Yudhy Kusuma"

	// cannot use channel to receive data
	//data := <- channel
}

// OnlyOut function with parameter receive-only channel
func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)

	// cannot use channel to send data
	//channel <- "Guntoro Yudhy Kusuma"

	// use channel to receive data
	data := <- channel
	fmt.Println("data", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(4 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	channel <- "guntoro"
	channel <- "yudhy"
	channel <- "kusuma"

	// panic deadlock
	//channel <- "aaaa"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println("finish")
}