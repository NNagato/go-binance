package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMIncomeHistoryService struct {
	c          *Client
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	limit      *int
}

func (s *GetUMIncomeHistoryService) Symbol(symbol string) *GetUMIncomeHistoryService {
	s.symbol = &symbol
	return s
}

func (s *GetUMIncomeHistoryService) IncomeType(incomeType string) *GetUMIncomeHistoryService {
	s.incomeType = &incomeType
	return s
}

func (s *GetUMIncomeHistoryService) StartTime(startTime int64) *GetUMIncomeHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetUMIncomeHistoryService) EndTime(endTime int64) *GetUMIncomeHistoryService {
	s.endTime = &endTime
	return s
}

func (s *GetUMIncomeHistoryService) Limit(limit int) *GetUMIncomeHistoryService {
	s.limit = &limit
	return s
}

func (s *GetUMIncomeHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]IncomeHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/income",
		secType:  secTypeSigned,
	}

	// Set optional parameters if they are not nil
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.incomeType != nil {
		r.setParam("incomeType", *s.incomeType)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]IncomeHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type IncomeHistory struct {
	Symbol     string `json:"symbol"`
	IncomeType string `json:"incomeType"`
	Income     string `json:"income"`
	Asset      string `json:"asset"`
	Info       string `json:"info"`
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`
	TradeID    string `json:"tradeId"`
}
