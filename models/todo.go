package models

type todo struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Cmpleted boobool `json:"completed"`
}
