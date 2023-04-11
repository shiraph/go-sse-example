package watcher

import (
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type Watcher struct {
	n int64
}

func (w *Watcher) OnStateChange(_ net.Conn, state http.ConnState) {
	switch state {
	case http.StateIdle, http.StateActive:
		// Do nothing.
	case http.StateNew:
		w.add(1)
	case http.StateHijacked, http.StateClosed:
		w.add(-1)
	}
	log.Printf("%v Connection Counts: %d", time.Now(), w.count())
}

func (w *Watcher) count() int {
	return int(atomic.LoadInt64(&w.n))
}

func (w *Watcher) add(c int64) {
	atomic.AddInt64(&w.n, c)
}
