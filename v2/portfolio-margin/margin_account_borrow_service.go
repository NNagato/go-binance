package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type MarginAccountBorrowService struct {
	c      *Client
	asset  string
	amount string
}

func (s *MarginAccountBorrowService) Asset(asset string) *MarginAccountBorrowService {
	s.asset = asset
	return s
}

func (s *MarginAccountBorrowService) Amount(amount string) *MarginAccountBorrowService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginAccountBorrowService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccountBorrowResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/marginLoan",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccountBorrowResponse)
	err = json.Unmarshal(data, res)
	return res, err
}

// MarginAccountBorrowResponse define create order response
type MarginAccountBorrowResponse struct {
	TranId int64 `json:"tranId"`
}
