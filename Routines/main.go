package main

import (
	"fmt"
	"sync"
	"time"
)

func workerNode(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker node %d started.\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker node %d ended.\n", id)

}

func main() {
	fmt.Println("Program start execution")

	var wg sync.WaitGroup

	wg.Add(1)
	go workerNode(1, &wg)

	wg.Add(1)
	go workerNode(2, &wg)

	wg.Wait()
	fmt.Println("Program terminated")

}
