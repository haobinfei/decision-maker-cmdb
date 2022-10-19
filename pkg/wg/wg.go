package wg

import "sync"

type Wg struct {
	sync.WaitGroup
}

func (w *Wg) Exec(function func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()
}
