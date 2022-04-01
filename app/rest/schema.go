package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type CreateEmployeeRequest struct {
	Name string `json:"name"`
}

type Employee struct {
	Base `json:",inline"`
	Name string `json:"name"`
}

type SearchEmployeesRequest struct {
	PageToken string `json:"page_token" query:"page_token"`
	PageSize  int    `json:"page_size" query:"page_size"`
}

type SearchEmployeesResponse struct {
	Employees     []*Employee `json:"employees"`
	NextPageToken string      `json:"next_page_token"`
}
