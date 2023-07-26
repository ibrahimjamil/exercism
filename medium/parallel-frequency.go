package medium

import (
	"sync"
)

func ParallelFrequencyCount(texts []string) map[rune]int {
	// using mutex lock as frequencyMap is common between goroutines
	// so having conflicts with write operation
	frequencyMap := make(map[rune]int)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for _, text := range texts {
		wg.Add(1)

		// goroutine ifee function
		go func(text string) {
			defer wg.Done()
			for _, char := range text {
				mutex.Lock()
				frequencyMap[char]++
				mutex.Unlock()
			}
		}(text)
	}

	wg.Wait()
	return frequencyMap
}
