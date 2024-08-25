package models

type BroadcastRequest struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp int64  `json:"timestamp"`
}

type BroadcastResponse struct {
	TxHash string `json:"tx_hash"`
}

type TransactionStatusResponse struct {
	TxStatus string `json:"tx_status"`
}
