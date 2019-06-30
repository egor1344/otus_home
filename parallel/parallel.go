package parallel

import (
	"fmt"
	"sync"
)

// CustomFunc обертка над обычной функцией
type CustomFunc func() error

// Parallel функция для выполнения задач
// number количество одновременных выполнений задач
func Parallel(number int, numError int, workers []CustomFunc) error {
	var wg sync.WaitGroup
	wg.Add(len(workers))

	for _, worker := range workers {
		go func () {
			err := worker()
			if err != nil {fmt.Println("error func", worker)}
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}
