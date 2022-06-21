package entity

type Mark struct {
	StudentID string `json:"student_no"`
	ModuleID  string `json:"module_code"`
	Mark      int    `json:"mark"`
}
