package entity

type Issue struct {
	Title       string `json:"title,required"`
	Description string `json:"description,required"`
}

type Bug struct {
	Title       string `json:"title"`
	Description string `json:"description,required"`
}

type Task struct {
	Title    string `json:"title,required"`
	Category string `json:"category"`
}

type Category string

const (
	Maintenance Category = "maintenance"
	Research    Category = "research"
	Test        Category = "test"
)

type TaskType string

const (
	TypeIssue TaskType = "issue"
	TypeBug   TaskType = "bug"
	TypeTask  TaskType = "task"
)
