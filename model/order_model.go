package model

// 数据库orders表
type Order struct {
	ID        int64   `json:"id"`
	OrderId   string  `json:"order_id"`
	UserName  string  `json:"user_name"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	FileUrl   string  `json:"file_url"`
	CreatedAt int64   `json:"created_at"`
}
