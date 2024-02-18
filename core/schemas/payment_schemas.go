package schemas

type DepositWithdrawRequest struct {
	Phone  string `json:"phone"`
	Amount uint   `json:"amount"`
}
