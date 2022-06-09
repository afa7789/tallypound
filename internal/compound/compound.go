package compound

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afa7789/tallypound/internal/cache"
	"github.com/afa7789/tallypound/internal/domain"
)

type CacheCompound struct {
	cache *cache.Cache
}

func NewCacheCompound(cache *cache.Cache) *CacheCompound {
	return &CacheCompound{
		cache: cache,
	}
}

// Proposals gets the list of proposals from the compound API.
// by running a do while and using their pagination.
func (cc *CacheCompound) Proposals() ([]domain.Proposal, error) {
	var proposals []domain.Proposal
	page := 1

	// what I would do to improve, use concurrency and paralelism on the following pages after the first one.
	// do while
	for ok := true; ok; {
		var body []byte
		// create query
		query := fmt.Sprintf("%s/?page_number=%d&page_size=10", domain.CompoundAPIProposals, page)
		// Request the proposals from the Compound API with url compound_api_proposals
		// and store the response in the variable res.
		res, err := http.Get(query)

		if err != nil {
			// get from cache
			b, boolItem := cc.cache.Get(query)
			if !boolItem {
				return nil, err
			}
			// cast b to bytes and set it to body
			body = b.([]byte)
		} else {
			// get the body out of the res
			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			// set on cache
			cc.cache.SetWithoutExpiration(query, body)
		}

		// unmarshal the body into a struct
		var pr domain.ProposalsResponse

		err = json.Unmarshal(body, &pr)
		if err != nil {
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

// Stats with the compound API answer mount a new struct with the stats.
func (cc *CacheCompound) Stats() (*domain.Stats, error) {
	stats := &domain.Stats{}
	page := 1

	// what I would do to improve, use concurrency and paralelism on the following pages after the first one.
	// do while
	for ok := true; ok; {
		// create query
		query := fmt.Sprintf("%s/?page_number=%d&page_size=10", domain.CompoundAPIProposals, page)
		// Request the proposals from the Compound API with url compound_api_proposals
		// and store the response in the variable res.
		res, err := http.Get(query)
		var body []byte
		if err != nil {
			// get from cache
			b, boolItem := cc.cache.Get(query)
			if !boolItem {
				return nil, err
			}
			// cast b to bytes and set it to body
			body = b.([]byte)
		} else {
			// get the body out of the res
			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			// set on cache
			cc.cache.SetWithoutExpiration(query, body)
		}

		// unmarshal the body into a struct
		var pr domain.ProposalsResponse

		err = json.Unmarshal(body, &pr)
		if err != nil {
			return nil, err
		}

		// calculate for the created stats above
		timestamp := 0
		olderState := domain.State{}
		// range proposals
		for _, proposal := range pr.Proposals {
			// check which state is the older, bigger timestamp on start.
			for _, state := range proposal.States {
				// if state is equal to "active"
				if timestamp < state.StartTime {
					timestamp = state.StartTime
					olderState = state
				}
			}

			// swich case each state
			// stats: pending , active ,canceled , defeated , succeeded , queued , expired , executed .
			switch olderState.State {
			case "pending":
				stats.Pending++
			case "active":
				stats.Active++
			case "canceled":
				stats.Canceled++
			case "defeated":
				stats.Defeated++
			case "succeeded":
				stats.Succeeded++
			case "queued":
				stats.Queued++
			case "expired":
				stats.Expired++
			case "executed":
				stats.Executed++
			}
		}

		// stop when ok is equal to false
		if pr.PaginationSumarry.PageNumber < pr.PaginationSumarry.TotalPages {
			ok = true
			page++
		} else {
			ok = false
		}
	}

	return stats, nil
}
