package portfoliomargin

import (
	"context"
	"encoding/json"
	"net/http"
)

type CreateUMConditionalOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	strategyType            *StrategyType
	timeInForce             *TimeInForceType
	quantity                *string
	reduceOnly              *string
	price                   *string
	workingType             *WorkingType
	priceProtect            *string
	newClientStrategyId     *string
	stopPrice               *string
	activationPrice         *string
	callbackRate            *string
	selfTradePreventionMode *SelfTradePreventionMode
	goodTillDate            *int64
}

func (s *CreateUMConditionalOrderService) Symbol(symbol string) *CreateUMConditionalOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateUMConditionalOrderService) Side(side SideType) *CreateUMConditionalOrderService {
	s.side = side
	return s
}

func (s *CreateUMConditionalOrderService) PositionSide(positionSide PositionSideType) *CreateUMConditionalOrderService {
	s.positionSide = &positionSide
	return s
}

func (s *CreateUMConditionalOrderService) StrategyType(strategyType StrategyType) *CreateUMConditionalOrderService {
	s.strategyType = &strategyType
	return s
}

func (s *CreateUMConditionalOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUMConditionalOrderService {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateUMConditionalOrderService) Quantity(quantity string) *CreateUMConditionalOrderService {
	s.quantity = &quantity
	return s
}

func (s *CreateUMConditionalOrderService) ReduceOnly(reduceOnly string) *CreateUMConditionalOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *CreateUMConditionalOrderService) Price(price string) *CreateUMConditionalOrderService {
	s.price = &price
	return s
}

func (s *CreateUMConditionalOrderService) WorkingType(workingType WorkingType) *CreateUMConditionalOrderService {
	s.workingType = &workingType
	return s
}

func (s *CreateUMConditionalOrderService) PriceProtect(priceProtect string) *CreateUMConditionalOrderService {
	s.priceProtect = &priceProtect
	return s
}

func (s *CreateUMConditionalOrderService) NewClientStrategyId(newClientStrategyId string) *CreateUMConditionalOrderService {
	s.newClientStrategyId = &newClientStrategyId
	return s
}

func (s *CreateUMConditionalOrderService) StopPrice(stopPrice string) *CreateUMConditionalOrderService {
	s.stopPrice = &stopPrice
	return s
}

func (s *CreateUMConditionalOrderService) ActivationPrice(activationPrice string) *CreateUMConditionalOrderService {
	s.activationPrice = &activationPrice
	return s
}

func (s *CreateUMConditionalOrderService) CallbackRate(callbackRate string) *CreateUMConditionalOrderService {
	s.callbackRate = &callbackRate
	return s
}

func (s *CreateUMConditionalOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateUMConditionalOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateUMConditionalOrderService) GoodTillDate(goodTillDate int64) *CreateUMConditionalOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

func (s *CreateUMConditionalOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
	}
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.strategyType != nil {
		m["strategyType"] = *s.strategyType
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
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.newClientStrategyId != nil {
		m["newClientStrategyId"] = *s.newClientStrategyId
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.goodTillDate != nil {
		m["goodTillDate"] = *s.goodTillDate
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateUMConditionalOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUMConditionalOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/conditional/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateUMConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUMConditionalOrderResponse define create order response
type CreateUMConditionalOrderResponse struct {
	NewClientStrategyID     string                  `json:"newClientStrategyId"`
	StrategyID              int                     `json:"strategyId"`
	StrategyStatus          string                  `json:"strategyStatus"`
	StrategyType            StrategyType            `json:"strategyType"`
	OrigQty                 string                  `json:"origQty"`
	Price                   string                  `json:"price"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	Side                    SideType                `json:"side"`
	PositionSide            PositionSideType        `json:"positionSide"`
	StopPrice               string                  `json:"stopPrice"`
	Symbol                  string                  `json:"symbol"`
	TimeInForce             TimeInForceType         `json:"timeInForce"`
	ActivatePrice           string                  `json:"activatePrice"`
	PriceRate               string                  `json:"priceRate"`
	BookTime                int64                   `json:"bookTime"`
	UpdateTime              int64                   `json:"updateTime"`
	WorkingType             WorkingType             `json:"workingType"`
	PriceProtect            bool                    `json:"priceProtect"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate"`
	RateLimitOrder10s       string                  `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m        string                  `json:"rateLimitOrder1m,omitempty"`
}
