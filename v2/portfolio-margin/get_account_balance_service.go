package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAccountBalanceService struct {
	c     *Client
	asset *string
}

func (s *GetAccountBalanceService) Asset(asset string) *GetAccountBalanceService {
	s.asset = &asset
	return s
}

// Do sends the request to get the account balance
func (s *GetAccountBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []AssetBalance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	if s.asset != nil {
		var rawRes AssetBalance
		err = json.Unmarshal(data, &rawRes)
		if err != nil {
			return nil, err
		}
		return []AssetBalance{rawRes}, nil
	}
	res = make([]AssetBalance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type AssetBalance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	UmWalletBalance     string `json:"umWalletBalance"`
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CmWalletBalance     string `json:"cmWalletBalance"`
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`
	NegativeBalance     string `json:"negativeBalance"`
	UpdateTime          int64  `json:"updateTime"`
}
