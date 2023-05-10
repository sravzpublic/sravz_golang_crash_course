package models

type Crypto struct {
	Type               string `json:"type"`
	BaseCurrency       string `json:"base_currency"`
	QuoteCurrency      string `json:"quote_currency"`
	Status             string `json:"status"`
	QuantityIncrement  string `json:"quantity_increment"`
	TickSize           string `json:"tick_size"`
	TakeRate           string `json:"take_rate"`
	MakeRate           string `json:"make_rate"`
	FeeCurrency        string `json:"fee_currency"`
	MarginTrading      bool   `json:"margin_trading"`
	MaxInitialLeverage string `json:"max_initial_leverage"`
}

type CryptoTable []Crypto

func (cT CryptoTable) Len() int           { return len(cT) }
func (cT CryptoTable) Less(i, j int) bool { return cT[i].BaseCurrency < cT[j].BaseCurrency }
func (cT CryptoTable) Swap(i, j int)      { cT[i], cT[j] = cT[j], cT[i] }
