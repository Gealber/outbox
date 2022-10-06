package model

type Cat struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Weight       float64 `json:"weight"`
	Intelligence int     `json:"intelligence"`
	Laziness     int     `json:"laziness"`
	Curiosity    int     `json:"curiosity"`
	Sociability  int     `json:"sociability"`
	Egoism       int     `json:"egoism"`
	MiauPower    int     `json:"miau_power"`
	Attack       int     `json:"attack"`
}

type CatResponse struct {
	Cat Cat `json:"cat"`
}

type CreateCatRequest struct {
	Name         string  `json:"name" binding:"required,lte=100"`
	Color        string  `json:"color" binding:"required,lte=20"`
	Weight       float64 `json:"weight" binding:"required"`
	Intelligence int     `json:"intelligence" binding:"required,min=1,max=5"`
	Laziness     int     `json:"laziness" binding:"required,min=1,max=10"`
	Curiosity    int     `json:"curiosity" binding:"required,min=5,max=10"`
	Sociability  int     `json:"sociability" binding:"required,min=1,max=10"`
	Egoism       int     `json:"egoism" binding:"required,min=1,max=10"`
	MiauPower    int     `json:"miau_power" binding:"required,min=1,max=10"`
	Attack       int     `json:"attack" binding:"required,min=1,max=10"`
}

type UpdateCatRequest struct {
	Name         string  `json:"name" binding:"omitempty,lte=100"`
	Color        string  `json:"color" binding:"omitempty,lte=20"`
	Weight       float64 `json:"weight" binding:"omitempty"`
	Intelligence int     `json:"intelligence" binding:"omitempty,min=1,max=5"`
	Laziness     int     `json:"laziness" binding:"omitempty,min=1,max=10"`
	Curiosity    int     `json:"curiosity" binding:"omitempty,min=5,max=10"`
	Sociability  int     `json:"sociability" binding:"omitempty,min=1,max=10"`
	Egoism       int     `json:"egoism" binding:"omitempty,min=1,max=10"`
	MiauPower    int     `json:"miau_power" binding:"omitempty,min=1,max=10"`
	Attack       int     `json:"attack" binding:"omitempty,min=1,max=10"`
}
