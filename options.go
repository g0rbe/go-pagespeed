package pagespeed

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
	URL         string   // The URL to fetch and analyze
	Key         string   // API Key
	Category    []string // A Lighthouse category to run; if none are given, only Performance category will be run.
	Locale      string   // The locale used to localize formatted results
	Strategy    string   // The analysis strategy (desktop or mobile) to use, and desktop is the default
	UTMCampaign string   // Campaign name for analytics.
	UTMSource   string   // Campaign source for analytics.
}

func (o *Options) RequestURL() string {

	v := ApiEndpoint

	v += "?url=" + o.URL

	if o.Key != "" {
		v += "&key=" + o.Key
	}

	for i := range o.Category {
		v += "&category=" + o.Category[i]
	}

	if o.Locale != "" {
		v += "&locale=" + o.Locale
	}

	if o.Strategy != "" {
		v += "&strategy=" + o.Strategy
	}

	if o.UTMCampaign != "" {
		v += "&utm_campaign=" + o.UTMCampaign
	}

	if o.UTMSource != "" {
		v += "&utm_source=" + o.UTMSource
	}

	return v
}
