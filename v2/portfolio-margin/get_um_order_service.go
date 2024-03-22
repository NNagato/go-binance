package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetUMOrderService struct {
	c                 *Client
	symbol            *string
	orderID           *string
	origClientOrderID *string
}

func (s *GetUMOrderService) Symbol(symbol string) *GetUMOrderService {
	s.symbol = &symbol
	return s
}

func (s *GetUMOrderService) OrderID(orderID string) *GetUMOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetUMOrderService) OrigClientOrderID(origClientOrderID string) *GetUMOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *UMOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMOrder)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type UMOrder struct {
	ClientOrderID           string                  `json:"clientOrderId"`
	CumQty                  string                  `json:"cumQty"`
	CumQuote                string                  `json:"cumQuote"`
	ExecutedQty             string                  `json:"executedQty"`
	OrderID                 int                     `json:"orderId"`
	AvgPrice                string                  `json:"avgPrice"`
	OrigQty                 string                  `json:"origQty"`
	Price                   string                  `json:"price"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	Status                  OrderStatusType         `json:"status"`
	Symbol                  string                  `json:"symbol"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	Type                    OrderType               `json:"type"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate"`
	UpdateTime              int64                   `json:"updateTime"`
}
