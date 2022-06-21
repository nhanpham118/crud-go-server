package entity

type Student struct {
	StudentID string  `json:"studentID"`
	SurName   string  `json:"surname"`
	ForeName  string  `json:"forename"`
}

type FullStudent struct {
	StudentID string  `json:"studentID"`
	SurName   string  `json:"surname"`
	ForeName  string  `json:"forename"`
	Scores    []Score `json:"scores"` // From mark - module table
}

type Score struct {
	ModuleName string `json:"module_name"`
	ModuleCode string `json:"module_code"`
	Mark       int    `json:"mark"`
}
