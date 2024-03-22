package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type TransferSideType string

const (
	TransferSideToUM   TransferSideType = "TO_UM"
	TransferSideFromUM TransferSideType = "FROM_UM"
)

type TransferBnbService struct {
	c            *Client
	amount       string
	transferSide TransferSideType
}

func (s *TransferBnbService) Amount(amount string) *TransferBnbService {
	s.amount = amount
	return s
}

func (s *TransferBnbService) TransferSide(transferSide TransferSideType) *TransferBnbService {
	s.transferSide = transferSide
	return s
}

func (s *TransferBnbService) Do(ctx context.Context, opts ...RequestOption) (*TransferBnbResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/bnb-transfer",
		secType:  secTypeSigned,
	}

	r.setParam("amount", s.amount)
	r.setParam("transferSide", s.transferSide)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(TransferBnbResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type TransferBnbResponse struct {
	TranID int `json:"tranId"`
}
