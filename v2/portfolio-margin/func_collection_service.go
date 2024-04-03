package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type FundingCollectionService struct {
	c     *Client
	asset string
}

func (s *FundingCollectionService) Asset(asset string) *FundingCollectionService {
	s.asset = asset
	return s
}

func (s *FundingCollectionService) Do(ctx context.Context, opts ...RequestOption) (*FundCollectionResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/asset-collection",
		secType:  secTypeSigned,
	}

	r.setParam("asset", s.asset)

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

type FundCollectionResponse struct {
	Msg string `json:"msg"`
}
