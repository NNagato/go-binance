package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMPositionRiskService struct {
	c      *Client
	symbol *string
}

func (s *GetUMPositionRiskService) Symbol(symbol string) *GetUMPositionRiskService {
	s.symbol = &symbol
	return s
}

// Do sends the request to get the account balance
func (s *GetUMPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []PositonRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/positionRisk",
		secType:  secTypeSigned, // or whatever the security type should be
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]PositonRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PositonRisk struct {
	EntryPrice       string `json:"entryPrice"`
	Leverage         string `json:"leverage"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Notional         string `json:"notional"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	LiquidationPrice string `json:"liquidationPrice"`
	PositionSide     string `json:"positionSide"`
	UpdateTime       int64  `json:"updateTime"`
}
