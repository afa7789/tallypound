package compound

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/afa7789/tallypound/internal/domain"
)

func Proposals() ([]byte, error) {

	// Request the proposals from the Compound API with url compound_api_proposals
	// and store the response in the variable res.

	// make an http request to a random site
	res, err := http.Get(domain.Compound_api_proposals)
	if err != nil {
		return nil, err
	}

	// get the body out of the res
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//unmarshal the body into a struct
	var proposals []Proposal
	err = json.Unmarshal(body, &proposals)

	return body, nil
}
