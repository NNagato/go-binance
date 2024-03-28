package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMLeverageService struct {
	c      *Client
	symbol *string
}

func (s *GetUMLeverageService) Symbol(symbol string) *GetUMLeverageService {
	s.symbol = &symbol
	return s
}

// Do sends the request to get the account balance
func (s *GetUMLeverageService) Do(ctx context.Context, opts ...RequestOption) (res []Leverage, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/leverageBracket",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	res = make([]Leverage, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Leverage struct {
	Symbol       string    `json:"symbol"`
	NotionalCoef string    `json:"notionalCoef"`
	Brackets     []Bracket `json:"brackets"`
}

type Bracket struct {
	Bracket          int     `json:"bracket"`
	InitialLeverage  int     `json:"initialLeverage"`
	NotionalCap      float64 `json:"notionalCap"`
	NotionalFloor    float64 `json:"notionalFloor"`
	MaintMarginRatio float64 `json:"maintMarginRatio"`
	Cum              float64 `json:"cum"`
}
