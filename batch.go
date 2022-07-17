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
	users := make(chan user, n)
	ch := make(chan struct{}, pool)

	for i := int64(0); i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(j int64) {
			defer wg.Done()
			users <- getOne(j)
			<-ch
		}(i)
	}
	wg.Wait()

	close(users)
	for user := range users {
		res = append(res, user)
	}

	return
}
