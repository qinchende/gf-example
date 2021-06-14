package hr

type Title struct {
	Id    int16
	Level int16
	Name  string `json:"name" binding:"required"`
}
