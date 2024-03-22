package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAccountInfoService struct {
	c *Client
}

// Do sends the request to get the account balance
func (s *GetAccountInfoService) Do(ctx context.Context, opts ...RequestOption) (res *AccountInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/account",
		secType:  secTypeSigned, // or whatever the security type should be
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(AccountInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type AccountInfo struct {
	UniMMR                   string `json:"uniMMR"`
	AccountEquity            string `json:"accountEquity"`
	ActualEquity             string `json:"actualEquity"`
	AccountInitialMargin     string `json:"accountInitialMargin"`
	AccountMaintMargin       string `json:"accountMaintMargin"`
	AccountStatus            string `json:"accountStatus"`
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"`
	TotalAvailableBalance    string `json:"totalAvailableBalance"`
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"`
	UpdateTime               int64  `json:"updateTime"`
}
