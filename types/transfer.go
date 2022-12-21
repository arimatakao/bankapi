package types

type TransferReq struct {
	IsSuccess     bool   `json:"isSuccess"`
	OperationDate string `json:"operationDate"`
}

type Transfer struct {
	To      string `json:"to"`
	Count   string `json:"count"`
	Message string `json:"message"`
}
