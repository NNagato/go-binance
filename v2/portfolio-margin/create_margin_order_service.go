package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateMarginOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	stopPrice               *string
	newClientOrderID        *string
	newOrderRespType        *NewOrderRespType
	icebergQty              *string
	sideEffectType          *SideEffectType
	selfTradePreventionMode *SelfTradePreventionMode
}

func (s *CreateMarginOrderService) Symbol(symbol string) *CreateMarginOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateMarginOrderService) Side(side SideType) *CreateMarginOrderService {
	s.side = side
	return s
}

func (s *CreateMarginOrderService) Type(orderType OrderType) *CreateMarginOrderService {
	s.orderType = orderType
	return s
}

func (s *CreateMarginOrderService) TimeInForce(timeInForce TimeInForceType) *CreateMarginOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateMarginOrderService) Quantity(quantity string) *CreateMarginOrderService {
	s.quantity = &quantity
	return s
}

func (s *CreateMarginOrderService) QuoteOrderQty(quoteOrderQty string) *CreateMarginOrderService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *CreateMarginOrderService) Price(price string) *CreateMarginOrderService {
	s.price = &price
	return s
}

func (s *CreateMarginOrderService) StopPrice(stopPrice string) *CreateMarginOrderService {
	s.stopPrice = &stopPrice
	return s
}

func (s *CreateMarginOrderService) NewClientOrderID(newClientOrderID string) *CreateMarginOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *CreateMarginOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateMarginOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *CreateMarginOrderService) IcebergQty(icebergQty string) *CreateMarginOrderService {
	s.icebergQty = &icebergQty
	return s
}

func (s *CreateMarginOrderService) SideEffectType(sideEffectType SideEffectType) *CreateMarginOrderService {
	s.sideEffectType = &sideEffectType
	return s
}

func (s *CreateMarginOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateMarginOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateMarginOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}
	if s.newOrderRespType != nil {
		m["newOrderRespType"] = *s.newOrderRespType
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = s.quoteOrderQty
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.icebergQty != nil {
		m["icebergQty"] = *s.icebergQty
	}
	if s.sideEffectType != nil {
		m["sideEffectType"] = *s.sideEffectType
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateMarginOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/margin/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateUMOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateMarginOrderResponse define create order response
type CreateMarginOrderResponse struct {
	Symbol            string           `json:"symbol"`
	OrderID           int64            `json:"orderId"`
	ClientOrderID     string           `json:"clientOrderId"`
	Price             string           `json:"price"`
	OrigQuantity      string           `json:"origQty"`
	ExecutedQuantity  string           `json:"executedQty"`
	CumQuote          string           `json:"cumQuote"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Status            OrderStatusType  `json:"status"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	Side              SideType         `json:"side"`
	UpdateTime        int64            `json:"updateTime"`
	AvgPrice          string           `json:"avgPrice"`
	PositionSide      PositionSideType `json:"positionSide"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}
