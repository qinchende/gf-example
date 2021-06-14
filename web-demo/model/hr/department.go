package hr

type Department struct {
	Id          int16
	SubDepartId int16
	Name        string `json:"name" binding:"required"`
}
