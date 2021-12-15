package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

func Poller(res *Resources) {
	for {
		// get the least recently-polled resource and make it as being polled
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}

		if r != nil {
			r.polling = true
		}

		res.lock.Unlock()
		if r == nil {
			continue
		}
		//poll the url

		//update the Resource's polling and lastPolled
		res.lock.Lock()
		r.polling = false
		r.lastPolled = time.Now().Unix()
		res.lock.Unlock()
	}
}

//channle
//type Resource string
//
//func Poller(in, out chan *Resource) {
//	for r := range in {
//		//poll the url
//		// send the porcessed Resource to out
//		out <- r
//	}
//}
