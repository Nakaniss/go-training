package dto

// クライアントに返す構造体。これはテストなので良いが、本名やEmailが表示されてしまうため注意。
type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
