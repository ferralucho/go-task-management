package entity

type Issue struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Bug struct {
	Description string `json:"description" binding:"required"`
}

type Task struct {
	Title    string `json:"title"binding:"required"`
	Category string `json:"category"`
}

type InternalCard struct {
	Name     string   `json:"name" binding:"required"`
	Desc     string   `json:"desc"`
	IdLabels []string `json:"idLabels" binding:"required"`
}

type Card struct {
	Name      string   `json:"name" binding:"required"`
	Desc      string   `json:"desc"`
	IdList    string   `json:"idList"`
	IdLabels  []string `json:"idLabels"`
	ShortLink string   `json:"shortLink"`
	ShortURL  string   `json:"shortUrl"`
	URL       string   `json:"url"`
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
