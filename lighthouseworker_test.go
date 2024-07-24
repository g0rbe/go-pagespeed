package pagespeed_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/g0rbe/go-pagespeed"
)

func TestLighthouseWorker(t *testing.T) {

	w := pagespeed.NewLighthouseWorker(pagespeed.FullAnalysisWithKey(os.Getenv("GOOGLE_CLOUD_KEY")), 24, 100)

	for i := 0; i < 100; i++ {
		w.Put(TestURLS[rand.Intn(len(TestURLS))])
	}

	tasksDone := 0

	go func() {

		time.Sleep(10 * time.Second)
		w.StartWorkers(4)

		for w.NumWorker() > 0 {

			time.Sleep(time.Duration(rand.Intn(60)) * time.Second)
			if !w.StopWorker() {
				fmt.Fprintf(os.Stderr, "Failed to stop worker, starting...\n")
				w.StartWorker()
			}
		}
	}()

	for w.NumTask() > 0 {

		r := w.Get()
		if r.Error != nil {
			t.Logf("%s %s\n", r.URL, r.Error)
		} else {
			t.Logf("%s %d\n", r.URL, r.Total())
		}
		tasksDone++

		t.Logf("Tasks: %d, Done: %d\n", w.NumTask(), tasksDone)
		t.Logf("Workers: %d\n", w.NumWorker())

	}

	w.Close()

}
