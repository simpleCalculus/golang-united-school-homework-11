package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	if pool <= 0 {
		return
	}

	if n <= 0 {
		return
	}

	var wg sync.WaitGroup
	var mx sync.Mutex

	for i := int64(0); i < pool; i++ {
		wg.Add(1)
		go func(j int64) {
			for {
				mx.Lock()
				size := len(res)
				if size == int(n) {
					mx.Unlock()
					break
				}
				res = append(res, getOne(int64(size)))
				mx.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return
}
