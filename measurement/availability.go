package measurement

import (
	"fmt"
	"net/http"
)

type availability struct {
	name          string
	summaryPhrase string
	url           string
	client        *http.Client
}

func (a *availability) Name() string {
	return a.name
}

func (a *availability) SummaryPhrase() string {
	return a.summaryPhrase
}

func (a *availability) PerformMeasurement() (string, string, string, bool) {
	res, err := a.client.Get(a.url)
	if err != nil {
		return err.Error(), "", "", false
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Sprintf("response had status %d; %s", res.StatusCode, res.Status), "", "", false
	}

	return "", "", "", true
}
