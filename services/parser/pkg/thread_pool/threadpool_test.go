package thread_pool

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestThreadPool_AddTask(t *testing.T) {
	pool := New(1024)
	var executedTasks atomic.Int32

	pool.RunWorkers(12)

	const iter = 100000
	wg := sync.WaitGroup{}
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		go func() {
			defer wg.Done()
			pool.AddTask(func() error {
				executedTasks.Add(1)
				return nil
			})
		}()
	}
	wg.Wait()

	pool.TerminateWait()

	if executedTasks.Load() != iter {
		t.Errorf("Expected 100000 tasks to be executed, got %d", executedTasks.Load())
	}
}
