package pagespeed_test

import (
	"fmt"
	"testing"

	"github.com/g0rbe/go-pagespeed"
)

func TestRun(t *testing.T) {

	r, err := pagespeed.Run(pagespeed.Options{URL: "https://gorbe.io/about", Category: pagespeed.CategoryAll})
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("%#v\n", r)
}

func ExampleRun() {

	r, err := pagespeed.Run(pagespeed.Options{URL: "https://github.com/", Category: pagespeed.CategoryAll})
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v\n", r)
}

func TestLighthouseScores(t *testing.T) {
	r, err := pagespeed.Run(pagespeed.Options{URL: "https://gorbe.io/about", Category: pagespeed.CategoryAll})
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("\n%s\n", r.LighthouseScores())
}
