package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// go run main.go
// curl --location 'localhost:3000'
func main() {

	http.HandleFunc("/blocking", func(w http.ResponseWriter, r *http.Request) {
		// counter := heavyTask()

		counter := concurrentTasks()
		w.Write([]byte(fmt.Sprintf("Counter: %v", counter)))
	})

	http.HandleFunc("/heavy-task", func(w http.ResponseWriter, r *http.Request) {
		counter := heavyTask()
		w.Write([]byte(fmt.Sprintf("Counter: %v", counter)))
	})

	addr := ":3000"

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func heavyTask() int {
	counter := 0
	for i := 0; i < 10_000_000_000; i++ {
		counter++
	}
	return counter
}

func goroutineTask(workers int) int {
	counter := 0
	for i := 0; i < 10_000_000_000/workers; i++ {
		counter++
	}
	return counter
}

func concurrentTasks() int {
	concurrentTasks := 1000
	wg := &sync.WaitGroup{}
	wg.Add(concurrentTasks)
	for i := 0; i < concurrentTasks; i++ {
		go func(workerId int, wg *sync.WaitGroup) {
			counter := goroutineTask(concurrentTasks)
			fmt.Printf("Worker ID: %v - %v", workerId, counter)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
	return 10_000_000_000
}
