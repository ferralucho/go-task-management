package entity

type Issue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Bug struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Task struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}
