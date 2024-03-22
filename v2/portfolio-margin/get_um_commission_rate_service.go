package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMCommissionRate struct {
	c      *Client
	symbol string
}

func (s *GetUMCommissionRate) Symbol(symbol string) *GetUMCommissionRate {
	s.symbol = symbol
	return s
}

// Do sends the request to get the user's margin commission rate
func (s *GetUMCommissionRate) Do(ctx context.Context, opts ...RequestOption) (res *CommissionRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/commissionRate",
		secType:  secTypeSigned,
	}

	// Set up the mandatory parameters
	r.setParam("symbol", s.symbol)

	// Make the API call and handle the response
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CommissionRate)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
