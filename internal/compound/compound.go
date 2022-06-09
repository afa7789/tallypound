package compound

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afa7789/tallypound/internal/domain"
)

func Proposals() ([]domain.Proposal, error) {
	var proposals []domain.Proposal
	page := 1

	// do while
	for ok := true; ok; {
		// create query
		query := fmt.Sprintf("%s/?page_number=%d&page_size=10", domain.CompoundAPIProposals, page)
		// Request the proposals from the Compound API with url compound_api_proposals
		// and store the response in the variable res.
		res, err := http.Get(query)
		if err != nil {
			print("oi2\n")

			return nil, err
		}

		// get the body out of the res
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			print("oi3\n")

			return nil, err
		}

		// unmarshal the body into a struct
		var pr domain.ProposalsResponse

		err = json.Unmarshal(body, &pr)
		if err != nil {
			fmt.Printf("%s", string(body))
			print("oi4\n")

			return nil, err
		}

		proposals = append(proposals, pr.Proposals...)

		// stop when ok is equal to false
		if pr.PaginationSumarry.PageNumber < pr.PaginationSumarry.TotalPages {
			ok = true
			page++
		} else {
			ok = false
		}
	}

	return proposals, nil
}
