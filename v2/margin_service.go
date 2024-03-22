package binance

import (
	"context"
	"net/http"
	"strings"
)

// TransactionResponse define transaction response
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

// MarginRepayResponse define margin repay response
type MarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarginRepay define margin repay
type MarginRepay struct {
	Asset     string                `json:"asset"`
	Amount    string                `json:"amount"`
	Interest  string                `json:"interest"`
	Principal string                `json:"principal"`
	Timestamp int64                 `json:"timestamp"`
	Status    MarginRepayStatusType `json:"status"`
	TxID      int64                 `json:"txId"`
}

// GetIsolatedMarginAccountService gets isolated margin account info
type GetIsolatedMarginAccountService struct {
	c *Client

	symbols []string
}

// Symbols set symbols to the isolated margin account
func (s *GetIsolatedMarginAccountService) Symbols(symbols ...string) *GetIsolatedMarginAccountService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *GetIsolatedMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IsolatedMarginAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/account",
		secType:  secTypeSigned,
	}

	if len(s.symbols) > 0 {
		r.setParam("symbols", strings.Join(s.symbols, ","))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IsolatedMarginAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// IsolatedMarginAccount defines isolated user assets of margin account
type IsolatedMarginAccount struct {
	TotalAssetOfBTC     string                `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string                `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string                `json:"totalNetAssetOfBtc"`
	Assets              []IsolatedMarginAsset `json:"assets"`
}

// IsolatedMarginAsset defines isolated margin asset information, like margin level, liquidation price... etc
type IsolatedMarginAsset struct {
	Symbol     string            `json:"symbol"`
	QuoteAsset IsolatedUserAsset `json:"quoteAsset"`
	BaseAsset  IsolatedUserAsset `json:"baseAsset"`

	IsolatedCreated   bool   `json:"isolatedCreated"`
	Enabled           bool   `json:"enabled"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// IsolatedUserAsset defines isolated user assets of the margin account
type IsolatedUserAsset struct {
	Asset         string `json:"asset"`
	Borrowed      string `json:"borrowed"`
	Free          string `json:"free"`
	Interest      string `json:"interest"`
	Locked        string `json:"locked"`
	NetAsset      string `json:"netAsset"`
	NetAssetOfBtc string `json:"netAssetOfBtc"`

	BorrowEnabled bool   `json:"borrowEnabled"`
	RepayEnabled  bool   `json:"repayEnabled"`
	TotalAsset    string `json:"totalAsset"`
}

// GetMarginAccountService get margin account info
type GetMarginAccountService struct {
	c *Client
}

// Do send request
func (s *GetMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginAccount define margin account info
type MarginAccount struct {
	BorrowEnabled       bool        `json:"borrowEnabled"`
	MarginLevel         string      `json:"marginLevel"`
	TotalAssetOfBTC     string      `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string      `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string      `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool        `json:"tradeEnabled"`
	TransferEnabled     bool        `json:"transferEnabled"`
	UserAssets          []UserAsset `json:"userAssets"`
}

// UserAsset define user assets of margin account
type UserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// MarginAsset define margin asset info
type MarginAsset struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// MarginPair define margin pair info
type MarginPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginAllPairsService get margin pair info
type GetMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAllPair, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	res = make([]*MarginAllPair, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	return res, nil
}

// MarginAllPair define margin pair info
type MarginAllPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginPriceIndexService get margin price index
type GetMarginPriceIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPriceIndexService) Symbol(symbol string) *GetMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPriceIndexService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPriceIndex, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/priceIndex",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPriceIndex)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginPriceIndex define margin price index
type MarginPriceIndex struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// ListMarginTradesService list trades
type ListMarginTradesService struct {
	c          *Client
	symbol     string
	startTime  *int64
	endTime    *int64
	limit      *int
	fromID     *int64
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginTradesService) Symbol(symbol string) *ListMarginTradesService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginTradesService) IsIsolated(isIsolated bool) *ListMarginTradesService {
	s.isIsolated = isIsolated
	return s
}

// StartTime set starttime
func (s *ListMarginTradesService) StartTime(startTime int64) *ListMarginTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginTradesService) EndTime(endTime int64) *ListMarginTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginTradesService) Limit(limit int) *ListMarginTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListMarginTradesService) FromID(fromID int64) *ListMarginTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *ListMarginTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeV3, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeV3{}, err
	}
	res = make([]*TradeV3, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*TradeV3{}, err
	}
	return res, nil
}

// GetMaxBorrowableService get max borrowable of asset
type GetMaxBorrowableService struct {
	c              *Client
	asset          string
	isolatedSymbol string
}

// Asset set asset
func (s *GetMaxBorrowableService) Asset(asset string) *GetMaxBorrowableService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *GetMaxBorrowableService) IsolatedSymbol(isolatedSymbol string) *GetMaxBorrowableService {
	s.isolatedSymbol = isolatedSymbol
	return s
}

// Do send request
func (s *GetMaxBorrowableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxBorrowable, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.isolatedSymbol != "" {
		r.setParam("isolatedSymbol", s.isolatedSymbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxBorrowable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxBorrowable define max borrowable response
type MaxBorrowable struct {
	Amount string `json:"amount"`
}

// GetMaxTransferableService get max transferable of asset
type GetMaxTransferableService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMaxTransferableService) Asset(asset string) *GetMaxTransferableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxTransferableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxTransferable, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxTransferable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxTransferable)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MaxTransferable define max transferable response
type MaxTransferable struct {
	Amount string `json:"amount"`
}

// StartIsolatedMarginUserStreamService create listen key for margin user stream service
type StartIsolatedMarginUserStreamService struct {
	c      *Client
	symbol string
}

// Symbol sets the user stream to isolated margin user stream
func (s *StartIsolatedMarginUserStreamService) Symbol(symbol string) *StartIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *StartIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveIsolatedMarginUserStreamService updates listen key for isolated margin user data stream
type KeepaliveIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string
	symbol    string
}

// Symbol set symbol to the isolated margin keepalive request
func (s *KeepaliveIsolatedMarginUserStreamService) Symbol(symbol string) *KeepaliveIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// ListenKey set listen key
func (s *KeepaliveIsolatedMarginUserStreamService) ListenKey(listenKey string) *KeepaliveIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseIsolatedMarginUserStreamService delete listen key
type CloseIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string

	symbol string
}

// ListenKey set listen key
func (s *CloseIsolatedMarginUserStreamService) ListenKey(listenKey string) *CloseIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Symbol set symbol to the isolated margin user stream close request
func (s *CloseIsolatedMarginUserStreamService) Symbol(symbol string) *CloseIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CloseIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// StartMarginUserStreamService create listen key for margin user stream service
type StartMarginUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	j, err := newJSON(data)
	if err != nil {
		return "", err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseMarginUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// GetAllMarginAssetsService get margin pair info
type GetAllMarginAssetsService struct {
	c *Client
}

// Do send request
func (s *GetAllMarginAssetsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAsset, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/margin/allAssets",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAsset{}, err
	}
	res = make([]*MarginAsset, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAsset{}, err
	}
	return res, nil
}

// GetIsolatedMarginAllPairsService get isolated margin pair info
type GetIsolatedMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetIsolatedMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*IsolatedMarginAllPair, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	res = make([]*IsolatedMarginAllPair, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	return res, nil
}

// IsolatedMarginAllPair define isolated margin pair info
type IsolatedMarginAllPair struct {
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}
