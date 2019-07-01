package parallel

import (
	"errors"
	"sync"
)

// CustomFunc обертка над обычной функцией
type CustomFunc func() error

// Parallel функция для выполнения задач
// number количество одновременных выполнений задач
func Parallel(number int, numError int, workers []CustomFunc) error {
	worker := make(chan func() error)
	errorCount := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(1)
	for i := 0; i < number; i++ {
		go func() error {
			if errorCount <= numError {
				for {
					select {
					case task := <-worker:
						if task() != nil {
							m.Lock()
							errorCount++
							m.Unlock()
						}
						wg.Done()
					}
				}
			} else {
				return errors.New("max count error")
			}
		}()
	}

	for i := 0; i < len(workers); {
		if errorCount <= numError {
			worker <- workers[i]
			wg.Add(1)
			i++
		} else {
			return errors.New("max count error")
		}

	}
	return nil
}
