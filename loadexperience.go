package pagespeed

type Distribution struct {
	Min        int     `json:"min"`
	Max        int     `json:"max"`
	Proportion float32 `json:"proportion"`
}

type Metric struct {
	Percentile    int            `json:"percentile"`
	Distributions []Distribution `json:"distributions"`
	Category      int            `json:"category"`
}

type LoadingExperience struct {
	ID              string            `json:"id"`
	Metrics         map[string]Metric `json:"metrics"`
	OverallCategory string            `json:"overall_category"`
	InitialURL      string            `json:"initial_url"`
}
