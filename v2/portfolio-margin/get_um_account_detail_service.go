package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMAccountDetailService struct {
	c *Client
}

func (s *GetUMAccountDetailService) Do(ctx context.Context, opts ...RequestOption) (*AccountDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/account",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(AccountDetail)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Asset struct {
	Asset                  string `json:"asset"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	UpdateTime             int64  `json:"updateTime"`
}

type Position struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int    `json:"updateTime"`
	BreakEvenPrice         string `json:"breakEvenPrice"`
}

type AccountDetail struct {
	TradeGroupID int        `json:"tradeGroupId"`
	Assets       []Asset    `json:"assets"`
	Positions    []Position `json:"positions"`
}
