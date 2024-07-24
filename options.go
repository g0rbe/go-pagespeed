package pagespeed

import "sync"

// Possible values for category paramater
const (
	CategoryAccessibility = "ACCESSIBILITY"
	CategoryBestPractices = "BEST_PRACTICES"
	CategoryPerformance   = "PERFORMANCE"
	CategorySEO           = "SEO"
	CategoryPWA           = "PWA"
)

// Possible values for strategy paramater
const (
	StrategyDesktop = "dektop"
	StrategyMobile  = "mobile"
)

var (
	// Shorthand to include every category
	CategoryAll = []string{CategoryAccessibility, CategoryBestPractices, CategoryPerformance, CategorySEO, CategoryPWA}
)

type Options struct {
	key         string   // API Key
	category    []string // A Lighthouse category to run; if none are given, only Performance category will be run.
	locale      string   // The locale used to localize formatted results
	strategy    string   // The analysis strategy (desktop or mobile) to use, and desktop is the default
	utmCampaign string   // Campaign name for analytics.
	utmSource   string   // Campaign source for analytics.
	m           *sync.RWMutex
}

func NewOptions() *Options {

	o := new(Options)
	o.m = new(sync.RWMutex)

	return o
}

func FullAnalysis() *Options {

	o := NewOptions()
	o.SetCategories(CategoryAll)

	return o
}

func FullAnalysisWithKey(k string) *Options {

	o := NewOptions()
	o.SetCategories(CategoryAll)
	o.SetKey(k)

	return o
}

func (o *Options) SetKey(v string) {

	o.m.Lock()

	o.key = v

	o.m.Unlock()
}

func (o *Options) SetCategories(c []string) {

	o.m.Lock()

	o.category = c

	o.m.Unlock()
}

// RequestURL creates the request URL to analyze URL u.
func (o *Options) RequestURL(u string) string {

	o.m.RLock()
	defer o.m.RUnlock()

	v := ApiEndpoint

	v += "?url=" + u

	if o.key != "" {
		v += "&key=" + o.key
	}

	for i := range o.category {
		v += "&category=" + o.category[i]
	}

	if o.locale != "" {
		v += "&locale=" + o.locale
	}

	if o.strategy != "" {
		v += "&strategy=" + o.strategy
	}

	if o.utmCampaign != "" {
		v += "&utm_campaign=" + o.utmCampaign
	}

	if o.utmSource != "" {
		v += "&utm_source=" + o.utmSource
	}

	return v
}
