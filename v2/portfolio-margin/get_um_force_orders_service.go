package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMForceOrdersService struct {
	c             *Client
	symbol        *string
	autoCloseType *string
	startTime     *int64
	endTime       *int64
	limit         *int
}

// Symbol sets the symbol parameter of the request.
func (s *GetUMForceOrdersService) Symbol(symbol string) *GetUMForceOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType sets the autoCloseType parameter of the request.
func (s *GetUMForceOrdersService) AutoCloseType(autoCloseType string) *GetUMForceOrdersService {
	s.autoCloseType = &autoCloseType
	return s
}

// StartTime sets the startTime parameter of the request.
func (s *GetUMForceOrdersService) StartTime(startTime int64) *GetUMForceOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter of the request.
func (s *GetUMForceOrdersService) EndTime(endTime int64) *GetUMForceOrdersService {
	s.endTime = &endTime
	return s
}

// Limit sets the limit parameter of the request.
func (s *GetUMForceOrdersService) Limit(limit int) *GetUMForceOrdersService {
	s.limit = &limit
	return s
}

// Do sends the request to get the force orders.
func (s *GetUMForceOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []UMOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/forceOrders",
		secType:  secTypeSigned,
	}

	// Set up the parameters
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.autoCloseType != nil {
		r.setParam("autoCloseType", *s.autoCloseType)
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

	// Make the API call and handle the response
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
