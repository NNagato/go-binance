package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type ChangeUMInitialLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

func (s *ChangeUMInitialLeverageService) Symbol(symbol string) *ChangeUMInitialLeverageService {
	s.symbol = symbol
	return s
}

func (s *ChangeUMInitialLeverageService) Leverage(leverage int) *ChangeUMInitialLeverageService {
	s.leverage = leverage
	return s
}

// Do sends the request to change user's initial leverage
func (s *ChangeUMInitialLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *LeverageResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned, // or whatever the security type should be
	}

	// Required parameters
	r.setParam("symbol", s.symbol)
	r.setParam("leverage", s.leverage)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(LeverageResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type LeverageResponse struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
