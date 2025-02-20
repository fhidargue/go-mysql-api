package models

type Student struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name" binding:"required"`     
	LastName string   `json:"lastName" binding:"required"` 
	Age      *int64   `json:"age" binding:"required"`              
	Grade    *float64 `json:"grade" binding:"required"`           
}

type StudentPatch struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name,omitempty"`     
	LastName string   `json:"lastName,omitempty"` 
	Age      *int64   `json:"age,omitempty"`              
	Grade    *float64 `json:"grade,omitempty"`           
}
