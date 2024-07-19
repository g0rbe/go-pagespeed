# go-pagespeed

[![Go Report Card](https://goreportcard.com/badge/github.com/g0rbe/go-pagespeed)](https://goreportcard.com/report/github.com/g0rbe/go-pagespeed)
[![Go Reference](https://pkg.go.dev/badge/github.com/g0rbe/go-pagespeed.svg)](https://pkg.go.dev/github.com/g0rbe/go-pagespeed)

Golang module to the PageSpeed Insights API v5.

Google Docs: [Get Started with the PageSpeed Insights API](https://developers.google.com/speed/docs/insights/v5/get-started)

Get:
```bash
go get github.com/g0rbe/go-pagespeed@latest
```

Get the latest tag (if Go module proxy is not updated):
```bash
go get "github.com/g0rbe/go-pagespeed@$(curl -s 'https://api.github.com/repos/g0rbe/go-pagespeed/tags' | jq -r '.[0].name')"
```

Get the latest commit (if Go module proxy is not updated):
```bash
go get "github.com/g0rbe/go-pagespeed@$(curl -s 'https://api.github.com/repos/g0rbe/go-pagespeed/commits' | jq -r '.[0].sha')"
```

## TODO

- `LighthouseResult.i18n`
- `LighthouseResult.fullPageScreenshot`