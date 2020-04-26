package entity

// DummyEmployee entity
type DummyEmployee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// DummyEmployees entity's slice
type DummyEmployees []DummyEmployee
