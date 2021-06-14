package hr

type User struct {
	Id       int
	Account  string `json:"account" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}
