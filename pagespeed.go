package pagespeed

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ApiEndpoint = "https://www.googleapis.com/pagespeedonline/v5/runPagespeed"
)

type Result struct {
	CaptchaResult           string            `json:"captchaResult"`
	Kind                    string            `json:"kind"`
	ID                      string            `json:"id"`
	LoadingExperience       LoadingExperience `json:"loadingExperience"`
	OriginLoadingExperience LoadingExperience `json:"originLoadingExperience"`
	LighthouseResult        LighthouseResult  `json:"lighthouseResult"`
	AnalysisUTCTimestamp    time.Time         `json:"analysisUTCTimestamp"`
}

func RunPagespeed(opt *Options) (*Result, error) {

	resp, err := http.Get(opt.RequestURL())
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	if resp.StatusCode != 200 {

		var e RuntimeError
		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal RuntimeError: %w", err)
		}

		return nil, e
	}

	res := new(Result)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return res, nil
}

func (r *Result) LighthouseScores() *LighthouseScores {

	if r == nil {
		return nil
	}

	v := new(LighthouseScores)

	v.URL = r.ID

	for k := range r.LighthouseResult.Categories {

		switch k {
		case "performance":
			v.Performance = int(r.LighthouseResult.Categories[k].Score * 100)
		case "accessibility":
			v.Accessibility = int(r.LighthouseResult.Categories[k].Score * 100)
		case "best-practices":
			v.BestPractices = int(r.LighthouseResult.Categories[k].Score * 100)
		case "seo":
			v.SEO = int(r.LighthouseResult.Categories[k].Score * 100)
		case "pwa":
			v.PWA = int(r.LighthouseResult.Categories[k].Score * 100)
		default:
			panic(fmt.Sprintf("invalid lighthouse category: %s", k))
		}
	}

	return v
}
