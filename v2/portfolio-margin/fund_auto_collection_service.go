package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type FundAutoCollectionService struct {
	c *Client
}

func (s *FundAutoCollectionService) Do(ctx context.Context, opts ...RequestOption) (*FundCollectionResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/auto-collection",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(FundCollectionResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
