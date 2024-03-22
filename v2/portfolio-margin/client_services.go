package portfoliomargin

// NewChangeUMInitialLeverageService
func (c *Client) NewChangeUMInitialLeverageService() *ChangeUMInitialLeverageService {
	return &ChangeUMInitialLeverageService{c: c}
}

// NewCreateCMOrderService create coin margin order
func (c *Client) NewCreateCMOrderService() *CreateCMOrderService {
	return &CreateCMOrderService{c: c}
}

// NewCreateMarginOrderService
func (c *Client) NewCreateMarginOrderService() *CreateMarginOrderService {
	return &CreateMarginOrderService{c: c}
}

// NewCreateUMConditionalOrderService
func (c *Client) NewCreateUMConditionalOrderService() *CreateUMConditionalOrderService {
	return &CreateUMConditionalOrderService{c: c}
}

// NewCreateUMOrderService
func (c *Client) NewCreateUMOrderService() *CreateUMOrderService {
	return &CreateUMOrderService{c: c}
}

// NewGetAccountBalanceService
func (c *Client) NewGetAccountBalanceService() *GetAccountBalanceService {
	return &GetAccountBalanceService{c: c}
}

// NewGetAccountInfoService
func (c *Client) NewGetAccountInfoService() *GetAccountInfoService {
	return &GetAccountInfoService{c: c}
}

// NewGetAllUMOrdersService
func (c *Client) NewGetAllUMOrdersService() *GetAllUMOrdersService {
	return &GetAllUMOrdersService{c: c}
}

// NewGetNegativeBalanceInterestHistoryService
func (c *Client) NewGetNegativeBalanceInterestHistoryService() *GetNegativeBalanceInterestHistoryService {
	return &GetNegativeBalanceInterestHistoryService{c: c}
}

// NewGetUMAccountDetailService
func (c *Client) NewGetUMAccountDetailService() *GetUMAccountDetailService {
	return &GetUMAccountDetailService{c: c}
}

// NewGetUMCommissionRate
func (c *Client) NewGetUMCommissionRate() *GetUMCommissionRate {
	return &GetUMCommissionRate{c: c}
}

// NewGetUMForceOrdersService
func (c *Client) NewGetUMForceOrdersService() *GetUMForceOrdersService {
	return &GetUMForceOrdersService{c: c}
}

// NewGetUMIncomeHistoryService
func (c *Client) NewGetUMIncomeHistoryService() *GetUMIncomeHistoryService {
	return &GetUMIncomeHistoryService{c: c}
}

// NewGetUMLeverageService
func (c *Client) NewGetUMLeverageService() *GetUMLeverageService {
	return &GetUMLeverageService{c: c}
}

// NewGetUMOrderService
func (c *Client) NewGetUMOrderService() *GetUMOrderService {
	return &GetUMOrderService{c: c}
}

// NewGetUMPositionRiskService
func (c *Client) NewGetUMPositionRiskService() *GetUMPositionRiskService {
	return &GetUMPositionRiskService{c: c}
}

// NewGetUMTradeList
func (c *Client) NewGetUMTradeList() *GetUMTradeList {
	return &GetUMTradeList{c: c}
}

// NewMarginAccountBorrowService
func (c *Client) NewMarginAccountBorrowService() *MarginAccountBorrowService {
	return &MarginAccountBorrowService{c: c}
}

// NewMarginAccountRepayService
func (c *Client) NewMarginAccountRepayService() *MarginAccountRepayService {
	return &MarginAccountRepayService{c: c}
}

// NewPingService
func (c *Client) NewPingService() *PingService {
	return &PingService{c: c}
}

// NewTransferBnbService
func (c *Client) NewTransferBnbService() *TransferBnbService {
	return &TransferBnbService{c: c}
}

// NewStartUserStreamDataService
func (c *Client) NewStartUserStreamDataService() *StartUserStreamDataService {
	return &StartUserStreamDataService{c: c}
}

// NewKeepAliveUserStreamDataService
func (c *Client) NewKeepAliveUserStreamDataService() *KeepAliveUserStreamDataService {
	return &KeepAliveUserStreamDataService{c: c}
}

// NewCloseUserStreamDataService
func (c *Client) NewCloseUserStreamDataService() *CloseUserStreamDataService {
	return &CloseUserStreamDataService{c: c}
}
