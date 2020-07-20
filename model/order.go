package model

// 数据库orders表
type Order struct {
	ID        uint    `json:"id"`
	UserName  string  `json:"user_name"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	FileUrl   string  `json:"file_url"`
	CreatedAt int64   `json:"created_at"`
}

type AddOrderReq struct {
	ID       string  `form:"id"`
	UserName string  `form:"user_name"`
	Amount   float64 `form:"amount"`
	Status   string  `form:"status"`
	FileUrl  string  `form:"file_url"`
}

type GetOrderReq struct {
	ID uint `form:"id"`
}

type OrderListReq struct {
	UserName string `form:"user_name"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
