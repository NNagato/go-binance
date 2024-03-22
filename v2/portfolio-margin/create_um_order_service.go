package portfoliomargin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUMOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                *string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	newOrderRespType        *NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64
}

func (s *CreateUMOrderService) Symbol(symbol string) *CreateUMOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateUMOrderService) Side(side SideType) *CreateUMOrderService {
	s.side = side
	return s
}

func (s *CreateUMOrderService) PositionSide(positionSide PositionSideType) *CreateUMOrderService {
	s.positionSide = &positionSide
	return s
}

func (s *CreateUMOrderService) Type(orderType OrderType) *CreateUMOrderService {
	s.orderType = orderType
	return s
}

func (s *CreateUMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUMOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateUMOrderService) Quantity(quantity string) *CreateUMOrderService {
	s.quantity = &quantity
	return s
}

func (s *CreateUMOrderService) ReduceOnly(reduceOnly bool) *CreateUMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *CreateUMOrderService) Price(price string) *CreateUMOrderService {
	s.price = &price
	return s
}

func (s *CreateUMOrderService) NewClientOrderID(newClientOrderID string) *CreateUMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *CreateUMOrderService) NewOrderRespType(newOrderRespType NewOrderRespType) *CreateUMOrderService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *CreateUMOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateUMOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateUMOrderService) GoodTillDate(goodTillDate int64) *CreateUMOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

func (s *CreateUMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
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
		fmt.Println("qty", s.quantity)
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
func (s *CreateUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/order", opts...)
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

// CreateUMOrderResponse define create order response
type CreateUMOrderResponse struct {
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
