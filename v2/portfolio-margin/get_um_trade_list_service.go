package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMTradeList struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	fromId    *int64
	limit     *int
}

func (s *GetUMTradeList) Symbol(symbol string) *GetUMTradeList {
	s.symbol = symbol
	return s
}

func (s *GetUMTradeList) StartTime(startTime int64) *GetUMTradeList {
	s.startTime = &startTime
	return s
}

func (s *GetUMTradeList) EndTime(endTime int64) *GetUMTradeList {
	s.endTime = &endTime
	return s
}

func (s *GetUMTradeList) FromId(fromId int64) *GetUMTradeList {
	s.fromId = &fromId
	return s
}

func (s *GetUMTradeList) Limit(limit int) *GetUMTradeList {
	s.limit = &limit
	return s
}

// Do sends the request to get the user's margin trades
func (s *GetUMTradeList) Do(ctx context.Context, opts ...RequestOption) (res []Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/userTrades",
		secType:  secTypeSigned,
	}

	// Set up the parameters
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	// Make the API call and handle the response
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Trade struct {
	Symbol          string `json:"symbol"`
	ID              int    `json:"id"`
	OrderID         int    `json:"orderId"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	MarginAsset     string `json:"marginAsset"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
	PositionSide    string `json:"positionSide"`
}
