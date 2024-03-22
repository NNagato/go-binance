package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAllUMOrdersService struct {
	c         *Client
	symbol    *string
	orderID   *string
	startTime *int64
	endTime   *int64
	limit     *int
}

func (s *GetAllUMOrdersService) Symbol(symbol string) *GetAllUMOrdersService {
	s.symbol = &symbol
	return s
}

func (s *GetAllUMOrdersService) OrderID(orderID string) *GetAllUMOrdersService {
	s.orderID = &orderID
	return s
}

func (s *GetAllUMOrdersService) StartTime(startTime int64) *GetAllUMOrdersService {
	s.startTime = &startTime
	return s
}

func (s *GetAllUMOrdersService) EndTime(endTime int64) *GetAllUMOrdersService {
	s.endTime = &endTime
	return s
}

func (s *GetAllUMOrdersService) Limit(limit int) *GetAllUMOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAllUMOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []UMOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/allOrders",
		secType:  secTypeSigned,
	}

	// Required parameters
	r.setParam("symbol", s.symbol)

	// Optional parameters
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
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
	res = make([]UMOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
