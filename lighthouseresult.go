package pagespeed

type Environment struct {
	NetworkUserAgent string  `json:"networkUserAgent"`
	HostUserAgent    string  `json:"hostUserAgent"`
	BenchmarkIndex   float32 `json:"benchmarkIndex"`
}

type ConfigSettings struct {
	EmulatedFormFactor string   `json:"emulatedFormFactor"`
	FormFactor         string   `json:"formFactor"`
	Locale             string   `json:"locale"`
	FinalUrl           string   `json:"finalUrl"`
	OnlyCategories     []string `json:"onlyCategories"`
	Channel            string   `json:"channel"`
}

type Audit struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Score            float32 `json:"score"`
	ScoreDisplayMode string  `json:"scoreDisplayMode"`
}

type AuditRef struct {
	ID     string  `json:"id"`
	Weight float32 `json:"weight"`
	Group  string  `json:"group"`
}

type Category struct {
	ID                string     `json:"id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Score             float32    `json:"score"`
	ManualDescription string     `json:"manualDescription"`
	AuditRefs         []AuditRef `json:"auditRefs"`
}

type CategoryGroup struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Timing struct {
	Total float32 `json:"total"`
}

type Entity struct {
	Name           string   `json:"name"`
	IsFirstParty   bool     `json:"isFirstParty"`
	IsUnrecognized bool     `json:"isUnrecognized"`
	Origins        []string `json:"origins"`
}

type LighthouseResult struct {
	RequestedUrl      string                   `json:"requestedUrl"`
	FinalUrl          string                   `json:"finalUrl"`
	MainDocumentUrl   string                   `json:"mainDocumentUrl"`
	FinalDisplayedUrl string                   `json:"finalDisplayedUrl"`
	LighthouseVersion string                   `json:"lighthouseVersion"`
	UserAgent         string                   `json:"userAgent"`
	FetchTime         string                   `json:"fetchTime"`
	Environment       Environment              `json:"environment"`
	RunWarnings       []string                 `json:"runWarnings"`
	ConfigSettings    ConfigSettings           `json:"configSettings"`
	Audits            map[string]Audit         `json:"audits"`
	Categories        map[string]Category      `json:"categories"`
	CategoryGroups    map[string]CategoryGroup `json:"categoryGroups"`
	RuntimeError      RuntimeError             `json:"runtimeError"`
	Timing            Timing                   `json:"timing"`
	Entities          []Entity                 `json:"entities"`
}
