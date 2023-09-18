package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func addTask(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	task := uuid.New()
	fmt.Println("Adding Task : ", task.String())
	ch <- task.String()
}

func recieveTask(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task Recieved : ", <-ch)
}

func main() {
	var ch chan string     //declaring channel
	ch = make(chan string) // defining channel

	var wg sync.WaitGroup

	defer close(ch) // closing the channel

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go addTask(ch, &wg)
		go recieveTask(ch, &wg)
	}

	wg.Wait()
	//time.Sleep(time.Second * 2)
	fmt.Println("All tasks Iterated")
}
