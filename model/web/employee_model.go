package web

type EmployeeCreateRequest struct {
	Name      string `json:"name" validate:"required,max=32,min=10"`
	Role      string `json:"role" validate:"required,max=32,min=3"` // e.g., Cashier, Manager
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone_number" validate:"required,min=10,max=30"`
	DateHired string `json:"date_hired]" validate:"required,min=6,max=30"`
}

type EmployeeUpdateRequest struct {
	Id        uint64 `json:"id" validate:"required,gte=0"`
	Name      string `json:"name" validate:"required,max=32,min=10"`
	Role      string `json:"role" validate:"required,max=32,min=3"` // e.g., Cashier, Manager
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone_number" validate:"required,min=10,max=30"`
	DateHired string `json:"date_hired]" validate:"required,min=6,max=30"`
}

type EmployeeResponse struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone_number"`
	DateHired string `json:"date_hired]"`
}
