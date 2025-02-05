package model

type BitGetBook struct {
	Action string `json:"action"`
	Arg    struct {
		InstType string `json:"instType"`
		Channel  string `json:"channel"`
		InstId   string `json:"instId"`
	} `json:"arg"`
	Data []BookData `json:"data"`
	Ts   int64      `json:"ts"`
}

type BookData struct {
	InstId          string `json:"instId"`
	LastPr          string `json:"lastPr"`
	BidPr           string `json:"bidPr"`
	AskPr           string `json:"askPr"`
	BidSz           string `json:"bidSz"`
	AskSz           string `json:"askSz"`
	Open24H         string `json:"open24h"`
	High24H         string `json:"high24h"`
	Low24H          string `json:"low24h"`
	Change24H       string `json:"change24h"`
	FundingRate     string `json:"fundingRate"`
	NextFundingTime string `json:"nextFundingTime"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	HoldingAmount   string `json:"holdingAmount"`
	BaseVolume      string `json:"baseVolume"`
	QuoteVolume     string `json:"quoteVolume"`
	OpenUtc         string `json:"openUtc"`
	SymbolType      string `json:"symbolType"`
	Symbol          string `json:"symbol"`
	DeliveryPrice   string `json:"deliveryPrice"`
	Ts              string `json:"ts"`
}
