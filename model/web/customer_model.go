package web

type CustomerCreateRequest struct {
	Name       string `json:"name" validate:"required,max=32,min=10"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone_number" validate:"required,min=10,max=30"`
	Address    string `json:"address" validate:"required,min=10,max=255"`
	LoyaltyPts int    `json:"loyalty_pts" validate:"required"`
}
type CustomerUpdateRequest struct {
	CustomerID uint64 `json:"id" validate:"required,gte=0"`
	Name       string `json:"name" validate:"required,max=32,min=10"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone_number" validate:"required,min=10,max=30"`
	Address    string `json:"address" validate:"required,min=10,max=255"`
	LoyaltyPts int    `json:"loyalty_pts" validate:"required"`
}

type CustomerResponse struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone_number"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_pts"`
}
