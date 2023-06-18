package dtos

type CreateActionParameter struct {
	Name     string `json:"key"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type CreateHttpAction struct {
	Name       string                  `json:"name"`
	Url        string                  `json:"url"`
	Method     string                  `json:"method"`
	Body       string                  `json:"body"`
	Headers    string                  `json:"headers"`
	TimeoutMs  uint64                  `json:"timeoutMs"`
	Parameters []CreateActionParameter `json:"parameters"`
}
