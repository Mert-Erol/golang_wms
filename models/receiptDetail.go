package models

type ReceiptDetail struct {
	StockCode       string
	Quantity        string
	TransactionType int //0 = Receipt, 1 = Order
}
