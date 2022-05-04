package consumer

type Messages struct {
	TRANSACTIONTYPE  int    `json:"TRANSACTIONTYPE"`
	BLOCKHASH        string `json:"BLOCKHASH"`
	BLOCKNUMBER      string `json:"BLOCKNUMBER"`
	TRANSACTIONHASH  string
	TRANSACTIONINDEX int
	LOGINDEX         int
	CONTRACTADDRESS  string
	EXCHANGEADDRESS  string
	FROMADDRESS      string
	TOADDRESS        string
	TOKENID          string
	TOKENURI         string
	TIMESTAMP        string
	TIMEUNIX         int
	VALUE            CurrencyInfo
	STATUS           string	`json:"STATUS"`
	ORDERTYPE        string `json:"ORDERTYPE"`
	ISPDMP           bool `json:"ISPDMP"`
}

// CurrencyInfo is...
type CurrencyInfo struct {
	Price    string `json:"Price"`
	Currency string `json:"Currency"`
}
