package manythreads

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func expensiveCalculation(num int) float64 {
	result := float64(num)
	for i := 0; i < 1000; i++ {
		result = math.Sin(result) * math.Cos(result)
	}
	return result
}

func multiplyByTwo(num int, resultChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := expensiveCalculation(num)
	resultChan <- result
}

func Run() {
	numRoutines := 100000

	// Goroutine version
	startGoroutine := time.Now()
	resultChan := make(chan float64, numRoutines) // Buffered channel
	var wg sync.WaitGroup

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go multiplyByTwo(i, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan) // Close the channel after all routines are done
	}()

	results := make([]float64, 0, numRoutines)
	for result := range resultChan {
		results = append(results, result)
	}

	elapsedGoroutine := time.Since(startGoroutine)
	fmt.Println("Goroutine version completed in", elapsedGoroutine)

	// Synchronized version
	startSync := time.Now()
	resultsSync := make([]float64, numRoutines)
	for i := 0; i < numRoutines; i++ {
		resultsSync[i] = expensiveCalculation(i)
	}
	elapsedSync := time.Since(startSync)
	fmt.Println("Synchronized version completed in", elapsedSync)
}
