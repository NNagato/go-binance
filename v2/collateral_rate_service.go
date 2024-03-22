package binance

import (
	"context"
	"net/http"
)

type MarginCollateralRateService struct {
	c *Client
}

func (s *MarginCollateralRateService) Do(ctx context.Context, opts ...RequestOption) ([]MarginCollateralRate, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/portfolio/collateralRate",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]MarginCollateralRate, 0)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type MarginCollateralRate struct {
	Asset          string `json:"asset"`
	CollateralRate string `json:"collateralRate"`
}
