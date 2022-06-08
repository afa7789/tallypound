package compound

import (
	"net/http"

	"github.com/afa7789/tallypound/internal/domain"
)

func Proposals() (*http.Response, error) {

	// Request the proposals from the Compound API with url compound_api_proposals
	// and store the response in the variable res.

	// make an http request to a random site
	res, err := http.Get(domain.Compound_api_proposals)
	if err != nil {
		return nil, err
	}

	return res, nil
}
