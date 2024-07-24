package pagespeed

import "fmt"

type LighthouseScores struct {
	URL           string `json:",omitempty" yaml:",omitempty"`
	Performance   int    `json:",omitempty" yaml:",omitempty"`
	Accessibility int    `json:",omitempty" yaml:",omitempty"`
	BestPractices int    `json:",omitempty" yaml:",omitempty"`
	SEO           int    `json:",omitempty" yaml:",omitempty"`
	PWA           int    `json:",omitempty" yaml:",omitempty"`
	Error         error  `json:",omitempty" yaml:",omitempty"`
}

func RunLighthouse(u string, opts *Options) (*LighthouseScores, error) {

	res, err := RunPagespeed(u, opts)

	return res.LighthouseScores(), err
}

func (l LighthouseScores) String() string {

	v := fmt.Sprintf("%s\n", l.URL)

	if l.Performance > 0 {
		v += fmt.Sprintf("\t- Performance:   %d\n", l.Performance)
	}

	if l.Accessibility > 0 {
		v += fmt.Sprintf("\t- Accessibility: %d\n", l.Accessibility)
	}

	if l.BestPractices > 0 {
		v += fmt.Sprintf("\t- BestPractices: %d\n", l.BestPractices)
	}

	if l.SEO > 0 {
		v += fmt.Sprintf("\t- SEO:           %d\n", l.SEO)
	}

	if l.PWA > 0 {
		v += fmt.Sprintf("\t- PWA:           %d\n", l.PWA)
	}

	return v
}

func (l *LighthouseScores) Total() int {
	return l.Performance + l.Accessibility + l.BestPractices + l.SEO + l.PWA
}
