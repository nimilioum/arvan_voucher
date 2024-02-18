package schemas

type UseVoucherRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}
