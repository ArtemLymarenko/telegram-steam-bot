package thread_pool

import (
	"log"
	"sync"
)

type TaskFunc func() error

type ThreadPool struct {
	pool       chan TaskFunc
	terminated bool
	mx         sync.RWMutex
}

func New(bufferSize int) *ThreadPool {
	return &ThreadPool{
		pool:       make(chan TaskFunc, bufferSize),
		terminated: false,
		mx:         sync.RWMutex{},
	}
}

func (t *ThreadPool) RunWorkers(totalWorkers int) {
	for i := 0; i < totalWorkers; i++ {
		go func() {
			for task := range t.pool {
				err := task()
				if err != nil {
					log.Println("Error while running task: ", err)
				}
			}
		}()
	}
}

func (t *ThreadPool) AddTask(task TaskFunc) {
	t.mx.RLock()
	defer t.mx.RUnlock()
	if t.terminated {
		log.Println("Can't add task to terminated pool: terminated")
		return
	}

	t.pool <- task
}

func (t *ThreadPool) TerminateWait() {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.terminated = true
	close(t.pool)
}
