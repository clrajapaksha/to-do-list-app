package entities

type Task struct {
	Id          string `json:"id"`
	Titile      string `json:"title"`
	Description string `json:"description"`
}

type TaskCreate struct {
	Titile      string `json:"title"`
	Description string `json:"description"`
}
