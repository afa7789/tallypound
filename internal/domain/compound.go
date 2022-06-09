package domain

const (
	compound    = "https://api.compound.finance"
	compoundAPI = "/api/v2"
	proposals   = "/governance/proposals"
	// or
	CompoundAPIProposals = compound + compoundAPI + proposals
)

type Proposal struct {
	AgainstVotes string  `json:"against_votes"` // this is a float representation
	Description  string  `json:"description"`   // just a text
	ForVotes     string  `json:"for_votes"`     // this is a float representation
	ID           int     `json:"id"`            // this is a string representation
	Title        string  `json:"title"`         // just a text
	States       []State `json:"states"`
}

type PaginationSumarry struct {
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalPages   int `json:"total_pages"`
	TotalEntries int `json:"total_entries"`
}

type State struct {
	EndTime   int    `json:"end_time"`
	StartTime int    `json:"start_time"`
	State     string `json:"state"`
	TrxHash   string `json:"trx_hash"`
}

type Request struct {
	Network     string `json:"network"`
	PageNumber  int    `json:"page_number"`
	PageSize    int    `json:"page_size"`
	ProposalIds []int  `json:"proposal_ids"`
	State       string `json:"state"`
	WithDetail  bool   `json:"with_detail"`
}

type ProposalsResponse struct {
	Error             interface{}       `json:"error"` // try to treat it as an integer is a good away
	PaginationSumarry PaginationSumarry `json:"pagination_summary"`
	Proposals         []Proposal        `json:"proposals"`
	Request           Request           `json:"request"`
}

// stats: pending , active ,canceled , defeated , succeeded , queued , expired , executed .
type Stats struct {
	Pending   int `json:"pending"`
	Active    int `json:"active"`
	Canceled  int `json:"canceled"`
	Defeated  int `json:"defeated"`
	Succeeded int `json:"succeeded"`
	Queued    int `json:"queued"`
	Expired   int `json:"expired"`
	Executed  int `json:"executed"`
}
