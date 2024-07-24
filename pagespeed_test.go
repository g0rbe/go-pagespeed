package pagespeed_test

import (
	"fmt"
	"testing"

	"github.com/g0rbe/go-pagespeed"
)

func TestRunPagespeed(t *testing.T) {

	_, err := pagespeed.RunPagespeed(&pagespeed.Options{URL: "https://gorbe.io/about", Category: pagespeed.CategoryAll})
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	//t.Logf("%#v\n", r)
}

func ExampleRunPagespeed() {

	r, err := pagespeed.RunPagespeed(&pagespeed.Options{URL: "https://github.com/", Category: pagespeed.CategoryAll})
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v\n", r)
}

func TestRunLighthouse(t *testing.T) {

	r, err := pagespeed.RunLighthouse(&pagespeed.Options{URL: "https://gorbe.io/about", Category: pagespeed.CategoryAll})
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("\n%s\n", r)
}

func ExampleRunLighthouse() {

	r, err := pagespeed.RunLighthouse(&pagespeed.Options{URL: "https://github.com/", Category: pagespeed.CategoryAll})
	if err != nil {
		// handle error
	}

	fmt.Printf("%#v\n", r)
}
