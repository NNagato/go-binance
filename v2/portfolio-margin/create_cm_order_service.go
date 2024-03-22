package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateCMOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         *string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	newOrderRespType *NewOrderRespType
}

func (s *CreateCMOrderService) Symbol(symbol string) *CreateCMOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateCMOrderService) Side(side SideType) *CreateCMOrderService {
	s.side = side
	return s
}

func (s *CreateCMOrderService) PositionSide(positionSide PositionSideType) *CreateCMOrderService {
	s.positionSide = &positionSide
	return s
}

func (s *CreateCMOrderService) Type(orderType OrderType) *CreateCMOrderService {
	s.orderType = orderType
	return s
}

func (s *CreateCMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateCMOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateCMOrderService) Quantity(quantity string) *CreateCMOrderService {
	s.quantity = &quantity
	return s
}

func (s *CreateCMOrderService) ReduceOnly(reduceOnly bool) *CreateCMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *CreateCMOrderService) Price(price string) *CreateCMOrderService {
	s.price = &price
	return s
}

func (s *CreateCMOrderService) NewClientOrderID(newClientOrderID string) *CreateCMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *CreateCMOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateCMOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *CreateCMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

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
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateCMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/cm/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateCMOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCMOrderResponse define create order response
type CreateCMOrderResponse struct {
	ClientOrderID     string           `json:"clientOrderId"`
	CumQty            string           `json:"cumQty"`
	CumBase           string           `json:"cumBase"`
	ExecutedQty       string           `json:"executedQty"`
	OrderID           int64            `json:"orderId"`
	AvgPrice          string           `json:"avgPrice"`
	OrigQty           string           `json:"origQty"`
	Price             string           `json:"price"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Side              SideType         `json:"side"`
	PositionSide      PositionSideType `json:"positionSide"`
	Status            OrderStatusType  `json:"status"`
	Symbol            string           `json:"symbol"`
	Pair              string           `json:"pair"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	UpdateTime        int64            `json:"updateTime"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}
