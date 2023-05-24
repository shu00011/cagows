package model

// json:○○と書くだけでreqをparseでかける。
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
