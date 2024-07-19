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

func Run(opt Options) (*Result, error) {

	resp, err := http.Get(opt.RequestURL())
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	res := new(Result)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return res, nil
}
