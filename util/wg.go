package util

import "sync"

type WaitGroup struct {
	wg  sync.WaitGroup
	err error
}

func (wg *WaitGroup) Go(fn func() error) *WaitGroup {
	if wg.err != nil {
		return wg
	}
	wg.wg.Add(1)
	go func() {
		defer wg.wg.Done()
		if err := fn(); err != nil {
			wg.stop(err)
		}
	}()
	return wg
}

func (wg *WaitGroup) stop(err error) {
	wg.err = err
}

func (wg *WaitGroup) Wait() error {
	wg.wg.Wait()
	err := wg.err
	wg.err = nil
	return err
}
