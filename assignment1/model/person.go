package model

type Students struct {
	ID          string `json:"id"`
	StudentCode string `json:"student_code"`
	StudentName string `json:"student_name"`
	StudentAddr string `json:"student_address"`
	StudentOcc  string `json:"student_occupation"`
	Reason      string `json:"joining_reason"`
}

type People struct {
	People []Students `json:"participants"`
}
