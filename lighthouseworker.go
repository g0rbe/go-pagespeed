package pagespeed

import (
	"context"
	"sync"
	"sync/atomic"
)

type LighthouseWorker struct {
	wg         *sync.WaitGroup
	m          *sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
	urlChan    chan string
	resultChan chan *LighthouseScores

	workersCancel []context.CancelFunc
	numTask       *atomic.Int64

	opts *Options
}

func NewLighthouseWorker(opt *Options, numWorker int, lenBuf int) *LighthouseWorker {

	lw := new(LighthouseWorker)

	lw.wg = new(sync.WaitGroup)
	lw.m = new(sync.RWMutex)

	lw.ctx, lw.cancel = context.WithCancel(context.Background())

	lw.urlChan = make(chan string, lenBuf)
	lw.resultChan = make(chan *LighthouseScores, lenBuf)

	lw.workersCancel = make([]context.CancelFunc, 0)
	lw.numTask = new(atomic.Int64)

	lw.opts = opt

	lw.StartWorkers(numWorker)

	return lw
}

func (w *LighthouseWorker) Put(u string) {

	defer w.numTask.Add(1)

	w.urlChan <- u

}

func (w *LighthouseWorker) Get() *LighthouseScores {

	defer w.numTask.Add(-1)

	return <-w.resultChan
}

func (w *LighthouseWorker) NumTask() int {

	return int(w.numTask.Load())
}

func (w *LighthouseWorker) NumWorker() int {
	w.m.RLock()
	defer w.m.RUnlock()

	return len(w.workersCancel)
}

func (w *LighthouseWorker) worker(ctx context.Context) {

	defer w.wg.Done()

	for {

		select {

		case u := <-w.urlChan:
			res, err := RunLighthouse(u, w.opts)
			if err != nil {
				w.resultChan <- &LighthouseScores{URL: u, Error: err}
				continue
			}

			w.resultChan <- res

		case <-ctx.Done():
			return
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *LighthouseWorker) StartWorker() {
	w.m.Lock()
	defer w.m.Unlock()

	w.wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	w.workersCancel = append(w.workersCancel, cancel)

	go w.worker(ctx)
}

func (w *LighthouseWorker) StopWorker() bool {

	w.m.Lock()
	defer w.m.Unlock()

	if len(w.workersCancel) == 0 {
		return false
	}

	cancel := w.workersCancel[0]

	w.workersCancel = w.workersCancel[1:]

	cancel()

	return true
}

func (w *LighthouseWorker) StartWorkers(n int) {

	for i := 0; i < n; i++ {

		w.StartWorker()
	}
}

func (w *LighthouseWorker) Close() error {

	close(w.urlChan)
	w.cancel()

	w.wg.Wait()

	close(w.resultChan)

	return nil
}
