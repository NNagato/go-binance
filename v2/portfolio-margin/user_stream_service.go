package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

type StartUserStreamDataService struct {
	c *Client
}

func (s *StartUserStreamDataService) Do(ctx context.Context, opts ...RequestOption) (*ListenKeyResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/listenKey",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ListenKeyResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type KeepAliveUserStreamDataService struct {
	c *Client
}

func (s *KeepAliveUserStreamDataService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/papi/v1/listenKey",
		secType:  secTypeSigned,
	}

	_, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

type CloseUserStreamDataService struct {
	c *Client
}

func (s *CloseUserStreamDataService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/listenKey",
		secType:  secTypeSigned,
	}

	_, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}
