package models

type Messages struct {
	TRANSACTIONTYPE  int    
	BLOCKHASH        string 
	BLOCKNUMBER      string 
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
	STATUS           string 
	ORDERTYPE        string 
	ISPDMP           bool   
}

// CurrencyInfo is...
type CurrencyInfo struct {
	Price    string 
	Currency string 
}
