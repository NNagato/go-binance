package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAutoRepayFuturesStatusService
type GetAutoRepayFuturesStatusService struct {
	c *Client
}

func (s *GetAutoRepayFuturesStatusService) Do(ctx context.Context, opts ...RequestOption) (res *AutoRepayFuturesStatusResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/repay-futures-switch",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(AutoRepayFuturesStatusResponse)
	err = json.Unmarshal(data, res)
	return
}

type AutoRepayFuturesStatusResponse struct {
	AutoRepay bool `json:"autoRepay"`
}

// ChangeAutoRepayFuturesStatusService
type ChangeAutoRepayFuturesStatusService struct {
	autoRepay string
	c         *Client
}

func (s *ChangeAutoRepayFuturesStatusService) AutoRepay(autoRepay string) *ChangeAutoRepayFuturesStatusService {
	s.autoRepay = autoRepay
	return s
}

// Do send request
func (s *ChangeAutoRepayFuturesStatusService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/repay-futures-switch",
		secType:  secTypeSigned,
	}
	r.setParam("autoRepay", s.autoRepay)

	_, _, err := s.c.callAPI(ctx, r, opts...)
	return err
}
