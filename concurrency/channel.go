
package main

import (
	"fmt"
)


func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2  
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    
    // Start 3 workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }
    
    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 5; r++ {
        <-results
    }
}