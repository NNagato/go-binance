package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetNegativeBalanceInterestHistoryService struct {
	c         *Client
	asset     *string
	startTime *int64
	endTime   *int64
	size      *int
}

func (s *GetNegativeBalanceInterestHistoryService) Asset(asset string) *GetNegativeBalanceInterestHistoryService {
	s.asset = &asset
	return s
}

func (s *GetNegativeBalanceInterestHistoryService) StartTime(startTime int64) *GetNegativeBalanceInterestHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetNegativeBalanceInterestHistoryService) EndTime(endTime int64) *GetNegativeBalanceInterestHistoryService {
	s.endTime = &endTime
	return s
}

func (s *GetNegativeBalanceInterestHistoryService) Size(size int) *GetNegativeBalanceInterestHistoryService {
	s.size = &size
	return s
}

// Do sends the request to get the history of negative balance interest.
func (s *GetNegativeBalanceInterestHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []NegativeBalanceInterestHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/portfolio/interest-history",
		secType:  secTypeSigned,
	}

	// Set up the parameters
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}

	// Make the API call and handle the response
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]NegativeBalanceInterestHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type NegativeBalanceInterestHistory struct {
	Asset               string `json:"asset"`
	Interest            string `json:"interest"`
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	InterestRate        string `json:"interestRate"`
	Principal           string `json:"principal"`
}
