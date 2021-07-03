package hr

type User struct {
	Id       int
	Account  string `json:"account" pms:"account" binding:"required"`
	Name     string `json:"name" pms:"name" binding:"required"`
	Nickname string `json:"nickname"`
	Age      int16  `json:"age" pms:"age" binding:"required"`
}
