package temporalsort

import (
	"sync"
	"time"
)

// Sort takes in an `[]int`, and returns the sorted version of it using the TemporalSort algorithm.
func Sort(arr []int) []int {
	res := []int{}
	var wg sync.WaitGroup

	for _, v := range arr {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(v) * time.Second)
		}()

		res = append(res, v)
	}

	wg.Wait()
	return res
}
