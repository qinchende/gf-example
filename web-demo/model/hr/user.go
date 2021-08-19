package hr

type User struct {
	Id       int
	Account  string `pms:"account" binding:"required"`
	Name     string `pms:"name" binding:"required"`
	Nickname string `pms:"nickname"`
	Age      int16  `pms:"age" binding:"required"`
}
