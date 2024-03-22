package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type MarginAccountRepayService struct {
	c      *Client
	asset  string
	amount string
}

func (s *MarginAccountRepayService) Asset(asset string) *MarginAccountRepayService {
	s.asset = asset
	return s
}

func (s *MarginAccountRepayService) Amount(amount string) *MarginAccountRepayService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginAccountRepayService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountRepayResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/marginLoan",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccountRepayResponse)
	err = json.Unmarshal(data, res)
	return res, err
}

// MarginAccountRepayResponse define create order response
type MarginAccountRepayResponse struct {
	TranId int64 `json:"tranId"`
}
