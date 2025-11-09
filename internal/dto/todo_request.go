package dto

type CreateTodoRequest struct {
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	Completed bool   `json:"completed"`
}
