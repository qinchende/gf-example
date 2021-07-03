package hr

type Title struct {
	Id    int16
	Level int16
	Name  string `json:"name" pms:"name" binding:"required"`
	Desc  string `json:"desc" pms:"desc"`
}
