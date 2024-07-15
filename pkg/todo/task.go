package todo

type Task struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty" binding:"required"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}
